package services

import (
	"errors"
	"log/slog"
	"time"

	"github.com/google/uuid"
	"github.com/nvytychakdev/vocab-mastery/internal/app/db"
	"github.com/nvytychakdev/vocab-mastery/internal/app/model"
)

type FlashcardSessionService interface {
	StartOrResume(userID uuid.UUID) (*model.FlashcardSession, error)
	GetUserEngagement(userID uuid.UUID) (*model.FlashcardEngagement, error)
	CheckAnswer(meaningID uuid.UUID, answerID uuid.UUID) (bool, error)
	SubmitAnswer(sessionID uuid.UUID) error
}

type flashcardSessionService struct {
	DB db.DB
}

func NewFlashcardSessionService(db db.DB) FlashcardSessionService {
	return &flashcardSessionService{DB: db}
}

func (fcs *flashcardSessionService) StartOrResume(userID uuid.UUID) (*model.FlashcardSession, error) {
	session, err := fcs.DB.FlashcardSession().GetActiveByUserID(userID)

	if err != nil {
		return nil, err
	}

	slog.Info("Session found", "session", session)
	if session != nil && !session.IsCompleted() {
		slog.Info("Reused existing session...")
		return session, nil
	}

	sessionID, err := fcs.DB.FlashcardSession().Create(model.FlashcardSessionCreate{
		UserID:     userID,
		CardsTotal: 15,
	})

	slog.Info("New session create for some reason")

	if err != nil {
		return nil, err
	}

	return fcs.DB.FlashcardSession().GetByID(sessionID)
}

// check existing engagement, if does not exists, create one and return
func (fcs *flashcardSessionService) GetUserEngagement(userID uuid.UUID) (*model.FlashcardEngagement, error) {
	engagement, err := fcs.DB.FlashcardEngagement().GetByUserID(userID)

	if err != nil {
		slog.Error("Error due to engagement get ", "err", err)
		return nil, err
	}

	if engagement != nil {
		return engagement, nil
	}

	err = fcs.DB.FlashcardEngagement().Create(model.FlashcardEngagementCreate{
		UserID:          userID,
		LastSessionDate: time.Now(),
		LastActiveAt:    time.Now(),
		ReminderStage:   "new",
	})

	if err != nil {
		slog.Error("Create engagement failed ", "err", err)
		return nil, err
	}

	return fcs.DB.FlashcardEngagement().GetByUserID(userID)
}

func (fcs *flashcardSessionService) CheckAnswer(meaningID uuid.UUID, answerID uuid.UUID) (bool, error) {
	meaning, err := fcs.DB.WordMeaning().GetByID(meaningID)

	if err != nil {
		return false, err
	}

	translations, translationsCount, err := fcs.DB.WordTranslation().ListByWordID(meaning.WordID)
	if translationsCount == 0 {
		return false, errors.New("No translations found for the provided meaning")
	}

	for _, t := range translations {
		if t.ID == answerID {
			return true, nil
		}
	}

	return false, nil
}

func (fcs *flashcardSessionService) SubmitAnswer(sessionID uuid.UUID) error {
	return nil
}
