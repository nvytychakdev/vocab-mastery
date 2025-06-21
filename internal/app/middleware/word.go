package middleware

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	httpError "github.com/nvytychakdev/vocab-mastery/internal/app/http-error"
	"github.com/nvytychakdev/vocab-mastery/internal/app/model"
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

func GetWordContext(r *http.Request) *model.Word {
	word, ok := r.Context().Value(WORD_KEY).(*model.Word)
	if !ok {
		slog.Error("Not able to retrieve word context")
		return nil
	}
	return word
}
