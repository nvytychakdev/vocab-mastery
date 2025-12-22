package model

import "time"

type Word struct {
	ID        string    `json:"id,omitempty"`
	Word      string    `json:"word"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}
