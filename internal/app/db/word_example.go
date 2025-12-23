package db

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx"
	"github.com/nvytychakdev/vocab-mastery/internal/app/model"
)

type WordExampleRepo interface {
	ListAllByMeaningIDs(meaningIds []string) ([]*model.WordExample, int, error)
}

type wordExampleRepo struct {
	conn *pgx.Conn
	psql sq.StatementBuilderType
}

func (db *PostgresDB) WordExample() WordExampleRepo {
	return &wordExampleRepo{conn: db.conn, psql: db.psql}
}

func (db *wordExampleRepo) ListAllByMeaningIDs(meaningIds []string) ([]*model.WordExample, int, error) {
	queryBuilder := db.psql.
		Select("id", "meaning_id", "text").
		From("word_examples").
		Where("meaning_id = ANY(?)", meaningIds)

	query, args, err := queryBuilder.ToSql()

	if err != nil {
		return nil, 0, err
	}

	rows, err := db.conn.Query(query, args...)
	if err != nil {
		return nil, 0, err
	}

	defer rows.Close()

	var meanings = []*model.WordExample{}
	for rows.Next() {
		var meaning model.WordExample
		err := rows.Scan(
			&meaning.ID,
			&meaning.MeaningID,
			&meaning.Text,
		)

		if err != nil {
			return nil, 0, err
		}
		meanings = append(meanings, &meaning)
	}

	return meanings, len(meanings), rows.Err()
}
