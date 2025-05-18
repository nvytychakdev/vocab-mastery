package auth

import (
	"errors"
	"net/http"
	"time"

	"github.com/go-chi/render"
	httpError "github.com/nvytychakdev/vocab-mastery/internal/app/http-error"
	"github.com/nvytychakdev/vocab-mastery/internal/app/middleware"
	"github.com/nvytychakdev/vocab-mastery/internal/app/model/user"
)

type ProfileResponse struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"craetedAt"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
}

func (u *ProfileResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func profile(w http.ResponseWriter, r *http.Request) {
	userId, ok := r.Context().Value(middleware.USER_ID_KEY).(string)

	if !ok {
		render.Render(w, r, httpError.NewErrorResponse(errors.New("user is not authorized"), http.StatusUnauthorized))
		return
	}

	user, err := user.GetUserByID(userId)
	if err != nil {
		render.Render(w, r, httpError.NewErrorResponse(err, http.StatusUnauthorized))
		return
	}

	response := &ProfileResponse{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		Email:     user.UserData.Email,
		Name:      user.UserData.Name,
	}
	render.Render(w, r, response)
}
