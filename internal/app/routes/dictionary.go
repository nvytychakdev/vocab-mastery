package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/nvytychakdev/vocab-mastery/internal/app/handler/dictionary"
	"github.com/nvytychakdev/vocab-mastery/internal/app/handler/word"
	"github.com/nvytychakdev/vocab-mastery/internal/app/middleware"
)

func DictionaryRouter(dh *dictionary.DictionaryHandler, wh *word.WordHandler, mw *middleware.Middleware) *chi.Mux {
	router := chi.NewRouter()
	router.Use(mw.Authorized)
	router.Post("/", dh.DictionaryCreate)
	router.With(mw.IncludeContext).With(mw.QueryOptionsContext).Get("/", dh.DictionaryGetList)
	router.Route("/{dictionaryId}", func(r chi.Router) {
		r.Use(mw.DictionaryContext)
		r.With(mw.IncludeContext).Get("/", dh.DictionaryGetByID)
		r.Delete("/", dh.DictionaryDeleteByID)
		WordDictionaryRouter(r, wh, mw)
	})

	return router
}
