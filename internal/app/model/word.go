package model

import (
	"time"

	"github.com/google/uuid"
)

type Word struct {
	ID        uuid.UUID `json:"id"`
	Word      string    `json:"word"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}

type WordMeaning struct {
	ID           uuid.UUID `json:"id"`
	WordID       uuid.UUID `json:"-"`
	Definition   string    `json:"definition"`
	PartOfSpeech string    `json:"partOfSpeech"`
}

type WordExample struct {
	ID        uuid.UUID `json:"id"`
	MeaningID uuid.UUID `json:"-"`
	Text      string    `json:"text"`
}

type WordSynonym struct {
	ID        uuid.UUID `json:"id"`
	Word      string    `json:"word"`
	CreatedAt time.Time `json:"-"`
	MeaningID uuid.UUID `json:"-"`
}
