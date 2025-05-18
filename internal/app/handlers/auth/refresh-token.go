package auth

import (
	"errors"
	"net/http"
	"time"

	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/nvytychakdev/vocab-mastery/internal/app/auth"
	httpError "github.com/nvytychakdev/vocab-mastery/internal/app/http-error"
	"github.com/nvytychakdev/vocab-mastery/internal/app/model/session"
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
	AccessTokenExpiresIn  string `json:"accessTokenExpiresIn"`
	RefreshToken          string `json:"refreshToken"`
	RefreshTokenExpiresIn string `json:"refreshTokenExpiresIn"`
}

func (s *RefreshTokenResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func refreshToken(w http.ResponseWriter, r *http.Request) {
	data := &RefreshTokenRequest{}

	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, httpError.NewErrorResponse(err, http.StatusBadRequest))
		return
	}

	token, claims, err := auth.ParseToken(data.RefreshToken)
	if err != nil || !token.Valid || claims.Type != auth.TOKEN_TYPE_REFRESH {
		render.Render(w, r, httpError.NewErrorResponse(errors.New("invalid token"), http.StatusUnauthorized))
		return
	}

	s, err := session.GetSessionByID(claims.SessionId)
	if err != nil || time.Now().Unix() > s.ExpiresAt.Unix() {
		render.Render(w, r, httpError.NewErrorResponse(errors.New("session expired"), http.StatusUnauthorized))
		return
	}

	if s.RefreshTokenID != claims.ID {
		render.Render(w, r, httpError.NewErrorResponse(errors.New("token was invalidated"), http.StatusUnauthorized))
		return
	}

	accessToken, accessTokenExpiresIn, err := auth.CreateAccessToken(claims.UserId)
	if err != nil {
		render.Render(w, r, httpError.NewErrorResponse(err, http.StatusInternalServerError))
		return
	}

	refreshTokenId := uuid.NewString()
	refreshToken, refreshTokenExpiresIn, err := auth.CreateRefreshToken(claims.SessionId, refreshTokenId)
	if err != nil {
		render.Render(w, r, httpError.NewErrorResponse(err, http.StatusInternalServerError))
		return
	}

	err = session.UpdateSessionJti(claims.SessionId, refreshTokenId)
	if err != nil {
		render.Render(w, r, httpError.NewErrorResponse(errors.New("session update failed"), http.StatusUnauthorized))
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
