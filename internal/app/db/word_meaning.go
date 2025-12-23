package db

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx"
	"github.com/nvytychakdev/vocab-mastery/internal/app/model"
)

type WordMeaningRepo interface {
	ListAllByWordIDs(wordIds []string) ([]*model.WordMeaning, int, error)
}

type wordMeaningRepo struct {
	conn *pgx.Conn
	psql sq.StatementBuilderType
}

func (db *PostgresDB) WordMeaning() WordMeaningRepo {
	return &wordMeaningRepo{conn: db.conn, psql: db.psql}
}

func (db *wordMeaningRepo) ListAllByWordIDs(wordIds []string) ([]*model.WordMeaning, int, error) {
	queryBuilder := db.psql.
		Select("wm.id", "wm.word_id", "pos.code AS part_of_speech", "wm.definition").
		From("word_meanings wm").
		Where("word_id = ANY(?)", wordIds).
		Join("parts_of_speech pos ON pos.id = wm.id")

	query, args, err := queryBuilder.ToSql()

	if err != nil {
		return nil, 0, err
	}

	rows, err := db.conn.Query(query, args...)
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
