package auth

import (
	"errors"
	"net/http"

	"github.com/go-chi/render"
	"github.com/nvytychakdev/vocab-mastery/internal/app/db"
	httpError "github.com/nvytychakdev/vocab-mastery/internal/app/http-error"
	"github.com/nvytychakdev/vocab-mastery/internal/app/model"
)

type ConfirmEmailRequest struct {
	Token string `json:"token"`
}

func (s *ConfirmEmailRequest) Bind(r *http.Request) error {
	if s.Token == "" {
		return errors.New("token is missing")
	}
	return nil
}

type ConfirmEmailResponse struct {
	AccessToken           string     `json:"accessToken"`
	AccessTokenExpiresIn  int64      `json:"accessTokenExpiresIn"`
	RefreshToken          string     `json:"refreshToken"`
	RefreshTokenExpiresIn int64      `json:"refreshTokenExpiresIn"`
	User                  SignInUser `json:"user"`
}

func (s *ConfirmEmailResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func confirmEmail(w http.ResponseWriter, r *http.Request) {
	data := &ConfirmEmailRequest{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, httpError.NewErrorResponse(http.StatusInternalServerError, httpError.ErrInvalidPayload, nil))
		return
	}

	userId, usedAt, err := db.GetNonExpiredUserToken(data.Token, model.EMAIL_CONFIRM_TOKEN)
	if err != nil {
		render.Render(w, r, httpError.NewErrorResponse(http.StatusInternalServerError, httpError.ErrConfirmTokenExpired, err))
		return
	}

	if usedAt != nil {
		render.Render(w, r, httpError.NewErrorResponse(http.StatusUnauthorized, httpError.ErrConfirmTokenAlreadyUsed, err))
		return
	}

	err = db.SetUserTokenUsed(data.Token)
	if err != nil {
		render.Render(w, r, httpError.NewErrorResponse(http.StatusInternalServerError, httpError.ErrInternalServer, err))
		return
	}

	err = db.SetUserEmailConfirmed(userId)
	if err != nil {
		render.Render(w, r, httpError.NewErrorResponse(http.StatusInternalServerError, httpError.ErrInternalServer, err))
		return
	}

	user, err := db.GetUserByID(userId)
	if err != nil {
		render.Render(w, r, httpError.NewErrorResponse(http.StatusInternalServerError, httpError.ErrInternalServer, err))
		return
	}

	signInComplete(w, r, user)
}
