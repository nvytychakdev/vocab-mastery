package dictionary

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/nvytychakdev/vocab-mastery/internal/app/middleware"
	"github.com/nvytychakdev/vocab-mastery/internal/app/model"
)

type DictionaryGetListResponse struct {
	Items []*model.Dictionary `json:"items"`
	Total int                 `json:"total"`
}

func (u *DictionaryGetListResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (auth *DictionaryHandler) DictionaryGetList(w http.ResponseWriter, r *http.Request) {
	userId, ok := r.Context().Value(middleware.USER_ID_KEY).(string)
	if !ok {
		return
	}

	dictionaries, err := auth.Deps.DB.GetAllDictionariesByUsedID(userId)
	if err != nil {
		return
	}

	response := &DictionaryGetListResponse{
		Items: dictionaries,
		Total: len(dictionaries),
	}

	render.Render(w, r, response)
}
