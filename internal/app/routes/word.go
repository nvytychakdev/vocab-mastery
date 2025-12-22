package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/nvytychakdev/vocab-mastery/internal/app/handler/translation"
	"github.com/nvytychakdev/vocab-mastery/internal/app/handler/word"
	"github.com/nvytychakdev/vocab-mastery/internal/app/middleware"
)

func WordDictionaryRouter(r chi.Router, wh *word.WordHandler, mw *middleware.Middleware) {
	r.Route("/words", func(r chi.Router) {
		r.Post("/", wh.WordCreate)
		r.With(mw.IncludeContext).With(mw.QueryOptionsContext).Get("/", wh.WordGetListByDictionary)
		r.Route("/{wordId}", func(r chi.Router) {
			r.Use(mw.WordContext)
			r.With(mw.IncludeContext).Get("/", wh.WordGetByID)
			r.Delete("/", wh.WordDeleteByID)

			th := &translation.TranslationHandler{Deps: wh.Deps}
			r.Route("/translations", TranslationRouter(th, mw))
		})
	})
}

func WordRouter(wh *word.WordHandler, mw *middleware.Middleware) *chi.Mux {
	router := chi.NewRouter()
	router.Use(mw.Authorized)
	router.With(mw.IncludeContext).With(mw.QueryOptionsContext).Get("/", wh.WordGetList)
	router.Route("/{wordId}", func(r chi.Router) {
		r.Use(mw.WordContext)
		r.With(mw.IncludeContext).Get("/", wh.WordGetByID)
		r.Delete("/", wh.WordDeleteByID)
	})

	return router
}
