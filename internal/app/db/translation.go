package db

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/nvytychakdev/vocab-mastery/internal/app/model"
)

type TranslationRepository interface {
	CreateTranslation(wordId string, word string, language string) (string, error)
	RemoveTranslationByID(translationId string) error
	GetTranslationByID(translationId string) (*model.Translation, error)
	GetAllTranslationsByWordID(wordId string) ([]*model.Translation, error)
}

func (db *PostgresDB) CreateTranslation(wordId string, word string, language string) (string, error) {
	query, args, err := db.psql.
		Insert("translations").
		Columns("word_id", "word", "language").
		Values(wordId, word, language).
		Suffix("RETURNING \"id\"").ToSql()

	if err != nil {
		return "", err
	}

	var translationId string
	err = db.conn.QueryRow(query, args...).Scan(&translationId)
	return translationId, err
}

func (db *PostgresDB) RemoveTranslationByID(translationId string) error {
	query, args, err := db.psql.Delete("translations").Where(sq.Eq{"id": translationId}).ToSql()

	if err != nil {
		return err
	}

	_, err = db.conn.Exec(query, args...)
	return err
}

func (db *PostgresDB) GetTranslationByID(translationId string) (*model.Translation, error) {
	query, args, err := db.psql.
		Select("id", "word_id", "word", "language", "created_at").
		From("translations").Where(sq.Eq{"id": translationId}).ToSql()

	if err != nil {
		return nil, err
	}

	var translation model.Translation
	err = db.conn.QueryRow(query, args...).Scan(
		&translation.ID,
		&translation.WordId,
		&translation.Word,
		&translation.Language,
		&translation.CreatedAt,
	)
	return &translation, err
}

func (db *PostgresDB) GetAllTranslationsByWordID(wordId string) ([]*model.Translation, error) {
	query, args, err := db.psql.
		Select("id", "word_id", "word", "language", "created_at").
		From("translations").Where(sq.Eq{"word_id": wordId}).ToSql()

	if err != nil {
		return nil, err
	}

	rows, err := db.conn.Query(query, args...)
	if err != nil {
		return nil, err
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
			return nil, err
		}
		words = append(words, &dictionary)
	}

	return words, rows.Err()
}
