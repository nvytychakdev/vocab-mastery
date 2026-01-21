package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/nvytychakdev/vocab-mastery/internal/app/handler/flashcard"
	"github.com/nvytychakdev/vocab-mastery/internal/app/middleware"
)

func FlashcardRouter(dh *flashcard.FlashcardHandler, mw *middleware.Middleware) *chi.Mux {
	router := chi.NewRouter()
	router.Use(mw.Authorized)
	router.Post("/sessions/start", dh.FlashcardSessionStart)
	router.Route("/sessions/{sessionId}", func(r chi.Router) {
		r.Use(mw.FlashcardSessionContext)
		r.With(mw.IncludeContext).Post("/answer", dh.FlashcardSessionAnswer)
	})
	return router
}
