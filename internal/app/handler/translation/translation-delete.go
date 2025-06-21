package translation

import (
	"net/http"

	"github.com/go-chi/render"
	httpError "github.com/nvytychakdev/vocab-mastery/internal/app/http-error"
	"github.com/nvytychakdev/vocab-mastery/internal/app/middleware"
)

type TranslationDeleteResponse struct {
	Success bool `json:"ok"`
}

func (d TranslationDeleteResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (th *TranslationHandler) TranslationDeleteByID(w http.ResponseWriter, r *http.Request) {
	translation := middleware.GetTranslationContext(r)

	err := th.Deps.DB.RemoveTranslationByID(translation.ID)
	if err != nil {
		render.Render(w, r, httpError.NewErrorResponse(http.StatusInternalServerError, httpError.ErrInternalServer, err))
		return
	}

	response := &TranslationDeleteResponse{
		Success: true,
	}
	render.Render(w, r, response)
}
