package db

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nvytychakdev/vocab-mastery/internal/app/model"
)

type UserWordProgressRepo interface {
	Create(progress model.UserWordProgressCreate) (uuid.UUID, error)
	GetByID(ID uuid.UUID) (*model.UserWordProgress, error)
	ListByUserID(userID uuid.UUID) ([]*model.UserWordProgress, int, error)
}

type userWordProgressRepo struct {
	conn *pgxpool.Pool
	psql sq.StatementBuilderType
}

func (db *PostgresDB) UserWordProgress() UserWordProgressRepo {
	return &userWordProgressRepo{conn: db.conn, psql: db.psql}
}

func (uwp *userWordProgressRepo) Create(progress model.UserWordProgressCreate) (uuid.UUID, error) {
	query, args, err := uwp.psql.
		Insert("user_words_progress").
		Columns(
			"user_id",
			"meaning_id",
			"status",
			"difficulty",
			"times_seen_recall",
			"times_correct_recall",
			"times_incorrect_recall",
			"next_review_at_recall",
			"times_seen_recognition",
			"times_correct_recognition",
			"times_incorrect_recognition",
			"next_review_at_recognition",
			"last_seen_at",
		).
		Values(
			progress.UserID,
			progress.MeaningID,
			progress.Status,
			progress.Difficulty,
			progress.TimesSeenRecall,
			progress.TimesCorrectRecall,
			progress.TimesIncorrectRecall,
			progress.NextReviewAtRecall,
			progress.TimesSeenRecognition,
			progress.TimesCorrectRecognition,
			progress.TimesIncorrectRecognition,
			progress.NextReviewAtRecognition,
			progress.LastSeenAt,
		).
		Suffix("RETURNING \"id\"").ToSql()

	if err != nil {
		return uuid.Nil, err
	}

	var progressID uuid.UUID
	err = uwp.conn.QueryRow(context.Background(), query, args...).Scan(&progressID)
	return progressID, err
}

func (uwp *userWordProgressRepo) GetByID(ID uuid.UUID) (*model.UserWordProgress, error) {
	query, args, err := uwp.psql.Select("*").From("user_words_progress").Where(sq.Eq{"id": ID}).ToSql()

	if err != nil {
		return nil, err
	}

	var progress model.UserWordProgress
	err = uwp.conn.QueryRow(context.Background(), query, args...).
		Scan(
			&progress.ID,
			&progress.UserID,
			&progress.MeaningID,
			&progress.Status,
			&progress.Difficulty,
			&progress.TimesSeenRecall,
			&progress.TimesCorrectRecall,
			&progress.TimesIncorrectRecall,
			&progress.NextReviewAtRecall,
			&progress.TimesSeenRecognition,
			&progress.TimesCorrectRecognition,
			&progress.TimesIncorrectRecognition,
			&progress.NextReviewAtRecognition,
			&progress.LastSeenAt,
			&progress.CreatedAt,
		)

	if err == pgx.ErrNoRows {
		return nil, nil
	}

	return &progress, err
}

func (uwp *userWordProgressRepo) ListByUserID(userID uuid.UUID) ([]*model.UserWordProgress, int, error) {
	query, args, err := uwp.psql.Select("*").From("user_words_progress").Where(sq.Eq{"user_id": userID}).ToSql()

	if err != nil {
		return nil, 0, err
	}

	rows, err := uwp.conn.Query(context.Background(), query, args...)
	if err != nil {
		return nil, 0, err
	}

	defer rows.Close()

	var progresses = []*model.UserWordProgress{}
	for rows.Next() {
		var progress model.UserWordProgress
		err := rows.Scan(
			&progress.ID,
			&progress.UserID,
			&progress.MeaningID,
			&progress.Status,
			&progress.Difficulty,
			&progress.TimesSeenRecall,
			&progress.TimesCorrectRecall,
			&progress.TimesIncorrectRecall,
			&progress.NextReviewAtRecall,
			&progress.TimesSeenRecognition,
			&progress.TimesCorrectRecognition,
			&progress.TimesIncorrectRecognition,
			&progress.NextReviewAtRecognition,
			&progress.LastSeenAt,
			&progress.CreatedAt,
		)

		if err != nil {
			return nil, 0, err
		}
		progresses = append(progresses, &progress)
	}

	return progresses, len(progresses), rows.Err()
}
