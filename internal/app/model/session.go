package model

import (
	"time"

	"github.com/google/uuid"
)

type Session struct {
	ID             uuid.UUID `json:"id,omitempty"`
	UserID         uuid.UUID `json:"userId"`
	RefreshTokenID uuid.UUID `json:"jti,omitempty"`
	ExpiresAt      time.Time `json:"expiresAt,omitempty"`
	CreatedAt      time.Time `json:"createdAt,omitempty"`
}
