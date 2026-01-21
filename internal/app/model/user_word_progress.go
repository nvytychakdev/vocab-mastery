package model

import (
	"time"

	"github.com/google/uuid"
)

type UserWordProgress struct {
	ID                        uuid.UUID          `json:"id"`
	MeaningID                 uuid.UUID          `json:"meaningId"`
	UserID                    uuid.UUID          `json:"userId"`
	Status                    WordProgressStatus `json:"status"`
	Difficulty                int                `json:"difficulty"`
	TimesSeenRecall           int                `json:"timesSeenRecall"`
	TimesCorrectRecall        int                `json:"timesCorrectRecall"`
	TimesIncorrectRecall      int                `json:"timesIncorrectRecall"`
	NextReviewAtRecall        time.Time          `json:"nextReviewAtRecall"`
	TimesSeenRecognition      int                `json:"timesSeenRecognition"`
	TimesCorrectRecognition   int                `json:"timesCorrectRecognition"`
	TimesIncorrectRecognition int                `json:"timesIncorrectRecognition"`
	NextReviewAtRecognition   time.Time          `json:"nextReviewAtRecognition"`
	LastSeenAt                time.Time          `json:"lastSeenAt"`
	CreatedAt                 time.Time          `json:"createdAt"`
}

type UserWordProgressCreate struct {
	MeaningID                 uuid.UUID          `json:"meaningId"`
	UserID                    uuid.UUID          `json:"userId"`
	Status                    WordProgressStatus `json:"status"`
	Difficulty                int                `json:"difficulty"`
	TimesSeenRecall           int                `json:"timesSeenRecall"`
	TimesCorrectRecall        int                `json:"timesCorrectRecall"`
	TimesIncorrectRecall      int                `json:"timesIncorrectRecall"`
	NextReviewAtRecall        time.Time          `json:"nextReviewAtRecall"`
	TimesSeenRecognition      int                `json:"timesSeenRecognition"`
	TimesCorrectRecognition   int                `json:"timesCorrectRecognition"`
	TimesIncorrectRecognition int                `json:"timesIncorrectRecognition"`
	NextReviewAtRecognition   time.Time          `json:"nextReviewAtRecognition"`
	LastSeenAt                time.Time          `json:"lastSeenAt"`
}
