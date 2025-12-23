package model

import (
	"time"

	"github.com/google/uuid"
)

const (
	EMAIL_CONFIRM_TOKEN  = "email_confirm"
	RESET_PASSWORD_TOKEN = "password_reset"
)

type UserToken struct {
	ID        uuid.UUID `json:"id,omitempty"`
	Token     uuid.UUID `json:"token"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"createdAt"`
	ExpiresAt time.Time `json:"expiresAt"`
	UsedAt    time.Time `json:"usedAt"`
}
