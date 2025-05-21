package auth

import (
	"errors"
	"net/http"

	"github.com/go-chi/render"
	httpError "github.com/nvytychakdev/vocab-mastery/internal/app/http-error"
)

// Request
type SignUpRequest struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

func (s *SignUpRequest) Bind(r *http.Request) error {
	if s.Email == "" {
		return errors.New("email field is required")
	}

	if s.Name == "" {
		return errors.New("name field is required")
	}

	if s.Password == "" {
		return errors.New("password field is required")
	}
	return nil
}

// Response
type SignUpResponse struct {
	ID             string `json:"id"`
	EmailConfirmed bool   `json:"emailConfirmed"`
}

func (s *SignUpResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// Handler
func (auth *AuthHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	var data = &SignUpRequest{}

	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, httpError.NewErrorResponse(http.StatusBadRequest, httpError.ErrInvalidPayload, err))
		return
	}

	existingUser, err := auth.Deps.DB.UserExists(data.Email)
	if err != nil {
		render.Render(w, r, httpError.NewErrorResponse(http.StatusInternalServerError, httpError.ErrInternalServer, err))
		return
	}

	if existingUser {
		render.Render(w, r, httpError.NewErrorResponse(http.StatusConflict, httpError.ErrUserAlreadyExists, nil))
		return
	}

	userId, err := auth.Deps.DB.CreateUser(data.Email, data.Password, data.Name)
	if err != nil {
		render.Render(w, r, httpError.NewErrorResponse(http.StatusInternalServerError, httpError.ErrInternalServer, err))
		return
	}

	auth.sendEmailConfirm(w, r, userId, data.Email)
}
