package translation

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/nvytychakdev/vocab-mastery/internal/app/middleware"
	"github.com/nvytychakdev/vocab-mastery/internal/app/model"
)

type TranslationGetListResponse struct {
	Items []*model.Translation `json:"items"`
	Total int                  `json:"total"`
}

func (u *TranslationGetListResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (th *TranslationHandler) TranslationGetList(w http.ResponseWriter, r *http.Request) {
	word, ok := r.Context().Value(middleware.WORD_KEY).(*model.Word)
	if !ok {
		return
	}

	translations, err := th.Deps.DB.GetAllTranslationsByWordID(word.ID)
	if err != nil {
		return
	}

	response := &TranslationGetListResponse{
		Items: translations,
		Total: len(translations),
	}

	render.Render(w, r, response)
}
