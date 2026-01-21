package db

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nvytychakdev/vocab-mastery/internal/app/model"
)

type WordExampleRepo interface {
	ListByMeaningIDs(meaningIds uuid.UUIDs) ([]*model.WordExample, int, error)
}

type wordExampleRepo struct {
	conn *pgxpool.Pool
	psql sq.StatementBuilderType
}

func (db *PostgresDB) WordExample() WordExampleRepo {
	return &wordExampleRepo{conn: db.conn, psql: db.psql}
}

func (db *wordExampleRepo) ListByMeaningIDs(meaningIDs uuid.UUIDs) ([]*model.WordExample, int, error) {
	queryBuilder := db.psql.
		Select("id", "meaning_id", "text").
		From("word_examples").
		Where(sq.Eq{"meaning_id": meaningIDs}).
		OrderBy("text DESC")

	query, args, err := queryBuilder.ToSql()

	if err != nil {
		return nil, 0, err
	}

	rows, err := db.conn.Query(context.Background(), query, args...)
	if err != nil {
		return nil, 0, err
	}

	defer rows.Close()

	var examples = []*model.WordExample{}
	for rows.Next() {
		var example model.WordExample
		err := rows.Scan(
			&example.ID,
			&example.MeaningID,
			&example.Text,
		)

		if err != nil {
			return nil, 0, err
		}
		examples = append(examples, &example)
	}

	return examples, len(examples), rows.Err()
}
