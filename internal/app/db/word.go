package db

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx"
	"github.com/nvytychakdev/vocab-mastery/internal/app/model"
)

type WordRepo interface {
	Create(dictionaryId string, word string, language string) (string, error)
	DeleteByID(wordId string) error
	GetByID(wordId string) (*model.Word, error)
	ListByDictionaryID(dictionaryId string, opts *model.QueryOptions) ([]*model.Word, int, error)
}

type wordRepo struct {
	conn *pgx.Conn
	psql sq.StatementBuilderType
}

func (db *PostgresDB) Word() WordRepo {
	return &wordRepo{conn: db.conn, psql: db.psql}
}

func (db *wordRepo) Create(dictionaryId string, word string, language string) (string, error) {
	query, args, err := db.psql.
		Insert("words").
		Columns("dictionary_id", "word", "language").
		Values(dictionaryId, word, language).
		Suffix("RETURNING \"id\"").ToSql()

	if err != nil {
		return "", err
	}

	var wordId string
	err = db.conn.QueryRow(query, args...).Scan(&wordId)
	return wordId, err
}

func (db *wordRepo) DeleteByID(wordId string) error {
	query, args, err := db.psql.Delete("words").Where(sq.Eq{"id": wordId}).ToSql()

	if err != nil {
		return err
	}

	_, err = db.conn.Exec(query, args...)
	return err
}

func (db *wordRepo) GetByID(wordId string) (*model.Word, error) {
	query, args, err := db.psql.
		Select("id", "dictionary_id", "word", "language", "created_at").
		From("words").Where(sq.Eq{"id": wordId}).ToSql()

	if err != nil {
		return nil, err
	}

	var word model.Word
	err = db.conn.QueryRow(query, args...).Scan(
		&word.ID,
		&word.DictionaryId,
		&word.Word,
		&word.Language,
		&word.CreatedAt,
	)
	return &word, err
}

func (db *wordRepo) ListByDictionaryID(dictionaryId string, opts *model.QueryOptions) ([]*model.Word, int, error) {
	queryBuilder := db.psql.
		Select("id", "dictionary_id", "word", "language", "created_at").
		From("words").Where(sq.Eq{"dictionary_id": dictionaryId})

	query, args, err := ApplyQueryOptions(queryBuilder, opts).ToSql()

	if err != nil {
		return nil, 0, err
	}

	rows, err := db.conn.Query(query, args...)
	if err != nil {
		return nil, 0, err
	}

	defer rows.Close()

	var words = []*model.Word{}
	for rows.Next() {
		var dictionary model.Word
		err := rows.Scan(
			&dictionary.ID,
			&dictionary.DictionaryId,
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
		Select("COUNT(*)").From("words").
		Where(sq.Eq{"dictionary_id": dictionaryId}).ToSql()

	if err != nil {
		return nil, 0, err
	}

	var total int
	err = db.conn.QueryRow(totalQuery, totalArgs...).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	return words, total, rows.Err()
}
