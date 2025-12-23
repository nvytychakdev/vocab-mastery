package model

import (
	"time"

	"github.com/google/uuid"
)

type Translation struct {
	ID        uuid.UUID `json:"id,omitempty"`
	WordId    uuid.UUID `json:"wordId,omitempty"`
	Word      string    `json:"word"`
	Language  string    `json:"language"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}
