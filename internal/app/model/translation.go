package model

import "time"

type Translation struct {
	ID        string    `json:"id,omitempty"`
	WordId    string    `json:"wordId,omitempty"`
	Word      string    `json:"word"`
	Language  string    `json:"language"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}
