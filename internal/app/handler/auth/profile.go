package auth

import (
	"net/http"
	"time"

	"github.com/go-chi/render"
	"github.com/nvytychakdev/vocab-mastery/internal/app/db"
	httpError "github.com/nvytychakdev/vocab-mastery/internal/app/http-error"
	"github.com/nvytychakdev/vocab-mastery/internal/app/middleware"
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

func Profile(w http.ResponseWriter, r *http.Request) {
	userId, ok := r.Context().Value(middleware.USER_ID_KEY).(string)

	if !ok {
		render.Render(w, r, httpError.NewErrorResponse(http.StatusForbidden, httpError.ErrUnauthorized, nil))
		return
	}

	user, err := db.GetUserByID(userId)
	if err != nil {
		render.Render(w, r, httpError.NewErrorResponse(http.StatusUnauthorized, httpError.ErrInternalServer, err))
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
