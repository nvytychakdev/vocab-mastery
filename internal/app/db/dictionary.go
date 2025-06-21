package db

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/nvytychakdev/vocab-mastery/internal/app/model"
)

type DictionaryRepository interface {
	CreateDictionary(userId string, name string, description string) (string, error)
	RemoveDictionaryByID(dictionaryId string) error
	GetDictionaryByID(dictionaryId string) (*model.Dictionary, error)
	GetAllDictionariesByUsedID(userId string, pagination *model.Pagination) ([]*model.Dictionary, int, error)
}

func (db *PostgresDB) CreateDictionary(userId string, name string, description string) (string, error) {
	query, args, err := db.psql.
		Insert("dictionaries").
		Columns("user_id", "name", "description").
		Values(userId, name, description).
		Suffix("RETURNING \"id\"").ToSql()

	if err != nil {
		return "", err
	}

	var dictionaryId string
	err = db.conn.QueryRow(query, args...).Scan(&dictionaryId)
	return dictionaryId, err
}

func (db *PostgresDB) RemoveDictionaryByID(dictionaryId string) error {
	query, args, err := db.psql.Delete("dictionaries").Where(sq.Eq{"id": dictionaryId}).ToSql()

	if err != nil {
		return err
	}

	_, err = db.conn.Exec(query, args...)
	return err
}

func (db *PostgresDB) GetDictionaryByID(dictionaryId string) (*model.Dictionary, error) {
	query, args, err := db.psql.
		Select("id", "user_id", "name", "description", "created_at").
		From("dictionaries").Where(sq.Eq{"id": dictionaryId}).ToSql()

	if err != nil {
		return nil, err
	}

	var dictionary model.Dictionary
	err = db.conn.QueryRow(query, args...).Scan(
		&dictionary.ID,
		&dictionary.UserID,
		&dictionary.Name,
		&dictionary.Description,
		&dictionary.CreatedAt,
	)
	return &dictionary, err
}

func (db *PostgresDB) GetAllDictionariesByUsedID(userId string, pagination *model.Pagination) ([]*model.Dictionary, int, error) {
	queryBuilder := db.psql.
		Select("id", "user_id", "name", "description", "created_at").
		From("dictionaries").Where(sq.Eq{"user_id": userId})

	if pagination != nil {
		queryBuilder = queryBuilder.Offset(uint64(pagination.Offset)).Limit(uint64(pagination.Limit))
	}

	query, args, err := queryBuilder.ToSql()

	if err != nil {
		return nil, 0, err
	}

	rows, err := db.conn.Query(query, args...)
	if err != nil {
		return nil, 0, err
	}

	defer rows.Close()

	var dictionaries = []*model.Dictionary{}
	for rows.Next() {
		var dictionary model.Dictionary
		err := rows.Scan(
			&dictionary.ID,
			&dictionary.UserID,
			&dictionary.Name,
			&dictionary.Description,
			&dictionary.CreatedAt,
		)

		if err != nil {
			return nil, 0, err
		}
		dictionaries = append(dictionaries, &dictionary)
	}

	totalQuery, totalArgs, err := db.psql.
		Select("COUNT(*)").From("dictionaries").
		Where(sq.Eq{"user_id": userId}).ToSql()

	if err != nil {
		return nil, 0, err
	}

	var total int
	err = db.conn.QueryRow(totalQuery, totalArgs...).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	return dictionaries, total, rows.Err()
}
