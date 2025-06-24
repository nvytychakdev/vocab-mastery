package dictionary

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/nvytychakdev/vocab-mastery/internal/app/middleware"
	"github.com/nvytychakdev/vocab-mastery/internal/app/model"
)

type DictionaryGetListResponse struct {
	Items []*model.Dictionary `json:"items"`
	*model.PaginationResponse
}

func (u *DictionaryGetListResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (auth *DictionaryHandler) DictionaryGetList(w http.ResponseWriter, r *http.Request) {
	opts := middleware.GetQueryOptionsContext(r)
	userId := middleware.GetAuthorizedUserId(r)

	dictionaries, total, err := auth.Deps.DB.Dictionary().ListByUserId(userId, opts)
	if err != nil {
		return
	}

	response := &DictionaryGetListResponse{
		Items: dictionaries,
		PaginationResponse: &model.PaginationResponse{
			Total:  total,
			Offset: opts.Pagination.Offset,
			Limit:  opts.Pagination.Limit,
		},
	}

	render.Render(w, r, response)
}
