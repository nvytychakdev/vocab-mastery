package db

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jackc/pgx"
	"github.com/nvytychakdev/vocab-mastery/internal/app/model"
)

type WordRepo interface {
	Create(dictionaryId uuid.UUID, word string, language string) (string, error)
	DeleteByID(wordId uuid.UUID) error
	GetByID(wordId uuid.UUID) (*model.Word, error)
	ListByDictionaryID(dictionaryId uuid.UUID, opts *model.QueryOptions) ([]*model.Word, int, error)
	ListAll(userId uuid.UUID, opts *model.QueryOptions) ([]*model.Word, int, error)
}

type wordRepo struct {
	conn *pgx.Conn
	psql sq.StatementBuilderType
}

func (db *PostgresDB) Word() WordRepo {
	return &wordRepo{conn: db.conn, psql: db.psql}
}

func (db *wordRepo) Create(dictionaryId uuid.UUID, word string, language string) (string, error) {
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

func (db *wordRepo) DeleteByID(wordId uuid.UUID) error {
	query, args, err := db.psql.Delete("words").Where(sq.Eq{"id": wordId}).ToSql()

	if err != nil {
		return err
	}

	_, err = db.conn.Exec(query, args...)
	return err
}

func (db *wordRepo) GetByID(wordId uuid.UUID) (*model.Word, error) {
	query, args, err := db.psql.
		Select("id", "word", "created_at").
		From("words").Where(sq.Eq{"id": wordId}).ToSql()

	if err != nil {
		return nil, err
	}

	var word model.Word
	err = db.conn.QueryRow(query, args...).Scan(
		&word.ID,
		&word.Word,
		&word.CreatedAt,
	)
	return &word, err
}

func (db *wordRepo) ListByDictionaryID(dictionaryId uuid.UUID, opts *model.QueryOptions) ([]*model.Word, int, error) {
	queryBuilder := db.psql.
		Select("id", "word", "created_at").
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
		var word model.Word
		err := rows.Scan(
			&word.ID,
			&word.Word,
			&word.CreatedAt,
		)

		if err != nil {
			return nil, 0, err
		}
		words = append(words, &word)
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

func listAllQuery(builder sq.SelectBuilder, userId uuid.UUID) sq.SelectBuilder {
	return builder.From("dictionaries d").
		Join("dictionary_words dw ON dw.dictionary_id = d.id").
		Join("words w ON dw.word_id = w.id").
		Where(sq.Or{
			sq.Eq{"d.owner_id": userId},
			sq.Eq{"d.is_default": true},
		})
}

func (db *wordRepo) ListAll(userId uuid.UUID, opts *model.QueryOptions) ([]*model.Word, int, error) {
	baseBuilder := db.psql.Select("w.id", "w.word", "w.created_at")
	queryBuilder := listAllQuery(baseBuilder, userId)
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
		var word model.Word
		err := rows.Scan(
			&word.ID,
			&word.Word,
			&word.CreatedAt,
		)

		if err != nil {
			return nil, 0, err
		}
		words = append(words, &word)
	}

	baseCountBuilder := db.psql.Select("COUNT(*)")
	totalQuery, totalArgs, err := listAllQuery(baseCountBuilder, userId).ToSql()

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
