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
	pagination := middleware.GetPaginationContext(r)
	userId := middleware.GetAuthorizedUserId(r)

	dictionaries, total, err := auth.Deps.DB.Dictionary().ListByUserId(userId, pagination)
	if err != nil {
		return
	}

	response := &DictionaryGetListResponse{
		Items: dictionaries,
		PaginationResponse: &model.PaginationResponse{
			Total:  total,
			Offset: pagination.Offset,
			Limit:  pagination.Limit,
		},
	}

	render.Render(w, r, response)
}
