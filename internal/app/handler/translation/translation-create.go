package translation

import (
	"errors"
	"net/http"

	"github.com/go-chi/render"
	httpError "github.com/nvytychakdev/vocab-mastery/internal/app/http-error"
	"github.com/nvytychakdev/vocab-mastery/internal/app/middleware"
)

type TranslationCreateRequest struct {
	Word     string `json:"word"`
	Language string `json:"language"`
}

func (s *TranslationCreateRequest) Bind(r *http.Request) error {
	if s.Word == "" {
		return errors.New("name field is required")
	}
	if s.Language == "" {
		return errors.New("description is required")
	}
	return nil
}

// Response
type TranslationCreateResponse struct {
	ID string `json:"id"`
}

func (*TranslationCreateResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (wh *TranslationHandler) TranslationCreate(w http.ResponseWriter, r *http.Request) {
	word := middleware.GetWordContext(r)

	var data = &TranslationCreateRequest{}

	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, httpError.NewErrorResponse(http.StatusBadRequest, httpError.ErrInvalidPayload, err))
		return
	}

	translationId, err := wh.Deps.DB.CreateTranslation(word.ID, data.Word, data.Language)
	if err != nil {
		render.Render(w, r, httpError.NewErrorResponse(http.StatusInternalServerError, httpError.ErrInternalServer, err))
		return
	}

	response := &TranslationCreateResponse{
		ID: translationId,
	}

	render.Render(w, r, response)
}
