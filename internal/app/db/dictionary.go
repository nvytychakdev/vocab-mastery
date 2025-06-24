package db

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx"
	"github.com/nvytychakdev/vocab-mastery/internal/app/model"
)

type DictionaryRepo interface {
	Create(userId string, name string, description string) (string, error)
	DeleteByID(id string) error
	GetByID(id string) (*model.Dictionary, error)
	ListByUserId(userID string, opts *model.QueryOptions) ([]*model.Dictionary, int, error)
}

type dictionaryRepo struct {
	conn *pgx.Conn
	psql sq.StatementBuilderType
}

func (db *PostgresDB) Dictionary() DictionaryRepo {
	return &dictionaryRepo{conn: db.conn, psql: db.psql}
}

func (db *dictionaryRepo) Create(userId string, name string, description string) (string, error) {
	query, args, err := db.psql.
		Insert("dictionaries").
		Columns("user_id", "name", "description").
		Values(userId, name, description).
		Suffix("RETURNING \"id\"").ToSql()

	if err != nil {
		return "", err
	}

	var dictionaryId string
	err = db.conn.QueryRow(query, args...).Scan(&dictionaryId)
	return dictionaryId, err
}

func (db *dictionaryRepo) DeleteByID(dictionaryId string) error {
	query, args, err := db.psql.Delete("dictionaries").Where(sq.Eq{"id": dictionaryId}).ToSql()

	if err != nil {
		return err
	}

	_, err = db.conn.Exec(query, args...)
	return err
}

func (db *dictionaryRepo) GetByID(dictionaryId string) (*model.Dictionary, error) {
	query, args, err := db.psql.
		Select("id", "user_id", "name", "description", "created_at").
		From("dictionaries").Where(sq.Eq{"id": dictionaryId}).ToSql()

	if err != nil {
		return nil, err
	}

	var dictionary model.Dictionary
	err = db.conn.QueryRow(query, args...).Scan(
		&dictionary.ID,
		&dictionary.UserID,
		&dictionary.Name,
		&dictionary.Description,
		&dictionary.CreatedAt,
	)
	return &dictionary, err
}

func (db *dictionaryRepo) ListByUserId(userId string, opts *model.QueryOptions) ([]*model.Dictionary, int, error) {
	queryBuilder := db.psql.
		Select("id", "user_id", "name", "description", "created_at").
		From("dictionaries").Where(sq.Eq{"user_id": userId})

	query, args, err := ApplyQueryOptions(queryBuilder, opts).ToSql()

	if err != nil {
		return nil, 0, err
	}

	rows, err := db.conn.Query(query, args...)
	if err != nil {
		return nil, 0, err
	}

	defer rows.Close()

	var dictionaries = []*model.Dictionary{}
	for rows.Next() {
		var dictionary model.Dictionary
		err := rows.Scan(
			&dictionary.ID,
			&dictionary.UserID,
			&dictionary.Name,
			&dictionary.Description,
			&dictionary.CreatedAt,
		)

		if err != nil {
			return nil, 0, err
		}
		dictionaries = append(dictionaries, &dictionary)
	}

	totalQuery, totalArgs, err := db.psql.
		Select("COUNT(*)").From("dictionaries").
		Where(sq.Eq{"user_id": userId}).ToSql()

	if err != nil {
		return nil, 0, err
	}

	var total int
	err = db.conn.QueryRow(totalQuery, totalArgs...).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	return dictionaries, total, rows.Err()
}
