package translation

import (
	"net/http"
	"time"

	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/nvytychakdev/vocab-mastery/internal/app/middleware"
)

type TranslationGetByIdResponse struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	Word      string    `json:"name"`
	Language  string    `json:"description"`
}

func (u *TranslationGetByIdResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (th *TranslationHandler) TranslationGetByID(w http.ResponseWriter, r *http.Request) {
	translation := middleware.GetTranslationContext(r)

	response := &TranslationGetByIdResponse{
		ID:        translation.ID,
		CreatedAt: translation.CreatedAt,
		Word:      translation.Word,
		Language:  translation.Language,
	}
	render.Render(w, r, response)
}
