package dictionary

import (
	"net/http"
	"time"

	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/nvytychakdev/vocab-mastery/internal/app/middleware"
	"github.com/nvytychakdev/vocab-mastery/internal/app/model"
)

type DictionaryGetByIdResponse struct {
	ID        uuid.UUID     `json:"id"`
	CreatedAt time.Time     `json:"craetedAt"`
	Title     string        `json:"name"`
	Words     []*model.Word `json:"words,omitempty"`
}

func (u *DictionaryGetByIdResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (dh *DictionaryHandler) DictionaryGetByID(w http.ResponseWriter, r *http.Request) {
	dictionary := middleware.GetDictionaryContext(r)
	include := middleware.GetIncludeContext(r)

	response := &DictionaryGetByIdResponse{
		ID:        dictionary.ID,
		CreatedAt: dictionary.CreatedAt,
		Title:     dictionary.Title,
	}

	if include != nil && include["words"] {
		words, _, err := dh.Deps.DB.Word().ListByDictionaryID(dictionary.ID, nil)
		if err == nil {
			response.Words = words
		}
	}

	render.Render(w, r, response)
}
