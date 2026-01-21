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

func (mw *Middleware) FlashcardSessionContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sessionIDParam := chi.URLParam(r, "sessionId")

		sessionID, err := uuid.Parse(sessionIDParam)
		if err != nil {
			render.Render(w, r, httpError.NewErrorResponse(http.StatusNotFound, httpError.ErrNotFound, err))
			return
		}

		session, err := mw.Deps.DB.FlashcardSession().GetByID(sessionID)
		if err != nil {
			render.Render(w, r, httpError.NewErrorResponse(http.StatusNotFound, httpError.ErrNotFound, err))
			return
		}

		ctx := context.WithValue(r.Context(), FLASHCARD_SESSION_KEY, session)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetFlashcardSessionContext(r *http.Request) *model.FlashcardSession {
	session, ok := r.Context().Value(FLASHCARD_SESSION_KEY).(*model.FlashcardSession)
	if !ok {
		slog.Error("Not able to parse flashcard session by id")
		return nil
	}

	return session
}
