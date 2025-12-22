package dictionary

import (
	"errors"
	"net/http"

	"github.com/go-chi/render"
	httpError "github.com/nvytychakdev/vocab-mastery/internal/app/http-error"
	"github.com/nvytychakdev/vocab-mastery/internal/app/middleware"
)

type DictionaryCreateRequest struct {
	Title string `json:"name"`
}

func (s *DictionaryCreateRequest) Bind(r *http.Request) error {
	if s.Title == "" {
		return errors.New("name field is required")
	}
	return nil
}

// Response
type DictionaryCreateResponse struct {
	ID string `json:"id"`
}

func (*DictionaryCreateResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (dh *DictionaryHandler) DictionaryCreate(w http.ResponseWriter, r *http.Request) {
	userId, ok := r.Context().Value(middleware.USER_ID_KEY).(string)
	if !ok {
		render.Render(w, r, httpError.NewErrorResponse(http.StatusInternalServerError, httpError.ErrInternalServer, nil))
		return
	}

	var data = &DictionaryCreateRequest{}

	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, httpError.NewErrorResponse(http.StatusBadRequest, httpError.ErrInvalidPayload, err))
		return
	}

	dictionaryId, err := dh.Deps.DB.Dictionary().Create(userId, data.Title)
	if err != nil {
		render.Render(w, r, httpError.NewErrorResponse(http.StatusInternalServerError, httpError.ErrInternalServer, err))
		return
	}

	response := &DictionaryCreateResponse{
		ID: dictionaryId,
	}

	render.Render(w, r, response)
}
