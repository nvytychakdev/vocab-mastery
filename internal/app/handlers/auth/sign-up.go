package auth

import (
	"errors"
	"net/http"

	"github.com/go-chi/render"
	httpError "github.com/nvytychakdev/vocab-mastery/internal/app/http-error"
	"github.com/nvytychakdev/vocab-mastery/internal/app/model/user"
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
	ID string `json:"id"`
}

func (s *SignUpResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// Handler
func signUp(w http.ResponseWriter, r *http.Request) {
	var data = &SignUpRequest{}

	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, httpError.NewErrorResponse(err, http.StatusBadRequest))
		return
	}

	existingUser, err := user.UserExists(data.Email)
	if err != nil {
		render.Render(w, r, httpError.NewErrorResponse(err, http.StatusInternalServerError))
		return
	}

	if existingUser {
		render.Render(w, r, httpError.NewErrorResponse(errors.New("user with following email already exists"), http.StatusBadRequest))
		return
	}

	userId, err := user.CreateUser(data.Email, data.Password, data.Name)
	if err != nil {
		render.Render(w, r, httpError.NewErrorResponse(err, http.StatusInternalServerError))
		return
	}

	var response = &SignUpResponse{
		ID: userId,
	}

	render.Render(w, r, response)
}
