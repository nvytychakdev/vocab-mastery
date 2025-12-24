package db

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nvytychakdev/vocab-mastery/internal/app/model"
)

type TranslationRepo interface {
	Create(wordID uuid.UUID, word string, language string) (string, error)
	DeleteByID(id uuid.UUID) error
	GetByID(id uuid.UUID) (*model.Translation, error)
	ListByWordID(wordId uuid.UUID, opts *model.QueryOptions) ([]*model.Translation, int, error)
	ListByWordIDs(wordIDs uuid.UUIDs) ([]*model.Translation, error)
}

type translationRepo struct {
	conn *pgxpool.Pool
	psql sq.StatementBuilderType
}

func (db *PostgresDB) Translation() TranslationRepo {
	return &translationRepo{conn: db.conn, psql: db.psql}
}

func (db *translationRepo) Create(wordId uuid.UUID, word string, language string) (string, error) {
	query, args, err := db.psql.
		Insert("translations").
		Columns("word_id", "word", "language").
		Values(wordId, word, language).
		Suffix("RETURNING \"id\"").ToSql()

	if err != nil {
		return "", err
	}

	var translationId string
	err = db.conn.QueryRow(context.Background(), query, args...).Scan(&translationId)
	return translationId, err
}

func (db *translationRepo) DeleteByID(translationId uuid.UUID) error {
	query, args, err := db.psql.Delete("translations").Where(sq.Eq{"id": translationId}).ToSql()

	if err != nil {
		return err
	}

	_, err = db.conn.Exec(context.Background(), query, args...)
	return err
}

func (db *translationRepo) GetByID(translationId uuid.UUID) (*model.Translation, error) {
	query, args, err := db.psql.
		Select("id", "word_id", "word", "language", "created_at").
		From("translations").Where(sq.Eq{"id": translationId}).ToSql()

	if err != nil {
		return nil, err
	}

	var translation model.Translation
	err = db.conn.QueryRow(context.Background(), query, args...).Scan(
		&translation.ID,
		&translation.WordId,
		&translation.Word,
		&translation.Language,
		&translation.CreatedAt,
	)
	return &translation, err
}

func (db *translationRepo) ListByWordID(wordId uuid.UUID, opts *model.QueryOptions) ([]*model.Translation, int, error) {
	queryBuilder := db.psql.
		Select("id", "word_id", "word", "language", "created_at").
		From("translations").Where(sq.Eq{"word_id": wordId})

	query, args, err := ApplyQueryOptions(queryBuilder, opts).ToSql()

	if err != nil {
		return nil, 0, err
	}

	rows, err := db.conn.Query(context.Background(), query, args...)
	if err != nil {
		return nil, 0, err
	}

	defer rows.Close()

	var words = []*model.Translation{}
	for rows.Next() {
		var dictionary model.Translation
		err := rows.Scan(
			&dictionary.ID,
			&dictionary.WordId,
			&dictionary.Word,
			&dictionary.Language,
			&dictionary.CreatedAt,
		)

		if err != nil {
			return nil, 0, err
		}
		words = append(words, &dictionary)
	}

	totalQuery, totalArgs, err := db.psql.
		Select("COUNT(*)").From("translations").
		Where(sq.Eq{"word_id": wordId}).ToSql()

	if err != nil {
		return nil, 0, err
	}

	var total int
	err = db.conn.QueryRow(context.Background(), totalQuery, totalArgs...).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	return words, total, rows.Err()
}

func (db *translationRepo) ListByWordIDs(wordIDs uuid.UUIDs) ([]*model.Translation, error) {
	queryBuilder := db.psql.
		Select("id", "word_id", "word", "language", "created_at").
		From("translations").Where(sq.Eq{"word_id": wordIDs})

	query, args, err := queryBuilder.ToSql()

	if err != nil {
		return nil, err
	}

	rows, err := db.conn.Query(context.Background(), query, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	translations := []*model.Translation{}
	for rows.Next() {
		var t model.Translation
		rows.Scan(
			&t.ID,
			&t.WordId,
			&t.Word,
			&t.Language,
			&t.CreatedAt,
		)
		translations = append(translations, &t)
	}

	return translations, nil
}
