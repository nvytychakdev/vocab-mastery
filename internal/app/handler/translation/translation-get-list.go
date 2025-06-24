package translation

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/nvytychakdev/vocab-mastery/internal/app/middleware"
	"github.com/nvytychakdev/vocab-mastery/internal/app/model"
)

type TranslationGetListResponse struct {
	Items []*model.Translation `json:"items"`
	*model.PaginationResponse
}

func (u *TranslationGetListResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (th *TranslationHandler) TranslationGetList(w http.ResponseWriter, r *http.Request) {
	opts := middleware.GetQueryOptionsContext(r)
	word := middleware.GetWordContext(r)

	translations, total, err := th.Deps.DB.Translation().ListByWordID(word.ID, opts)
	if err != nil {
		return
	}

	response := &TranslationGetListResponse{
		Items: translations,
		PaginationResponse: &model.PaginationResponse{
			Total:  total,
			Offset: opts.Pagination.Offset,
			Limit:  opts.Pagination.Limit,
		},
	}

	render.Render(w, r, response)
}
