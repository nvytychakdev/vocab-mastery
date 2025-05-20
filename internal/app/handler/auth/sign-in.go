package auth

import (
	"errors"
	"net/http"

	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/nvytychakdev/vocab-mastery/internal/app/auth"
	"github.com/nvytychakdev/vocab-mastery/internal/app/db"
	httpError "github.com/nvytychakdev/vocab-mastery/internal/app/http-error"
	"github.com/nvytychakdev/vocab-mastery/internal/app/model"
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

type SignInUser struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

// Response
type SignInResponse struct {
	AccessToken           string     `json:"accessToken"`
	AccessTokenExpiresIn  int64      `json:"accessTokenExpiresIn"`
	RefreshToken          string     `json:"refreshToken"`
	RefreshTokenExpiresIn int64      `json:"refreshTokenExpiresIn"`
	User                  SignInUser `json:"user"`
}

func (s *SignInResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

type EmailConfirmResponse struct {
	Sent bool `json:"sent"`
}

func (e *EmailConfirmResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// Handler
func SignIn(w http.ResponseWriter, r *http.Request) {
	user, data := verifyUser(w, r)

	if !user.IsEmailConfirmed {
		sendEmailConfirm(w, r, user.ID, user.Email)
		return
	}

	signInAuthorize(w, r, user, data)
}

func verifyUser(w http.ResponseWriter, r *http.Request) (*model.UserWithPwd, *SignInRequest) {
	data := &SignInRequest{}

	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, httpError.NewErrorResponse(http.StatusBadRequest, httpError.ErrInvalidPayload, err))
		return nil, nil
	}

	userExists, err := db.Instance.UserExists(data.Email)
	if err != nil {
		render.Render(w, r, httpError.NewErrorResponse(http.StatusInternalServerError, httpError.ErrInternalServer, err))
		return nil, nil
	}

	if !userExists {
		render.Render(w, r, httpError.NewErrorResponse(http.StatusUnauthorized, httpError.ErrUserNotFound, nil))
		return nil, nil
	}

	user, err := db.Instance.GetUserWithPawdByEmail(data.Email)
	if err != nil {
		render.Render(w, r, httpError.NewErrorResponse(http.StatusInternalServerError, httpError.ErrInternalServer, err))
		return nil, nil
	}

	return user, data
}

func signInAuthorize(w http.ResponseWriter, r *http.Request, user *model.UserWithPwd, data *SignInRequest) {
	passwordMatch := utils.ComparePassword(user.Password, data.Password)
	if !passwordMatch {
		render.Render(w, r, httpError.NewErrorResponse(http.StatusUnauthorized, httpError.ErrPasswordMismatch, nil))
		return
	}

	signInComplete(w, r, &user.User)
}

func signInComplete(w http.ResponseWriter, r *http.Request, user *model.User) {
	accessToken, accessTokenExpiresIn, err := auth.TokenService.CreateAccessToken(user.ID)
	if err != nil {
		render.Render(w, r, httpError.NewErrorResponse(http.StatusInternalServerError, httpError.ErrInternalServer, err))
		return
	}

	refreshTokenId := uuid.NewString()
	sessionId, err := db.Instance.CreateSession(user.ID, refreshTokenId)
	if err != nil {
		render.Render(w, r, httpError.NewErrorResponse(http.StatusInternalServerError, httpError.ErrInternalServer, err))
		return
	}

	refreshToken, refreshTokenExpiresIn, err := auth.TokenService.CreateRefreshToken(sessionId, refreshTokenId)
	if err != nil {
		render.Render(w, r, httpError.NewErrorResponse(http.StatusInternalServerError, httpError.ErrInternalServer, err))
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
