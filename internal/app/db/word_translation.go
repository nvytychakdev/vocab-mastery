package db

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nvytychakdev/vocab-mastery/internal/app/model"
)

type WordTranslationRepo interface {
	GetByID(translationID uuid.UUID) (*model.WordTranslation, error)
	ListByIDs(translationIDs uuid.UUIDs) ([]*model.WordTranslation, int, error)
	ListByWordID(wordID uuid.UUID) ([]*model.WordTranslation, int, error)
	ListByMeaningIDs(wordIDs uuid.UUIDs) ([]*model.WordTranslation, int, error)
}

type wordTranslationRepo struct {
	conn *pgxpool.Pool
	psql sq.StatementBuilderType
}

func (db *PostgresDB) WordTranslation() WordTranslationRepo {
	return &wordTranslationRepo{conn: db.conn, psql: db.psql}
}

func (db *wordTranslationRepo) ListByMeaningIDs(meaningIDs uuid.UUIDs) ([]*model.WordTranslation, int, error) {
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

func (db *wordTranslationRepo) GetByID(translationID uuid.UUID) (*model.WordTranslation, error) {
	query, args, err := db.psql.
		Select("wt.id", "wt.meaning_id", "wt.language", "wt.translation").
		From("word_translations wt").
		Where(sq.Eq{"wt.id": translationID}).ToSql()

	if err != nil {
		return nil, err
	}

	var translation model.WordTranslation
	err = db.conn.QueryRow(context.Background(), query, args...).
		Scan(
			&translation.ID,
			&translation.MeaningID,
			&translation.Language,
			&translation.Translation,
		)
	return &translation, err
}

func (db *wordTranslationRepo) ListByIDs(translationIDs uuid.UUIDs) ([]*model.WordTranslation, int, error) {
	query, args, err := db.psql.
		Select("wt.id", "wt.meaning_id", "wt.language", "wt.translation").
		From("word_translations wt").
		Join(
			"unnest(?::uuid[]) WITH ORDINALITY AS u(id, pos) ON wt.id = u.id",
			translationIDs,
		).
		OrderBy("u.pos").
		ToSql()

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

func (db *wordTranslationRepo) ListByWordID(wordID uuid.UUID) ([]*model.WordTranslation, int, error) {
	query, args, err := db.psql.
		Select("wt.id", "wt.meaning_id", "wt.language", "wt.translation").
		From("word_translations wt").
		Join("word_meanings wm ON wm.id = wt.meaning_id").
		Where(sq.Eq{"wm.word_id": wordID}).ToSql()

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
