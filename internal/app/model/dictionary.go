package model

import (
	"time"
)

type Dictionary struct {
	ID          string    `json:"id,omitempty"`
	UserID      string    `json:"-"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt,omitempty"`
}
