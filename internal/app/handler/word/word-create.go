package word

import (
	"errors"
	"net/http"

	"github.com/go-chi/render"
	httpError "github.com/nvytychakdev/vocab-mastery/internal/app/http-error"
	"github.com/nvytychakdev/vocab-mastery/internal/app/middleware"
	"github.com/nvytychakdev/vocab-mastery/internal/app/model"
)

type WordCreateRequest struct {
	Word     string `json:"word"`
	Language string `json:"language"`
}

func (s *WordCreateRequest) Bind(r *http.Request) error {
	if s.Word == "" {
		return errors.New("name field is required")
	}
	if s.Language == "" {
		return errors.New("description is required")
	}
	return nil
}

// Response
type WordCreateResponse struct {
	ID string `json:"id"`
}

func (*WordCreateResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (wh *WordHandler) WordCreate(w http.ResponseWriter, r *http.Request) {
	dictionary, ok := r.Context().Value(middleware.DICTIONARY_KEY).(*model.Dictionary)
	if !ok {
		render.Render(w, r, httpError.NewErrorResponse(http.StatusInternalServerError, httpError.ErrInternalServer, nil))
		return
	}

	var data = &WordCreateRequest{}

	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, httpError.NewErrorResponse(http.StatusBadRequest, httpError.ErrInvalidPayload, err))
		return
	}

	wordId, err := wh.Deps.DB.CreateWord(dictionary.ID, data.Word, data.Language)
	if err != nil {
		render.Render(w, r, httpError.NewErrorResponse(http.StatusInternalServerError, httpError.ErrInternalServer, err))
		return
	}

	response := &WordCreateResponse{
		ID: wordId,
	}

	render.Render(w, r, response)
}
