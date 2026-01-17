package auth

import (
	"net/http"
	"time"

	"github.com/go-chi/render"
	"github.com/google/uuid"
	httpError "github.com/nvytychakdev/vocab-mastery/internal/app/http-error"
	"github.com/nvytychakdev/vocab-mastery/internal/app/middleware"
)

type ProfileResponse struct {
	ID         uuid.UUID `json:"id"`
	CreatedAt  time.Time `json:"createdAt"`
	Email      string    `json:"email"`
	Name       string    `json:"name"`
	PictureUrl *string   `json:"pictureUrl"`
}

func (u *ProfileResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (auth *AuthHandler) Profile(w http.ResponseWriter, r *http.Request) {
	userId := middleware.GetAuthorizedUserId(r)

	user, err := auth.Deps.DB.User().GetByID(userId)
	if err != nil {
		render.Render(w, r, httpError.NewErrorResponse(http.StatusUnauthorized, httpError.ErrInternalServer, err))
		return
	}

	response := &ProfileResponse{
		ID:         user.ID,
		CreatedAt:  user.CreatedAt,
		Email:      user.UserData.Email,
		Name:       user.UserData.Name,
		PictureUrl: user.UserData.PictureUrl,
	}
	render.Render(w, r, response)
}
