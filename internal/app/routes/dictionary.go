package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/nvytychakdev/vocab-mastery/internal/app/handler/dictionary"
	"github.com/nvytychakdev/vocab-mastery/internal/app/handler/word"
	"github.com/nvytychakdev/vocab-mastery/internal/app/middleware"
)

func DictionaryRouter(dh *dictionary.DictionaryHandler, mw *middleware.Middleware) *chi.Mux {
	router := chi.NewRouter()
	router.Use(mw.Authorized)
	router.Post("/", dh.DictionaryCreate)
	router.Get("/", dh.DictionaryGetList)
	router.Route("/{dictionaryId}", func(r chi.Router) {
		r.Use(mw.DictionaryContext)
		r.Get("/", dh.DictionaryGetByID)
		r.Delete("/", dh.DictionaryDeleteByID)

		wh := &word.WordHandler{Deps: dh.Deps}
		WordRouter(r, wh, mw)
	})

	return router
}
