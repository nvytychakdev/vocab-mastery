package model

import (
	"time"

	"github.com/google/uuid"
)

type UserData struct {
	Email      string  `json:"email"`
	Name       string  `json:"name"`
	PictureUrl *string `json:"pictureUrl"`
}

type User struct {
	UserData
	ID        uuid.UUID `json:"id,omitempty"`
	CreatedAt time.Time `json:"createdAt"`
}

type UserWithPwd struct {
	User
	Password         string `json:"-"`
	IsEmailConfirmed bool   `json:"-"`
}
