package auth

import (
	"errors"
	"net/http"
	"time"

	"github.com/go-chi/render"
	"github.com/google/uuid"
	httpError "github.com/nvytychakdev/vocab-mastery/internal/app/http-error"
	"github.com/nvytychakdev/vocab-mastery/internal/app/services"
)

type RefreshTokenRequest struct {
	RefreshToken string `json:"refreshToken"`
}

func (rt *RefreshTokenRequest) Bind(r *http.Request) error {
	if rt.RefreshToken == "" {
		return errors.New("email field is required")
	}

	return nil
}

type RefreshTokenResponse struct {
	AccessToken           string `json:"accessToken"`
	AccessTokenExpiresIn  int64  `json:"accessTokenExpiresIn"`
	RefreshToken          string `json:"refreshToken"`
	RefreshTokenExpiresIn int64  `json:"refreshTokenExpiresIn"`
}

func (s *RefreshTokenResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (auth *AuthHandler) RefreshToken(w http.ResponseWriter, r *http.Request) {
	data := &RefreshTokenRequest{}

	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, httpError.NewErrorResponse(http.StatusBadRequest, httpError.ErrInvalidPayload, err))
		return
	}

	token, claims, err := auth.Deps.TokenService.ParseToken(data.RefreshToken)
	if err != nil || !token.Valid || claims.Type != services.TokenTypeRefresh {
		render.Render(w, r, httpError.NewErrorResponse(http.StatusUnauthorized, httpError.ErrInvalidToken, err))
		return
	}

	s, err := auth.Deps.DB.GetSessionByID(claims.SessionId)
	if err != nil || time.Now().Unix() > s.ExpiresAt.Unix() {
		render.Render(w, r, httpError.NewErrorResponse(http.StatusUnauthorized, httpError.ErrSessionExpired, err))
		return
	}

	if s.RefreshTokenID != claims.ID {
		render.Render(w, r, httpError.NewErrorResponse(http.StatusUnauthorized, httpError.ErrTokenRevoked, nil))
		return
	}

	accessToken, accessTokenExpiresIn, err := auth.Deps.TokenService.CreateAccessToken(s.UserID)
	if err != nil {
		render.Render(w, r, httpError.NewErrorResponse(http.StatusInternalServerError, httpError.ErrInternalServer, err))
		return
	}

	refreshTokenId := uuid.NewString()
	refreshToken, refreshTokenExpiresIn, err := auth.Deps.TokenService.CreateRefreshToken(claims.SessionId, refreshTokenId)
	if err != nil {
		render.Render(w, r, httpError.NewErrorResponse(http.StatusInternalServerError, httpError.ErrInternalServer, err))
		return
	}

	err = auth.Deps.DB.UpdateSessionJti(claims.SessionId, refreshTokenId)
	if err != nil {
		render.Render(w, r, httpError.NewErrorResponse(http.StatusInternalServerError, httpError.ErrInternalServer, err))
		return
	}

	response := &RefreshTokenResponse{
		AccessToken:           accessToken,
		AccessTokenExpiresIn:  accessTokenExpiresIn,
		RefreshToken:          refreshToken,
		RefreshTokenExpiresIn: refreshTokenExpiresIn,
	}

	render.Render(w, r, response)
}
