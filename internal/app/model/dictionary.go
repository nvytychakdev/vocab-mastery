package model

import (
	"time"
)

type Dictionary struct {
	ID        string    `json:"id,omitempty"`
	UserID    string    `json:"-"`
	Title     string    `json:"title"`
	Level     string    `json:"level"`
	IsDefault bool      `json:"isDefault"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}

type DictionaryWord struct {
	DictionaryID string    `json:"dictionaryId,omitempty"`
	WordId       string    `json:"wordId,omitempty"`
	AddedAt      time.Time `json:"addedAt,omitempty"`
}
