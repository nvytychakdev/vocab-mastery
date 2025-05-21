package auth

import (
	"errors"
	"net/http"

	"github.com/go-chi/render"
	httpError "github.com/nvytychakdev/vocab-mastery/internal/app/http-error"
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
func (auth *AuthHandler) SignOut(w http.ResponseWriter, r *http.Request) {
	var data = &SignOutRequest{}

	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, httpError.NewErrorResponse(http.StatusBadRequest, httpError.ErrInvalidPayload, err))
		return
	}

	token, claims, err := auth.Deps.AuthService.ParseToken(data.RefreshToken)
	if err != nil || !token.Valid {
		render.Render(w, r, httpError.NewErrorResponse(http.StatusUnauthorized, httpError.ErrUnauthorized, err))
		return
	}

	err = auth.Deps.DB.DeleteSessionByID(claims.SessionId)
	if err != nil {
		render.Render(w, r, httpError.NewErrorResponse(http.StatusInternalServerError, httpError.ErrInternalServer, err))
		return
	}

	var response = &SignOutResponse{
		Success: true,
	}

	render.Render(w, r, response)
}
