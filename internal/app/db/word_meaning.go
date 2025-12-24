package db

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nvytychakdev/vocab-mastery/internal/app/model"
)

type WordMeaningRepo interface {
	ListAllByWordIDs(wordIds uuid.UUIDs) ([]*model.WordMeaning, int, error)
}

type wordMeaningRepo struct {
	conn *pgxpool.Pool
	psql sq.StatementBuilderType
}

func (db *PostgresDB) WordMeaning() WordMeaningRepo {
	return &wordMeaningRepo{conn: db.conn, psql: db.psql}
}

func (db *wordMeaningRepo) ListAllByWordIDs(wordIDs uuid.UUIDs) ([]*model.WordMeaning, int, error) {
	queryBuilder := db.psql.
		Select("wm.id", "wm.word_id", "pos.code AS part_of_speech", "wm.definition").
		From("word_meanings wm").
		Join("parts_of_speech pos ON wm.part_of_speech_id = pos.id").
		Where(sq.Eq{
			"wm.word_id": wordIDs,
		}).
		OrderBy("definition DESC")

	query, args, err := queryBuilder.ToSql()

	if err != nil {
		return nil, 0, err
	}

	rows, err := db.conn.Query(context.Background(), query, args...)
	if err != nil {
		return nil, 0, err
	}

	defer rows.Close()

	var meanings = []*model.WordMeaning{}
	for rows.Next() {
		var meaning model.WordMeaning
		err := rows.Scan(
			&meaning.ID,
			&meaning.WordID,
			&meaning.PartOfSpeech,
			&meaning.Definition,
		)

		if err != nil {
			return nil, 0, err
		}
		meanings = append(meanings, &meaning)
	}

	return meanings, len(meanings), rows.Err()
}
