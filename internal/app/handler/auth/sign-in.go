package auth

import (
	"errors"
	"net/http"

	"github.com/go-chi/render"
	"github.com/google/uuid"
	httpError "github.com/nvytychakdev/vocab-mastery/internal/app/http-error"
	"github.com/nvytychakdev/vocab-mastery/internal/app/model"
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

type SignInTokens struct {
	AccessToken           string `json:"accessToken"`
	AccessTokenExpiresIn  int64  `json:"accessTokenExpiresIn"`
	RefreshToken          string `json:"refreshToken"`
	RefreshTokenExpiresIn int64  `json:"refreshTokenExpiresIn"`
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
func (auth *AuthHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	user, data := auth.verifyUser(w, r)

	if user == nil || data == nil {
		return
	}

	if !user.IsEmailConfirmed {
		auth.sendEmailConfirm(w, r, user.ID, user.Email)
		return
	}

	auth.signInAuthorize(w, r, user, data)
}

func (auth *AuthHandler) verifyUser(w http.ResponseWriter, r *http.Request) (*model.UserWithPwd, *SignInRequest) {
	data := &SignInRequest{}

	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, httpError.NewErrorResponse(http.StatusBadRequest, httpError.ErrInvalidPayload, err))
		return nil, nil
	}

	userExists, err := auth.Deps.DB.User().ExistsByProvider(data.Email, "local")
	if err != nil {
		render.Render(w, r, httpError.NewErrorResponse(http.StatusInternalServerError, httpError.ErrInternalServer, err))
		return nil, nil
	}

	if !userExists {
		render.Render(w, r, httpError.NewErrorResponse(http.StatusUnauthorized, httpError.ErrUserNotFound, nil))
		return nil, nil
	}

	user, err := auth.Deps.DB.User().GetByEmailWithPwd(data.Email)
	if err != nil {
		render.Render(w, r, httpError.NewErrorResponse(http.StatusInternalServerError, httpError.ErrInternalServer, err))
		return nil, nil
	}

	return user, data
}

func (auth *AuthHandler) signInAuthorize(w http.ResponseWriter, r *http.Request, user *model.UserWithPwd, data *SignInRequest) {
	passwordMatch := auth.Deps.PasswordService.ComparePassword(user.Password, data.Password)
	if !passwordMatch {
		render.Render(w, r, httpError.NewErrorResponse(http.StatusUnauthorized, httpError.ErrPasswordMismatch, nil))
		return
	}

	auth.signInComplete(w, r, &user.User)
}

func (auth *AuthHandler) signInComplete(w http.ResponseWriter, r *http.Request, user *model.User) {
	tokens, err := auth.signInGenerateTokens(user)
	if err != nil {
		render.Render(w, r, httpError.NewErrorResponse(http.StatusInternalServerError, httpError.ErrInternalServer, err))
		return
	}

	response := &SignInResponse{
		AccessToken:           tokens.AccessToken,
		AccessTokenExpiresIn:  tokens.AccessTokenExpiresIn,
		RefreshToken:          tokens.RefreshToken,
		RefreshTokenExpiresIn: tokens.RefreshTokenExpiresIn,
		User: SignInUser{
			ID:    user.ID,
			Email: user.Email,
		},
	}

	render.Render(w, r, response)
}

func (auth *AuthHandler) signInGenerateTokens(user *model.User) (*SignInTokens, error) {
	accessToken, accessTokenExpiresIn, err := auth.Deps.AuthService.CreateAccessToken(user.ID)
	if err != nil {
		return nil, err
	}
	refreshTokenId := uuid.NewString()
	sessionId, err := auth.Deps.DB.Session().Create(user.ID, refreshTokenId)
	if err != nil {
		return nil, err
	}

	refreshToken, refreshTokenExpiresIn, err := auth.Deps.AuthService.CreateRefreshToken(sessionId, refreshTokenId)
	if err != nil {
		return nil, err
	}

	return &SignInTokens{
		AccessToken:           accessToken,
		AccessTokenExpiresIn:  accessTokenExpiresIn,
		RefreshToken:          refreshToken,
		RefreshTokenExpiresIn: refreshTokenExpiresIn,
	}, nil
}
