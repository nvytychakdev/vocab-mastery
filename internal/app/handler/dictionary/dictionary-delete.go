package dictionary

import (
	"net/http"

	"github.com/go-chi/render"
	httpError "github.com/nvytychakdev/vocab-mastery/internal/app/http-error"
	"github.com/nvytychakdev/vocab-mastery/internal/app/middleware"
)

type DictionaryDeleteResponse struct {
	Success bool `json:"ok"`
}

func (d DictionaryDeleteResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (auth *DictionaryHandler) DictionaryDeleteByID(w http.ResponseWriter, r *http.Request) {
	dictionary := middleware.GetDictionaryContext(r)

	err := auth.Deps.DB.RemoveDictionaryByID(dictionary.ID)
	if err != nil {
		render.Render(w, r, httpError.NewErrorResponse(http.StatusInternalServerError, httpError.ErrInternalServer, err))
		return
	}

	response := &DictionaryDeleteResponse{
		Success: true,
	}
	render.Render(w, r, response)
}
