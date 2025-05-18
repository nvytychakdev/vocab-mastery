package auth

import (
	"errors"
	"net/http"

	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/nvytychakdev/vocab-mastery/internal/app/auth"
	httpError "github.com/nvytychakdev/vocab-mastery/internal/app/http-error"
	"github.com/nvytychakdev/vocab-mastery/internal/app/model/session"
	"github.com/nvytychakdev/vocab-mastery/internal/app/model/user"
	"github.com/nvytychakdev/vocab-mastery/internal/app/utils"
)

// Request
type SignInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (s *SignInRequest) Bind(r *http.Request) error {
	if s.Email == "" {
		return errors.New("email field is required")
	}

	if s.Password == "" {
		return errors.New("password field is requried")
	}

	return nil
}

// Response
type SignInResponse struct {
	AccessToken           string     `json:"accessToken"`
	AccessTokenExpiresIn  string     `json:"accessTokenExpiresIn"`
	RefreshToken          string     `json:"refreshToken"`
	RefreshTokenExpiresIn string     `json:"refreshTokenExpiresIn"`
	User                  SignInUser `json:"user"`
}

type SignInUser struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

func (s *SignInResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// Handler
func signIn(w http.ResponseWriter, r *http.Request) {
	data := &SignInRequest{}

	if err := render.Bind(r, data); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	userExists, err := user.UserExists(data.Email)
	if err != nil {
		render.Render(w, r, httpError.NewErrorResponse(err, http.StatusInternalServerError))
		return
	}

	if !userExists {
		render.Render(w, r, httpError.NewErrorResponse(errors.New("user does not exists"), http.StatusUnauthorized))
		return
	}

	user, err := user.GetUserWithPawdByEmail(data.Email)
	if err != nil {
		render.Render(w, r, httpError.NewErrorResponse(err, http.StatusInternalServerError))
		return
	}

	passwordMatch := utils.ComparePassword(user.Password, data.Password)
	if !passwordMatch {
		render.Render(w, r, httpError.NewErrorResponse(errors.New("password does not match"), http.StatusUnauthorized))
		return
	}

	accessToken, accessTokenExpiresIn, err := auth.CreateAccessToken(user.ID)
	if err != nil {
		render.Render(w, r, httpError.NewErrorResponse(err, http.StatusInternalServerError))
		return
	}

	refreshTokenId := uuid.NewString()
	sessionId, err := session.CraeteSession(user.ID, refreshTokenId)
	if err != nil {
		render.Render(w, r, httpError.NewErrorResponse(err, http.StatusInternalServerError))
		return
	}

	refreshToken, refreshTokenExpiresIn, err := auth.CreateRefreshToken(sessionId, refreshTokenId)
	if err != nil {
		render.Render(w, r, httpError.NewErrorResponse(err, http.StatusInternalServerError))
		return
	}

	response := &SignInResponse{
		AccessToken:           accessToken,
		AccessTokenExpiresIn:  accessTokenExpiresIn,
		RefreshToken:          refreshToken,
		RefreshTokenExpiresIn: refreshTokenExpiresIn,
		User: SignInUser{
			ID:    user.ID,
			Email: user.Email,
		},
	}

	render.Render(w, r, response)
}
