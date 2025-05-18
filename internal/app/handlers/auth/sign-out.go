package auth

import (
	"errors"
	"net/http"

	"github.com/go-chi/render"
	"github.com/nvytychakdev/vocab-mastery/internal/app/auth"
	httpError "github.com/nvytychakdev/vocab-mastery/internal/app/http-error"
	"github.com/nvytychakdev/vocab-mastery/internal/app/model/session"
)

// Request
type SignOutRequest struct {
	RefreshToken string `json:"refreshToken"`
}

func (s *SignOutRequest) Bind(r *http.Request) error {
	if s.RefreshToken == "" {
		return errors.New("email field is required")
	}
	return nil
}

// Response
type SignOutResponse struct {
	Success bool `json:"ok"`
}

func (s *SignOutResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// Handler
func signOut(w http.ResponseWriter, r *http.Request) {
	var data = &SignOutRequest{}

	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, httpError.NewErrorResponse(err, http.StatusBadRequest))
		return
	}

	token, claims, err := auth.ParseToken(data.RefreshToken)
	if err != nil || !token.Valid {
		render.Render(w, r, httpError.NewErrorResponse(errors.New("invalid token"), http.StatusUnauthorized))
		return
	}

	err = session.DeleteSessionByID(claims.SessionId)
	if err != nil {
		render.Render(w, r, httpError.NewErrorResponse(errors.New("session does not exists"), http.StatusUnauthorized))
		return
	}

	var response = &SignOutResponse{
		Success: true,
	}

	render.Render(w, r, response)
}
