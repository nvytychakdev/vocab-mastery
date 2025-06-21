package middleware

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	httpError "github.com/nvytychakdev/vocab-mastery/internal/app/http-error"
)

func (mw *Middleware) WordContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		wordId := chi.URLParam(r, "wordId")
		word, err := mw.Deps.DB.GetWordByID(wordId)
		if err != nil {
			render.Render(w, r, httpError.NewErrorResponse(http.StatusNotFound, httpError.ErrNotFound, err))
			return
		}

		ctx := context.WithValue(r.Context(), WORD_KEY, word)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
