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
	ID           string `json:"id"`
	WordID       string `json:"-"`
	Definition   string `json:"definition"`
	PartOfSpeech string `json:"partOfSpeech"`
}

type WordExample struct {
	ID        uuid.UUID `json:"id"`
	MeaningID string    `json:"-"`
	Text      string    `json:"text"`
}

type WordSynonym struct {
	SynonymWordID string `json:"-"`
	MeaningID     string `json:"-"`
}
