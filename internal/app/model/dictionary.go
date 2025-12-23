package model

import (
	"time"

	"github.com/google/uuid"
)

type Dictionary struct {
	ID        uuid.UUID `json:"id,omitempty"`
	OwnerID   *string   `json:"-"`
	Title     string    `json:"title"`
	Level     *string   `json:"level"`
	IsDefault bool      `json:"isDefault"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}

type DictionaryWord struct {
	DictionaryID uuid.UUID `json:"dictionaryId,omitempty"`
	WordId       uuid.UUID `json:"wordId,omitempty"`
	AddedAt      time.Time `json:"addedAt,omitempty"`
}
