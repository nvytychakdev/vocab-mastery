package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/nvytychakdev/vocab-mastery/internal/app/handler/dictionary"
	"github.com/nvytychakdev/vocab-mastery/internal/app/middleware"
)

func DictionaryRouter(dic *dictionary.DictionaryHandler, mw *middleware.Middleware) *chi.Mux {
	router := chi.NewRouter()
	router.Use(mw.Authorized)
	router.Post("/", dic.DictionaryCreate)
	router.Get("/", dic.DictionaryGetList)
	router.Route("/{dictionaryId}", func(r chi.Router) {
		r.Use(mw.DictionaryContext)
		r.Get("/", dic.DictionaryGetByID)
		r.Delete("/", dic.DictionaryDeleteByID)
	})

	return router
}
