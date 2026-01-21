package db

import (
	"context"
	"errors"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nvytychakdev/vocab-mastery/internal/app/model"
)

type FlashcardEngagementRepo interface {
	Create(engagement model.FlashcardEngagementCreate) error
	GetByUserID(userID uuid.UUID) (*model.FlashcardEngagement, error)
	UpdateSessionDateByUserID(userID uuid.UUID, date time.Time) error
	UpdateLastActiveAtByUserID(userID uuid.UUID, activeAt time.Time) error
	UpdateDatesByUserID(userID uuid.UUID, sessionDate time.Time, activeAt time.Time) error
}

type flashcardEngagementRepo struct {
	conn *pgxpool.Pool
	psql sq.StatementBuilderType
}

func (db *PostgresDB) FlashcardEngagement() FlashcardEngagementRepo {
	return &flashcardEngagementRepo{conn: db.conn, psql: db.psql}
}

func (fc *flashcardEngagementRepo) Create(engagement model.FlashcardEngagementCreate) error {
	query, args, err := fc.psql.
		Insert("flashcard_engagement").
		Columns("user_id", "last_active_at", "reminder_stage").
		Values(
			engagement.UserID,
			engagement.LastActiveAt,
			engagement.ReminderStage,
		).ToSql()

	if err != nil {
		return err
	}

	_, err = fc.conn.Exec(context.Background(), query, args...)
	return err
}

func (fc *flashcardEngagementRepo) GetByUserID(userID uuid.UUID) (*model.FlashcardEngagement, error) {
	query, args, err := fc.psql.Select("*").From("flashcard_engagement").Where(sq.Eq{"user_id": userID}).ToSql()

	if err != nil {
		return nil, err
	}

	var engagement model.FlashcardEngagement
	err = fc.conn.QueryRow(context.Background(), query, args...).
		Scan(
			&engagement.UserID,
			&engagement.LastActiveAt,
			&engagement.LastSessionDate,
			&engagement.ReminderStage,
			&engagement.MissedDaysCount,
			&engagement.NextReminderAt,
			&engagement.CreateAt,
			&engagement.UpdatedAt,
		)

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}

	return &engagement, err
}

func (fc *flashcardEngagementRepo) UpdateSessionDateByUserID(userID uuid.UUID, date time.Time) error {
	query, args, err := fc.psql.Update("flashcard_engagement").
		SetMap(map[string]interface{}{
			"last_session_date": date,
			"updated_at":        time.Now(),
		}).
		Where(sq.Eq{"user_id": userID}).ToSql()

	if err != nil {
		return err
	}

	_, err = fc.conn.Exec(context.Background(), query, args...)
	return err
}

func (fc *flashcardEngagementRepo) UpdateLastActiveAtByUserID(userID uuid.UUID, date time.Time) error {
	query, args, err := fc.psql.Update("flashcard_engagement").
		SetMap(map[string]interface{}{
			"last_active_at": date,
			"updated_at":     time.Now(),
		}).
		Where(sq.Eq{"user_id": userID}).ToSql()

	if err != nil {
		return err
	}

	_, err = fc.conn.Exec(context.Background(), query, args...)
	return err
}

func (fc *flashcardEngagementRepo) UpdateDatesByUserID(userID uuid.UUID, sessionDate time.Time, lastActiveAt time.Time) error {
	query, args, err := fc.psql.Update("flashcard_engagement").
		SetMap(map[string]interface{}{
			"last_active_at":    lastActiveAt,
			"last_session_date": sessionDate,
			"updated_at":        time.Now(),
		}).
		Where(sq.Eq{"user_id": userID}).ToSql()

	if err != nil {
		return err
	}

	_, err = fc.conn.Exec(context.Background(), query, args...)
	return err
}
