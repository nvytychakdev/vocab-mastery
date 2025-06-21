package model

import "time"

type Word struct {
	ID           string    `json:"id,omitempty"`
	DictionaryId string    `json:"dictionaryId,omitempty"`
	Word         string    `json:"word"`
	Language     string    `json:"language"`
	CreatedAt    time.Time `json:"createdAt,omitempty"`
}
