package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	TOKEN_TYPE_ACCESS  = "access"
	TOKEN_TYPE_REFRESH = "refresh"
)

type TokenClaims struct {
	Type      string `json:"type"`
	SessionId string `json:"sessionId,omitempty"`
	UserId    string `json:"userId,omitempty"`
	jwt.RegisteredClaims
}

const tokenSecret = "secret-phrase"

func ParseToken(tokenString string) (*jwt.Token, *TokenClaims, error) {
	claims := &TokenClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(tokenSecret), nil
	})

	return token, claims, err
}

func CreateAccessToken(userId string) (string, string, error) {
	claims := &TokenClaims{
		Type:   TOKEN_TYPE_ACCESS,
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "vocab-mastery",
			Subject:   "auth",
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
		},
	}

	return createToken(claims)
}

func CreateRefreshToken(sessionId string, jti string) (string, string, error) {
	claims := &TokenClaims{
		Type:      TOKEN_TYPE_REFRESH,
		SessionId: sessionId,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "vocab-mastery",
			Subject:   "auth",
			ID:        jti,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(30 * 24 * time.Hour)),
		},
	}

	return createToken(claims)
}

func createToken(claims *TokenClaims) (string, string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(tokenSecret))
	if err != nil {
		return "", "", err
	}

	return tokenString, claims.ExpiresAt.Format(time.RFC3339), nil
}
