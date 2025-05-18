package model

import (
	"time"
)

type Session struct {
	ID             string    `json:"id,omitempty"`
	UserID         string    `json:"userId"`
	RefreshTokenID string    `json:"jti,omitempty"`
	ExpiresAt      time.Time `json:"expiresAt,omitempty"`
	CreatedAt      time.Time `json:"createdAt,omitempty"`
}
