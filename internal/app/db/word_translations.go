package db

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nvytychakdev/vocab-mastery/internal/app/model"
)

type WordTranslationsRepo interface {
	ListAllByMeaningIDs(wordIDs uuid.UUIDs) ([]*model.WordTranslation, int, error)
}

type wordTranslationsRepo struct {
	conn *pgxpool.Pool
	psql sq.StatementBuilderType
}

func (db *PostgresDB) WordTranslation() WordTranslationsRepo {
	return &wordTranslationsRepo{conn: db.conn, psql: db.psql}
}

func (db *wordTranslationsRepo) ListAllByMeaningIDs(meaningIDs uuid.UUIDs) ([]*model.WordTranslation, int, error) {
	queryBuilder := db.psql.
		Select("id", "meaning_id", "language", "translation").
		From("word_translations").
		Where(sq.Eq{"meaning_id": meaningIDs})

	query, args, err := queryBuilder.ToSql()

	if err != nil {
		return nil, 0, err
	}

	rows, err := db.conn.Query(context.Background(), query, args...)
	if err != nil {
		return nil, 0, err
	}

	defer rows.Close()

	var translations = []*model.WordTranslation{}
	for rows.Next() {
		var translation model.WordTranslation
		err := rows.Scan(
			&translation.ID,
			&translation.MeaningID,
			&translation.Language,
			&translation.Translation,
		)

		if err != nil {
			return nil, 0, err
		}
		translations = append(translations, &translation)
	}

	return translations, len(translations), rows.Err()
}
