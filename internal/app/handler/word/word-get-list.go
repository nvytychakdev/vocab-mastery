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
	*model.PaginationResponse
}

func (u *WordGetListResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (wh *WordHandler) WordGetList(w http.ResponseWriter, r *http.Request) {
	pagination := middleware.GetPaginationContext(r)
	dictionary := middleware.GetDictionaryContext(r)

	words, totalWords, err := wh.Deps.DB.GetAllWordsByDictionaryID(dictionary.ID, pagination)
	if err != nil {
		slog.Error("Not able to get words by dictionary id", "error", err)
		return
	}

	response := &WordGetListResponse{
		Items: words,
		PaginationResponse: &model.PaginationResponse{
			Total:  totalWords,
			Offset: pagination.Offset,
			Limit:  pagination.Limit,
		},
	}

	render.Render(w, r, response)
}
