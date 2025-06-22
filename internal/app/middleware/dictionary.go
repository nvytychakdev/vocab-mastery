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

func (mw *Middleware) DictionaryContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		dictionaryId := chi.URLParam(r, "dictionaryId")
		dictionary, err := mw.Deps.DB.Dictionary().GetByID(dictionaryId)
		if err != nil {
			render.Render(w, r, httpError.NewErrorResponse(http.StatusNotFound, httpError.ErrNotFound, err))
			return
		}

		ctx := context.WithValue(r.Context(), DICTIONARY_KEY, dictionary)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetDictionaryContext(r *http.Request) *model.Dictionary {
	dictionary, ok := r.Context().Value(DICTIONARY_KEY).(*model.Dictionary)
	if !ok {
		slog.Error("Not ableto parse dictionary by id")
		return nil
	}

	return dictionary
}
