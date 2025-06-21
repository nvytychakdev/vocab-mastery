package translation

import (
	"net/http"
	"time"

	"github.com/go-chi/render"
	"github.com/nvytychakdev/vocab-mastery/internal/app/middleware"
	"github.com/nvytychakdev/vocab-mastery/internal/app/model"
)

type TranslationGetByIdResponse struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"craetedAt"`
	Word      string    `json:"name"`
	Language  string    `json:"description"`
}

func (u *TranslationGetByIdResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (th *TranslationHandler) TranslationGetByID(w http.ResponseWriter, r *http.Request) {
	translation := r.Context().Value(middleware.TRANSLATION_KEY).(*model.Translation)

	response := &TranslationGetByIdResponse{
		ID:        translation.ID,
		CreatedAt: translation.CreatedAt,
		Word:      translation.Word,
		Language:  translation.Language,
	}
	render.Render(w, r, response)
}
