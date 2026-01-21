package db

import (
	"context"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nvytychakdev/vocab-mastery/internal/app/model"
)

type FlashcardSessionRepo interface {
	Create(session model.FlashcardSessionCreate) (uuid.UUID, error)
	UpdateCurrentAnswer(sessionID uuid.UUID, meaningID *uuid.UUID, answerID *uuid.UUID, choicesIDs uuid.UUIDs) error
	UpdateCompletedCounter(sessionID uuid.UUID, completed int) error
	UpdateEndedAt(sessionID uuid.UUID, endedAt time.Time) error
	GetByID(ID uuid.UUID) (*model.FlashcardSession, error)
	GetByUserID(userID uuid.UUID) (*model.FlashcardSession, error)
	GetActiveByUserID(userID uuid.UUID) (*model.FlashcardSession, error)
	GetRandomMeaningToLearn(userID uuid.UUID, dictionaryID uuid.UUID, sessionID uuid.UUID) (uuid.UUID, float32, error)
	ListRandomAnswers(excludeMeaningID uuid.UUID) ([]*model.WordTranslation, int, error)
	GetRandomAnswerByWordID(wordId uuid.UUID) (*model.WordTranslation, error)
}

type flashcardSessionRepo struct {
	conn *pgxpool.Pool
	psql sq.StatementBuilderType
}

func (db *PostgresDB) FlashcardSession() FlashcardSessionRepo {
	return &flashcardSessionRepo{conn: db.conn, psql: db.psql}
}

func (fc *flashcardSessionRepo) Create(session model.FlashcardSessionCreate) (uuid.UUID, error) {
	query, args, err := fc.psql.
		Insert("flashcard_sessions").
		Columns("user_id", "cards_total").
		Values(
			session.UserID,
			session.CardsTotal,
		).
		Suffix("RETURNING \"id\"").ToSql()

	if err != nil {
		return uuid.Nil, err
	}

	var sessionID uuid.UUID
	err = fc.conn.QueryRow(context.Background(), query, args...).Scan(&sessionID)
	return sessionID, err
}

func (fc *flashcardSessionRepo) UpdateCurrentAnswer(sessionID uuid.UUID, meaningID *uuid.UUID, translationID *uuid.UUID, choicesIDs uuid.UUIDs) error {
	query, args, err := fc.psql.Update("flashcard_sessions").
		SetMap(map[string]interface{}{
			"current_meaning_id":             meaningID,
			"current_meaning_translation_id": translationID,
			"current_meaning_choices_ids":    choicesIDs,
		}).
		Where(sq.Eq{"id": sessionID}).ToSql()

	if err != nil {
		return err
	}

	_, err = fc.conn.Exec(context.Background(), query, args...)
	return err
}

func (fc *flashcardSessionRepo) UpdateCompletedCounter(sessionID uuid.UUID, completed int) error {
	query, args, err := fc.psql.Update("flashcard_sessions").
		SetMap(map[string]interface{}{
			"cards_completed": completed,
		}).
		Where(sq.Eq{"id": sessionID}).ToSql()

	if err != nil {
		return err
	}

	_, err = fc.conn.Exec(context.Background(), query, args...)
	return err
}

func (fc *flashcardSessionRepo) UpdateEndedAt(sessionID uuid.UUID, endedAt time.Time) error {
	query, args, err := fc.psql.Update("flashcard_sessions").
		SetMap(map[string]interface{}{
			"ended_at": endedAt,
		}).
		Where(sq.Eq{"id": sessionID}).ToSql()

	if err != nil {
		return err
	}

	_, err = fc.conn.Exec(context.Background(), query, args...)
	return err
}

func (fc *flashcardSessionRepo) GetByID(ID uuid.UUID) (*model.FlashcardSession, error) {
	query, args, err := fc.psql.Select("*").From("flashcard_sessions").Where(sq.Eq{"id": ID}).ToSql()

	if err != nil {
		return nil, err
	}

	var session model.FlashcardSession
	err = fc.conn.QueryRow(context.Background(), query, args...).
		Scan(
			&session.ID,
			&session.UserID,
			&session.StartedAt,
			&session.EndedAt,
			&session.CurrentMeaningID,
			&session.CurrentMeaningTranslationID,
			&session.CurrentMeaningChoicesIDs,
			&session.CardsTotal,
			&session.CardsCompleted,
		)

	if err == pgx.ErrNoRows {
		return nil, nil
	}

	return &session, err
}

func (fc *flashcardSessionRepo) GetByUserID(userID uuid.UUID) (*model.FlashcardSession, error) {
	query, args, err := fc.psql.Select("*").From("flashcard_sessions").Where(sq.Eq{"user_id": userID}).ToSql()

	if err != nil {
		return nil, err
	}

	var session model.FlashcardSession
	err = fc.conn.QueryRow(context.Background(), query, args...).
		Scan(
			&session.ID,
			&session.UserID,
			&session.StartedAt,
			&session.EndedAt,
			&session.CurrentMeaningID,
			&session.CurrentMeaningTranslationID,
			&session.CurrentMeaningChoicesIDs,
			&session.CardsTotal,
			&session.CardsCompleted,
		)

	if err == pgx.ErrNoRows {
		return nil, nil
	}

	return &session, err
}

func (fc *flashcardSessionRepo) GetActiveByUserID(userID uuid.UUID) (*model.FlashcardSession, error) {
	query, args, err := fc.psql.Select("*").From("flashcard_sessions").
		Where(sq.Eq{"user_id": userID}).
		Where("ended_at IS NULL").
		ToSql()

	if err != nil {
		return nil, err
	}

	var session model.FlashcardSession
	err = fc.conn.QueryRow(context.Background(), query, args...).
		Scan(
			&session.ID,
			&session.UserID,
			&session.StartedAt,
			&session.EndedAt,
			&session.CurrentMeaningID,
			&session.CurrentMeaningTranslationID,
			&session.CurrentMeaningChoicesIDs,
			&session.CardsTotal,
			&session.CardsCompleted,
		)

	if err == pgx.ErrNoRows {
		return nil, nil
	}

	return &session, err
}

func (fc *flashcardSessionRepo) GetRandomMeaningToLearn(userID uuid.UUID, dictionaryID uuid.UUID, sessionID uuid.UUID) (uuid.UUID, float32, error) {
	query, args, err := fc.psql.
		Select(
			"m.id AS meaning_id",
			`
			(
				CASE
					WHEN uwp.status = 'learning' THEN 60
					WHEN uwp.status = 'new'      THEN 20
					WHEN uwp.status = 'review'   THEN 10
					ELSE 10
				END
				+ random() * 10
			) AS score
		`,
		).
		From("word_meanings m").
		LeftJoin(
			"user_words_progress uwp ON uwp.meaning_id = m.id AND uwp.user_id = ?",
			userID,
		).
		// Where(sq.Eq{
		// 	"m.dictionary_id": dictionaryID,
		// }).
		// Where(sq.LtOrEq{
		// 	"m.level": maxLevel,
		// }).
		Where(
			sq.Expr(`
			m.id NOT IN (
				SELECT meaning_id
				FROM flashcard_attempts
				WHERE session_id = ?
			)
		`, sessionID),
		).
		OrderBy("score DESC").
		Limit(1).ToSql()

	if err != nil {
		return uuid.Nil, 0, err
	}

	var meaningID uuid.UUID
	var score float32
	err = fc.conn.QueryRow(context.Background(), query, args...).Scan(&meaningID, &score)
	return meaningID, score, err
}

func (fc *flashcardSessionRepo) ListRandomAnswers(excludeMeaningID uuid.UUID) ([]*model.WordTranslation, int, error) {
	query, args, err := fc.psql.
		Select("DISTINCT ON (meaning_id) *").
		From("word_translations").
		Where(sq.NotEq{"meaning_id": excludeMeaningID}).
		OrderBy("meaning_id, RANDOM()").
		Limit(3).ToSql()

	rows, err := fc.conn.Query(context.Background(), query, args...)
	if err != nil {
		return nil, 0, err
	}

	defer rows.Close()

	translations := []*model.WordTranslation{}
	for rows.Next() {
		var t model.WordTranslation
		rows.Scan(
			&t.ID,
			&t.MeaningID,
			&t.Language,
			&t.Translation,
		)
		translations = append(translations, &t)
	}

	return translations, len(translations), nil
}

func (fc *flashcardSessionRepo) GetRandomAnswerByWordID(wordId uuid.UUID) (*model.WordTranslation, error) {
	query, args, err := fc.psql.
		Select("wt.*").
		From("word_translations wt").
		Join("word_meanings wm ON wm.id = wt.meaning_id").
		Where(sq.Eq{"wm.word_id": wordId}).
		OrderBy("RANDOM()").
		Limit(1).ToSql()

	var translation model.WordTranslation
	err = fc.conn.QueryRow(context.Background(), query, args...).
		Scan(
			&translation.ID,
			&translation.MeaningID,
			&translation.Language,
			&translation.Translation,
		)
	return &translation, err
}
