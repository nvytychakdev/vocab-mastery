package db

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nvytychakdev/vocab-mastery/internal/app/model"
)

type FlashcardDayRepo interface {
	Create(day model.FlashcardDayCreate) (uuid.UUID, error)
	GetByID(ID uuid.UUID) (*model.FlashcardDay, error)
	GetByUserID(userID uuid.UUID) (*model.FlashcardDay, error)
}

type flashcardDayRepo struct {
	conn *pgxpool.Pool
	psql sq.StatementBuilderType
}

func (db *PostgresDB) FlashcardDay() FlashcardDayRepo {
	return &flashcardDayRepo{conn: db.conn, psql: db.psql}
}

func (fc *flashcardDayRepo) Create(day model.FlashcardDayCreate) (uuid.UUID, error) {
	query, args, err := fc.psql.
		Insert("flashcard_days").
		Columns("user_id", "date", "timezone").
		Values(
			day.UserID,
			day.Date,
			day.Timezone,
		).
		Suffix("RETURNING \"id\"").ToSql()

	if err != nil {
		return uuid.Nil, err
	}

	var dayID uuid.UUID
	err = fc.conn.QueryRow(context.Background(), query, args...).Scan(&dayID)
	return dayID, err
}

func (fc *flashcardDayRepo) GetByID(ID uuid.UUID) (*model.FlashcardDay, error) {
	query, args, err := fc.psql.Select("*").From("flashcard_days").Where(sq.Eq{"id": ID}).ToSql()

	if err != nil {
		return nil, err
	}

	var day model.FlashcardDay
	err = fc.conn.QueryRow(context.Background(), query, args...).
		Scan(
			&day.ID,
			&day.UserID,
			&day.Date,
			&day.Timezone,
			&day.StartedAt,
			&day.SessionsCount,
			&day.CardsAnswered,
			&day.CardsCorrect,
			&day.CreateAt,
			&day.UpdatedAt,
		)
	return &day, err
}

func (fc *flashcardDayRepo) GetByUserID(userID uuid.UUID) (*model.FlashcardDay, error) {
	query, args, err := fc.psql.Select("*").From("flashcard_days").Where(sq.Eq{"user_id": userID}).ToSql()

	if err != nil {
		return nil, err
	}

	var day model.FlashcardDay
	err = fc.conn.QueryRow(context.Background(), query, args...).
		Scan(
			&day.ID,
			&day.UserID,
			&day.Date,
			&day.Timezone,
			&day.StartedAt,
			&day.SessionsCount,
			&day.CardsAnswered,
			&day.CardsCorrect,
			&day.CreateAt,
			&day.UpdatedAt,
		)

	if err == pgx.ErrNoRows {
		return nil, nil
	}

	return &day, err
}
