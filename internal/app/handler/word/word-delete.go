package word

import (
	"net/http"

	"github.com/go-chi/render"
	httpError "github.com/nvytychakdev/vocab-mastery/internal/app/http-error"
	"github.com/nvytychakdev/vocab-mastery/internal/app/middleware"
)

type WordDeleteResponse struct {
	Success bool `json:"ok"`
}

func (d WordDeleteResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (wh *WordHandler) WordDeleteByID(w http.ResponseWriter, r *http.Request) {
	word := middleware.GetWordContext(r)

	err := wh.Deps.DB.Word().DeleteByID(word.ID)
	if err != nil {
		render.Render(w, r, httpError.NewErrorResponse(http.StatusInternalServerError, httpError.ErrInternalServer, err))
		return
	}

	response := &WordDeleteResponse{
		Success: true,
	}
	render.Render(w, r, response)
}
