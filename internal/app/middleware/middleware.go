package middleware

import "github.com/nvytychakdev/vocab-mastery/internal/app/services"

type contextKey string

const WORD_KEY contextKey = "word"
const USER_ID_KEY contextKey = "userId"
const SESSION_ID_KEY contextKey = "userId"
const DICTIONARY_KEY contextKey = "dictionary"
const TRANSLATION_KEY contextKey = "translation"
const INCLUDE_KEY contextKey = "include"
const PAGINATION_KEY contextKey = "pagination"

type Middleware struct {
	Deps *services.Deps
}

func NewMiddleware(deps *services.Deps) *Middleware {
	return &Middleware{Deps: deps}
}
