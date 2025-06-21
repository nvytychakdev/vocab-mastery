package word

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/render"
	"github.com/nvytychakdev/vocab-mastery/internal/app/middleware"
	"github.com/nvytychakdev/vocab-mastery/internal/app/model"
)

type WordGetListResponse struct {
	Items []*model.Word `json:"items"`
	Total int           `json:"total"`
}

func (u *WordGetListResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (wh *WordHandler) WordGetList(w http.ResponseWriter, r *http.Request) {
	dictionary, ok := r.Context().Value(middleware.DICTIONARY_KEY).(*model.Dictionary)
	if !ok {
		slog.Error("Not ableto parse dictionary by id")
		return
	}

	words, err := wh.Deps.DB.GetAllWordsByDictionaryID(dictionary.ID)
	if err != nil {
		slog.Error("Not able to get words by dictionary id", "error", err)
		return
	}

	response := &WordGetListResponse{
		Items: words,
		Total: len(words),
	}

	render.Render(w, r, response)
}
