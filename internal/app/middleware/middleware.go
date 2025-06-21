package middleware

import "github.com/nvytychakdev/vocab-mastery/internal/app/services"

type contextKey string

const USER_ID_KEY contextKey = "userId"
const SESSION_ID_KEY contextKey = "userId"
const DICTIONARY_KEY contextKey = "dictionary"
const WORD_KEY contextKey = "word"
const TRANSLATION_KEY contextKey = "translation"

type Middleware struct {
	Deps *services.Deps
}
