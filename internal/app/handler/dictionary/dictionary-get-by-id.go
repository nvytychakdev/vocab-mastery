package dictionary

import (
	"net/http"
	"time"

	"github.com/go-chi/render"
	"github.com/nvytychakdev/vocab-mastery/internal/app/middleware"
	"github.com/nvytychakdev/vocab-mastery/internal/app/model"
)

type DictionaryGetByIdResponse struct {
	ID          string    `json:"id"`
	CreatedAt   time.Time `json:"craetedAt"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
}

func (u *DictionaryGetByIdResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (auth *DictionaryHandler) DictionaryGetByID(w http.ResponseWriter, r *http.Request) {
	dictionary := r.Context().Value(middleware.DICTIONARY_KEY).(*model.Dictionary)

	response := &DictionaryGetByIdResponse{
		ID:          dictionary.ID,
		CreatedAt:   dictionary.CreatedAt,
		Name:        dictionary.Name,
		Description: dictionary.Description,
	}
	render.Render(w, r, response)
}
