package model

import (
	"time"

	"github.com/google/uuid"
)

type FlashcardEngagement struct {
	UserID          uuid.UUID  `json:"userId"`
	LastActiveAt    time.Time  `json:"lastActiveAt"`
	LastSessionDate *time.Time `json:"lastSessionDate"`
	ReminderStage   string     `json:"reminderStage"`
	MissedDaysCount int        `json:"missedDaysCount"`
	NextReminderAt  *time.Time `json:"nextReminderAt"`
	CreateAt        time.Time  `json:"createdAt"`
	UpdatedAt       time.Time  `json:"updatedAt"`
}

type FlashcardEngagementCreate struct {
	UserID          uuid.UUID
	LastSessionDate time.Time
	LastActiveAt    time.Time
	ReminderStage   string
}

type FlashcardDay struct {
	ID            uuid.UUID `json:"id"`
	UserID        uuid.UUID `json:"userId"`
	Date          time.Time `json:"date"`
	Timezone      string    `json:"timezone"`
	StartedAt     time.Time `json:"startedAt"`
	CompletedAt   time.Time `json:"completedAt"`
	SessionsCount int       `json:"sessionsCount"`
	CardsAnswered int       `json:"cardsAnswered"`
	CardsCorrect  int       `json:"cardsCorrect"`
	CreateAt      time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}

type FlashcardDayCreate struct {
	UserID   uuid.UUID
	Date     time.Time
	Timezone string
}

type FlashcardSession struct {
	ID                          uuid.UUID  `json:"id"`
	UserID                      uuid.UUID  `json:"userId"`
	StartedAt                   time.Time  `json:"startedAt"`
	EndedAt                     *time.Time `json:"endedAt"`
	CurrentMeaningID            uuid.UUID  `json:"-"`
	CurrentMeaningTranslationID uuid.UUID  `json:"-"`
	CurrentMeaningChoicesIDs    uuid.UUIDs `json:"-"`
	CardsTotal                  int        `json:"cardsTotal"`
	CardsCompleted              int        `json:"cardsCompleted"`
}

func (fcs *FlashcardSession) IsCompleted() bool {
	return fcs.CardsTotal == fcs.CardsCompleted
}

type FlashcardSessionCreate struct {
	UserID     uuid.UUID
	CardsTotal int
}

type FlashcardAttempt struct {
	ID             uuid.UUID `json:"id"`
	SessionID      uuid.UUID `json:"sessionId"`
	MeaningID      uuid.UUID `json:"meaningId"`
	Direction      string    `json:"direction"`
	PromptLanguage string    `json:"promptLanguage"`
	AnswerLanguage string    `json:"answerLanguage"`
	IsCorrect      bool      `json:"isCorrect"`
	ResponseTimeMs int       `json:"responseTimeMs"`
	CreateAt       time.Time `json:"createdAt"`
}

type FlashcardAttemptCreate struct {
	SessionID      uuid.UUID
	MeaningID      uuid.UUID
	Direction      string
	PromptLanguage string
	AnswerLanguage string
	IsCorrect      bool
	ResponseTimeMs int
}

type FlashcardCard struct {
	MeaningID uuid.UUID             `json:"meaningId"`
	WordID    uuid.UUID             `json:"wordId"`
	Word      string                `json:"word"`
	Meaning   string                `json:"meaning"`
	Type      string                `json:"type"`
	Choices   []FlashcardCardChoice `json:"choices"`
}

type FlashcardCardChoice struct {
	TranslationID uuid.UUID `json:"translationId"`
	Translation   string    `json:"translation"`
}

type FlashcardAnswerResult struct {
	IsCorrect      bool      `json:"isCorrect"`
	SelectedAnswer uuid.UUID `json:"selectedAnswer"`
	CorrectAnswer  uuid.UUID `json:"correctAnswer"`
}
