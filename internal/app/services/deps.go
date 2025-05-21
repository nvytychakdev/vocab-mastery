package services

import (
	"github.com/nvytychakdev/vocab-mastery/internal/app/db"
)

type Deps struct {
	DB              db.DB
	AuthService     AuthService
	PasswordService PasswordService
}
