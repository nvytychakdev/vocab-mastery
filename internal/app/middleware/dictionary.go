package middleware

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	httpError "github.com/nvytychakdev/vocab-mastery/internal/app/http-error"
)

func (mw *Middleware) DictionaryContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		dictionaryId := chi.URLParam(r, "dictionaryId")
		dictionary, err := mw.Deps.DB.GetDictionaryByID(dictionaryId)
		if err != nil {

			render.Render(w, r, httpError.NewErrorResponse(http.StatusNotFound, httpError.ErrNotFound, err))
			return
		}

		ctx := context.WithValue(r.Context(), DICTIONARY_KEY, dictionary)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
