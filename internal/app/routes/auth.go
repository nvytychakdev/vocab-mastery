package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/nvytychakdev/vocab-mastery/internal/app/handler/auth"
	"github.com/nvytychakdev/vocab-mastery/internal/app/middleware"
)

func AuthRouter() *chi.Mux {
	router := chi.NewRouter()
	router.Post("/sign-in", auth.SignIn)
	router.Post("/sign-up", auth.SignUp)
	router.Post("/refresh-token", auth.RefreshToken)
	router.Post("/confirm-email", auth.ConfirmEmail)
	router.Post("/resend-confirm-email", auth.ResendEmailConfirm)

	authorizedRouter := chi.NewRouter()
	authorizedRouter.Use(middleware.Authorized)
	authorizedRouter.Get("/profile", auth.Profile)
	authorizedRouter.Post("/sign-out", auth.SignOut)

	router.Mount("/", authorizedRouter)

	return router
}
