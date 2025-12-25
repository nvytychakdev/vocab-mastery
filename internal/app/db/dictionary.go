package db

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nvytychakdev/vocab-mastery/internal/app/model"
)

type DictionaryRepo interface {
	Create(userId uuid.UUID, title string) (uuid.UUID, error)
	DeleteByID(id uuid.UUID) error
	GetByID(id uuid.UUID) (*model.Dictionary, error)
	ListByUserId(userID uuid.UUID, opts *model.QueryOptions) ([]*model.Dictionary, int, error)
}

type dictionaryRepo struct {
	conn *pgxpool.Pool
	psql sq.StatementBuilderType
}

func (db *PostgresDB) Dictionary() DictionaryRepo {
	return &dictionaryRepo{conn: db.conn, psql: db.psql}
}

func (db *dictionaryRepo) Create(userId uuid.UUID, title string) (uuid.UUID, error) {
	query, args, err := db.psql.
		Insert("dictionaries").
		Columns("owner_id", "title", "is_default").
		Values(userId, title, false).
		Suffix("RETURNING \"id\"").ToSql()

	if err != nil {
		return uuid.Nil, err
	}

	var dictionaryId uuid.UUID
	err = db.conn.QueryRow(context.Background(), query, args...).Scan(&dictionaryId)
	return dictionaryId, err
}

func (db *dictionaryRepo) DeleteByID(dictionaryId uuid.UUID) error {
	query, args, err := db.psql.Delete("dictionaries").Where(sq.Eq{"id": dictionaryId}).ToSql()

	if err != nil {
		return err
	}

	_, err = db.conn.Exec(context.Background(), query, args...)
	return err
}

func (db *dictionaryRepo) GetByID(dictionaryId uuid.UUID) (*model.Dictionary, error) {
	query, args, err := db.psql.
		Select("id", "owner_id", "title", "level", "is_default", "created_at").
		From("dictionaries").Where(sq.Eq{"id": dictionaryId}).ToSql()

	if err != nil {
		return nil, err
	}

	var dictionary model.Dictionary
	err = db.conn.QueryRow(context.Background(), query, args...).Scan(
		&dictionary.ID,
		&dictionary.OwnerID,
		&dictionary.Title,
		&dictionary.Level,
		&dictionary.IsDefault,
		&dictionary.CreatedAt,
	)
	return &dictionary, err
}

func (db *dictionaryRepo) ListByUserId(userId uuid.UUID, opts *model.QueryOptions) ([]*model.Dictionary, int, error) {

	queryBuilder := db.psql.
		Select("id", "owner_id", "title", "level", "is_default", "created_at").
		From("dictionaries").Where(
		sq.Or{
			sq.Eq{"owner_id": userId},
			sq.Eq{"is_default": true},
		},
	)

	query, args, err := ApplyQueryOptions(queryBuilder, opts).ToSql()
	if err != nil {
		return nil, 0, err
	}

	rows, err := db.conn.Query(context.Background(), query, args...)
	if err != nil {
		return nil, 0, err
	}

	defer rows.Close()

	var dictionaries = []*model.Dictionary{}
	for rows.Next() {
		var dictionary model.Dictionary
		err := rows.Scan(
			&dictionary.ID,
			&dictionary.OwnerID,
			&dictionary.Title,
			&dictionary.Level,
			&dictionary.IsDefault,
			&dictionary.CreatedAt,
		)

		if err != nil {
			return nil, 0, err
		}
		dictionaries = append(dictionaries, &dictionary)
	}

	totalQuery, totalArgs, err := db.psql.
		Select("COUNT(*)").From("dictionaries").
		Where(
			sq.Or{
				sq.Eq{"owner_id": userId},
				sq.Eq{"is_default": true},
			},
		).ToSql()

	if err != nil {
		return nil, 0, err
	}

	var total int
	err = db.conn.QueryRow(context.Background(), totalQuery, totalArgs...).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	return dictionaries, total, rows.Err()
}
