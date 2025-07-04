package auth

import (
	"errors"
	"net/http"

	"github.com/go-chi/render"
	httpError "github.com/nvytychakdev/vocab-mastery/internal/app/http-error"
)

type ResendTokenConfirm struct {
	Email string `json:"email"`
}

func (rt *ResendTokenConfirm) Bind(r *http.Request) error {
	if rt.Email == "" {
		return errors.New("email is required")
	}
	return nil
}

func (auth *AuthHandler) ResendEmailConfirm(w http.ResponseWriter, r *http.Request) {
	data := &ResendTokenConfirm{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, httpError.NewErrorResponse(http.StatusBadRequest, httpError.ErrInvalidPayload, err))
		return
	}

	user, err := auth.Deps.DB.User().GetByEmail(data.Email)
	if err != nil {
		render.Render(w, r, httpError.NewErrorResponse(http.StatusInternalServerError, httpError.ErrInternalServer, err))
		return
	}

	auth.sendEmailConfirm(w, r, user.ID, user.Email)
}
