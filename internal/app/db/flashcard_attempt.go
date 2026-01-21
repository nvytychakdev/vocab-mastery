package db

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nvytychakdev/vocab-mastery/internal/app/model"
)

type FlashcardAttemptRepo interface {
	Create(attempt model.FlashcardAttemptCreate) (uuid.UUID, error)
	GetByID(ID uuid.UUID) (*model.FlashcardAttempt, error)
	GetAllByUserID(userID uuid.UUID) ([]*model.FlashcardAttempt, int, error)
	GetCorrectBySessionID(sessionID uuid.UUID) ([]*model.FlashcardAttempt, int, error)
}

type flashcardAttemptRepo struct {
	conn *pgxpool.Pool
	psql sq.StatementBuilderType
}

func (db *PostgresDB) FlashcardAttempt() FlashcardAttemptRepo {
	return &flashcardAttemptRepo{conn: db.conn, psql: db.psql}
}

func (fc *flashcardAttemptRepo) Create(attempt model.FlashcardAttemptCreate) (uuid.UUID, error) {
	query, args, err := fc.psql.
		Insert("flashcard_attempts").
		Columns("session_id", "meaning_id", "direction", "prompt_language", "answer_language", "is_correct", "response_time_ms").
		Values(
			attempt.SessionID,
			attempt.MeaningID,
			attempt.Direction,
			attempt.PromptLanguage,
			attempt.AnswerLanguage,
			attempt.IsCorrect,
			attempt.ResponseTimeMs,
		).
		Suffix("RETURNING \"id\"").ToSql()

	if err != nil {
		return uuid.Nil, err
	}

	var attemptId uuid.UUID
	err = fc.conn.QueryRow(context.Background(), query, args...).Scan(&attemptId)
	return attemptId, err
}

func (fc *flashcardAttemptRepo) GetByID(ID uuid.UUID) (*model.FlashcardAttempt, error) {
	query, args, err := fc.psql.Select("*").From("flashcard_attempts").Where(sq.Eq{"id": ID}).ToSql()

	if err != nil {
		return nil, err
	}

	var attempt model.FlashcardAttempt
	err = fc.conn.QueryRow(context.Background(), query, args...).
		Scan(
			&attempt.ID,
			&attempt.SessionID,
			&attempt.MeaningID,
			&attempt.Direction,
			&attempt.PromptLanguage,
			&attempt.AnswerLanguage,
			&attempt.IsCorrect,
			&attempt.ResponseTimeMs,
			&attempt.CreateAt,
		)

	if err == pgx.ErrNoRows {
		return nil, nil
	}

	return &attempt, err
}

func (fc *flashcardAttemptRepo) GetAllByUserID(userID uuid.UUID) ([]*model.FlashcardAttempt, int, error) {
	query, args, err := fc.psql.
		Select("fa.*").
		From("flashcard_attempts fa").
		Join("flashcard_sessions fs ON fs.id = fa.session_id").
		Where(sq.Eq{"fs.user_id": userID}).
		ToSql()

	if err != nil {
		return nil, 0, err
	}

	rows, err := fc.conn.Query(context.Background(), query, args...)
	if err != nil {
		return nil, 0, err
	}

	defer rows.Close()

	var attempts = []*model.FlashcardAttempt{}
	for rows.Next() {
		var attempt model.FlashcardAttempt
		err := rows.Scan(
			&attempt.ID,
			&attempt.SessionID,
			&attempt.MeaningID,
			&attempt.Direction,
			&attempt.PromptLanguage,
			&attempt.AnswerLanguage,
			&attempt.IsCorrect,
			&attempt.ResponseTimeMs,
			&attempt.CreateAt,
		)

		if err != nil {
			return nil, 0, err
		}
		attempts = append(attempts, &attempt)
	}

	return attempts, len(attempts), rows.Err()
}

func (fc *flashcardAttemptRepo) GetCorrectBySessionID(sessionID uuid.UUID) ([]*model.FlashcardAttempt, int, error) {
	query, args, err := fc.psql.
		Select("fa.*").
		From("flashcard_attempts fa").
		Where(sq.And{
			sq.Eq{"fa.session_id": sessionID},
			sq.Eq{"is_correct": true},
		}).
		ToSql()

	if err != nil {
		return nil, 0, err
	}

	rows, err := fc.conn.Query(context.Background(), query, args...)
	if err != nil {
		return nil, 0, err
	}

	defer rows.Close()

	var attempts = []*model.FlashcardAttempt{}
	for rows.Next() {
		var attempt model.FlashcardAttempt
		err := rows.Scan(
			&attempt.ID,
			&attempt.SessionID,
			&attempt.MeaningID,
			&attempt.Direction,
			&attempt.PromptLanguage,
			&attempt.AnswerLanguage,
			&attempt.IsCorrect,
			&attempt.ResponseTimeMs,
			&attempt.CreateAt,
		)

		if err != nil {
			return nil, 0, err
		}
		attempts = append(attempts, &attempt)
	}

	return attempts, len(attempts), rows.Err()
}
