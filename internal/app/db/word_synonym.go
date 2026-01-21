package db

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nvytychakdev/vocab-mastery/internal/app/model"
)

type WordSynonymRepo interface {
	ListByMeaningIDs(wordIds uuid.UUIDs) ([]*model.WordSynonym, int, error)
}

type wordSynonymRepo struct {
	conn *pgxpool.Pool
	psql sq.StatementBuilderType
}

func (db *PostgresDB) WordSynonym() WordSynonymRepo {
	return &wordSynonymRepo{conn: db.conn, psql: db.psql}
}

func (db *wordSynonymRepo) ListByMeaningIDs(wordIDs uuid.UUIDs) ([]*model.WordSynonym, int, error) {
	queryBuilder := db.psql.
		Select("w.id AS id", "ws.meaning_id", "w.word AS word", "w.created_at AS created_at").
		From("word_synonyms ws").
		Join("words w ON w.id = ws.synonym_word_id").
		Where(sq.Eq{"ws.meaning_id": wordIDs}).
		OrderBy("word DESC")

	query, args, err := queryBuilder.ToSql()

	if err != nil {
		return nil, 0, err
	}

	rows, err := db.conn.Query(context.Background(), query, args...)
	if err != nil {
		return nil, 0, err
	}

	defer rows.Close()

	var synonyms = []*model.WordSynonym{}
	for rows.Next() {
		var synonym model.WordSynonym
		err := rows.Scan(
			&synonym.ID,
			&synonym.MeaningID,
			&synonym.Word,
			&synonym.CreatedAt,
		)

		if err != nil {
			return nil, 0, err
		}
		synonyms = append(synonyms, &synonym)
	}

	return synonyms, len(synonyms), rows.Err()
}
