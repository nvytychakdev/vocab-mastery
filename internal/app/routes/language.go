package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/nvytychakdev/vocab-mastery/internal/app/handler/language"
	"github.com/nvytychakdev/vocab-mastery/internal/app/middleware"
)

func LanguageRoutes(dh *language.LanguageHandler, mw *middleware.Middleware) *chi.Mux {
	router := chi.NewRouter()
	router.Use(mw.Authorized)
	router.Post("/translate", dh.LanguageTranslate)
	return router
}
