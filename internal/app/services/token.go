package services

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	TokenTypeAccess  = "access"
	TokenTypeRefresh = "refresh"
)

type TokenClaims struct {
	Type      string `json:"type"`
	SessionId string `json:"sessionId,omitempty"`
	UserId    string `json:"userId,omitempty"`
	jwt.RegisteredClaims
}

type AuthService interface {
	ParseToken(tokenString string) (*jwt.Token, *TokenClaims, error)
	CreateAccessToken(userId string) (string, int64, error)
	CreateRefreshToken(sessionId string, jti string) (string, int64, error)
}

type authService struct {
	TokenSecret string
}

func NewAuthService() *authService {
	return &authService{TokenSecret: "Secret Phrase"}
}

func (as *authService) ParseToken(tokenString string) (*jwt.Token, *TokenClaims, error) {
	claims := &TokenClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(as.TokenSecret), nil
	})

	return token, claims, err
}

func (as *authService) CreateAccessToken(userId string) (string, int64, error) {
	claims := &TokenClaims{
		Type:   TokenTypeAccess,
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "vocab-mastery",
			Subject:   "auth",
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
		},
	}

	return createToken(as.TokenSecret, claims)
}

func (as *authService) CreateRefreshToken(sessionId string, jti string) (string, int64, error) {
	claims := &TokenClaims{
		Type:      TokenTypeRefresh,
		SessionId: sessionId,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "vocab-mastery",
			Subject:   "auth",
			ID:        jti,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(30 * 24 * time.Hour)),
		},
	}

	return createToken(as.TokenSecret, claims)
}

func createToken(secret string, claims *TokenClaims) (string, int64, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", 0, err
	}

	expiresIn := claims.ExpiresAt.Unix() - time.Now().Unix()
	return tokenString, expiresIn, nil
}
