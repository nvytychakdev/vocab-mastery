package services

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/coreos/go-oidc"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/oauth2"
)

const (
	TokenTypeAccess  = "access"
	TokenTypeRefresh = "refresh"
)

type TokenClaims struct {
	Type      string    `json:"type"`
	SessionId uuid.UUID `json:"sessionId,omitempty"`
	UserId    uuid.UUID `json:"userId,omitempty"`
	jwt.RegisteredClaims
}

type AuthService interface {
	ParseToken(tokenString string) (*jwt.Token, *TokenClaims, error)
	CreateAccessToken(userId uuid.UUID) (string, int64, error)
	CreateRefreshToken(sessionId uuid.UUID, jti string) (string, int64, error)
	HandleGoogleOAuth(config *oauth2.Config, code string, claims interface{}) error
}

type authService struct {
	TokenSecret string
	GoogleOAuth struct {
		Provider *oidc.Provider
		Verifier *oidc.IDTokenVerifier
	}
}

func NewAuthService() *authService {
	ctx := context.Background()
	provider, err := oidc.NewProvider(ctx, "https://accounts.google.com")
	if err != nil {
		slog.Error("Google provider can not be created", "err", err)
	}

	verifier := provider.Verifier(&oidc.Config{ClientID: os.Getenv("GOOGLE_CLIENT_ID")})
	return &authService{TokenSecret: "Secret Phrase", GoogleOAuth: struct {
		Provider *oidc.Provider
		Verifier *oidc.IDTokenVerifier
	}{
		Provider: provider,
		Verifier: verifier,
	}}
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

func (as *authService) CreateAccessToken(userId uuid.UUID) (string, int64, error) {
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

func (as *authService) CreateRefreshToken(sessionId uuid.UUID, jti string) (string, int64, error) {
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

func (as *authService) HandleGoogleOAuth(config *oauth2.Config, code string, claims interface{}) error {
	ctx := context.Background()
	token, err := config.Exchange(ctx, code)
	if err != nil {
		return err
	}

	rawIDToken, ok := token.Extra("id_token").(string)
	if !ok {
		return err
	}

	idToken, err := as.GoogleOAuth.Verifier.Verify(ctx, rawIDToken)
	if err != nil {
		return err
	}

	if err := idToken.Claims(&claims); err != nil {
		return err
	}

	return nil
}
