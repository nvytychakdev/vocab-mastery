package model

import "time"

const (
	EMAIL_CONFIRM_TOKEN  = "email_confirm"
	RESET_PASSWORD_TOKEN = "password_reset"
)

type UserToken struct {
	ID        string    `json:"id,omitempty"`
	Token     string    `json:"token"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"createdAt"`
	ExpiresAt time.Time `json:"expiresAt"`
	UsedAt    time.Time `json:"usedAt"`
}
