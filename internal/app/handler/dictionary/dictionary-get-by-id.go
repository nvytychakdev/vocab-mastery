package dictionary

import (
	"net/http"
	"time"

	"github.com/go-chi/render"
	"github.com/nvytychakdev/vocab-mastery/internal/app/middleware"
	"github.com/nvytychakdev/vocab-mastery/internal/app/model"
)

type DictionaryGetByIdResponse struct {
	ID          string        `json:"id"`
	CreatedAt   time.Time     `json:"craetedAt"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Words       []*model.Word `json:"words,omitempty"`
}

func (u *DictionaryGetByIdResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (dh *DictionaryHandler) DictionaryGetByID(w http.ResponseWriter, r *http.Request) {
	dictionary := middleware.GetDictionaryContext(r)
	include := middleware.GetIncludeContext(r)

	response := &DictionaryGetByIdResponse{
		ID:          dictionary.ID,
		CreatedAt:   dictionary.CreatedAt,
		Name:        dictionary.Name,
		Description: dictionary.Description,
	}

	if include != nil && include["words"] {
		words, _, err := dh.Deps.DB.Word().GetByDictionaryID(dictionary.ID, nil)
		if err == nil {
			response.Words = words
		}
	}

	render.Render(w, r, response)
}
