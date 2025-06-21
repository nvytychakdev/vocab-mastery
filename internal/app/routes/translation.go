package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/nvytychakdev/vocab-mastery/internal/app/handler/translation"
	"github.com/nvytychakdev/vocab-mastery/internal/app/middleware"
)

func TranslationRouter(th *translation.TranslationHandler, mw *middleware.Middleware) func(r chi.Router) {
	return func(router chi.Router) {
		router.Post("/", th.TranslationCreate)
		router.Get("/", th.TranslationGetList)
		router.Route("/{translationId}", func(r chi.Router) {
			r.Use(mw.TranslationContext)
			r.Get("/", th.TranslationGetByID)
			r.Delete("/", th.TranslationDeleteByID)
		})
	}

}
