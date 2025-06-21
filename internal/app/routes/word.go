package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/nvytychakdev/vocab-mastery/internal/app/handler/translation"
	"github.com/nvytychakdev/vocab-mastery/internal/app/handler/word"
	"github.com/nvytychakdev/vocab-mastery/internal/app/middleware"
)

func WordRouter(r chi.Router, wh *word.WordHandler, mw *middleware.Middleware) {
	r.Route("/words", func(r chi.Router) {
		r.Post("/", wh.WordCreate)
		r.Get("/", wh.WordGetList)
		r.Route("/{wordId}", func(r chi.Router) {
			r.Use(mw.WordContext)
			r.Get("/", wh.WordGetByID)
			r.Delete("/", wh.WordDeleteByID)

			th := &translation.TranslationHandler{Deps: wh.Deps}
			r.Route("/translations", TranslationRouter(th, mw))
		})
	})
}
