package auth

import (
	"github.com/go-chi/chi/v5"
	"github.com/nvytychakdev/vocab-mastery/internal/app/middleware"
)

func AuthRouter() *chi.Mux {
	router := chi.NewRouter()
	router.Post("/sign-in", signIn)
	router.Post("/sign-up", signUp)
	router.Post("/refresh-token", refreshToken)

	authorizedRouter := chi.NewRouter()
	authorizedRouter.Use(middleware.Authorized)
	authorizedRouter.Get("/profile", profile)
	authorizedRouter.Post("/sign-out", signOut)

	router.Mount("/", authorizedRouter)

	return router
}
