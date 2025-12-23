package middleware

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	httpError "github.com/nvytychakdev/vocab-mastery/internal/app/http-error"
	"github.com/nvytychakdev/vocab-mastery/internal/app/model"
)

func (mw *Middleware) TranslationContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		translationIdParam := chi.URLParam(r, "translationId")

		translationId, err := uuid.Parse(translationIdParam)
		if err != nil {
			render.Render(w, r, httpError.NewErrorResponse(http.StatusNotFound, httpError.ErrNotFound, err))
			return
		}

		translation, err := mw.Deps.DB.Translation().GetByID(translationId)
		if err != nil {
			render.Render(w, r, httpError.NewErrorResponse(http.StatusNotFound, httpError.ErrNotFound, err))
			return
		}

		ctx := context.WithValue(r.Context(), TRANSLATION_KEY, translation)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetTranslationContext(r *http.Request) *model.Translation {
	translation, ok := r.Context().Value(TRANSLATION_KEY).(*model.Translation)
	if !ok {
		slog.Error("Not able to retrieve translation context")
		return nil
	}
	return translation
}
