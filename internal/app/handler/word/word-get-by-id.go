package word

import (
	"net/http"
	"time"

	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/nvytychakdev/vocab-mastery/internal/app/middleware"
	"github.com/nvytychakdev/vocab-mastery/internal/app/model"
)

type WordGetByIdResponse struct {
	ID           uuid.UUID            `json:"id"`
	CreatedAt    time.Time            `json:"craetedAt"`
	Word         string               `json:"name"`
	Language     string               `json:"description"`
	Translations []*model.Translation `json:"translations,omitempty"`
}

func (u *WordGetByIdResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (wh *WordHandler) WordGetByID(w http.ResponseWriter, r *http.Request) {
	word := middleware.GetWordContext(r)
	include := middleware.GetIncludeContext(r)

	response := &WordGetByIdResponse{
		ID:        word.ID,
		Word:      word.Word,
		CreatedAt: word.CreatedAt,
	}

	if include != nil && include["translations"] {
		translations, _, err := wh.Deps.DB.Translation().ListByWordID(word.ID, nil)
		if err == nil {
			response.Translations = translations
		}
	}

	render.Render(w, r, response)
}
