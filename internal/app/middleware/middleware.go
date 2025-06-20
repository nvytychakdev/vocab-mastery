package middleware

import "github.com/nvytychakdev/vocab-mastery/internal/app/services"

type contextKey string

const USER_ID_KEY contextKey = "userId"
const SESSION_ID_KEY contextKey = "userId"
const DICTIONARY_KEY contextKey = "dictionary"

type Middleware struct {
	Deps *services.Deps
}
