package word

import (
	"net/http"
	"time"

	"github.com/go-chi/render"
	"github.com/nvytychakdev/vocab-mastery/internal/app/middleware"
	"github.com/nvytychakdev/vocab-mastery/internal/app/model"
)

type WordGetByIdResponse struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"craetedAt"`
	Word      string    `json:"name"`
	Language  string    `json:"description"`
}

func (u *WordGetByIdResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (wh *WordHandler) WordGetByID(w http.ResponseWriter, r *http.Request) {
	word := r.Context().Value(middleware.WORD_KEY).(*model.Word)

	response := &WordGetByIdResponse{
		ID:        word.ID,
		CreatedAt: word.CreatedAt,
		Word:      word.Word,
		Language:  word.Language,
	}
	render.Render(w, r, response)
}
