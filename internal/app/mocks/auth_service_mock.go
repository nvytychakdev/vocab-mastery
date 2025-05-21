package mocks

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/nvytychakdev/vocab-mastery/internal/app/services"
)

type MockAuthService struct {
	ParseTokenFn func(token string) (*jwt.Token, *services.TokenClaims, error)
}

func (m *MockAuthService) ParseToken(token string) (*jwt.Token, *services.TokenClaims, error) {
	return m.ParseTokenFn(token)
}

func (m *MockAuthService) CreateAccessToken(userId string) (string, int64, error) { return "", 0, nil }

func (m *MockAuthService) CreateRefreshToken(sessionId string, jti string) (string, int64, error) {
	return "", 0, nil
}
