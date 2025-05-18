package model

import (
	"time"
)

type UserData struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

type User struct {
	UserData
	ID        string    `json:"id,omitempty"`
	CreatedAt time.Time `json:"createdAt"`
}

type UserWithPwd struct {
	User
	Password string `json:"-"`
}
