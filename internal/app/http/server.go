package http

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/docgen"
	"github.com/go-chi/render"
	"github.com/nvytychakdev/vocab-mastery/internal/app/db"
	"github.com/nvytychakdev/vocab-mastery/internal/app/handler/auth"
	"github.com/nvytychakdev/vocab-mastery/internal/app/handler/dictionary"
	vmMiddleware "github.com/nvytychakdev/vocab-mastery/internal/app/middleware"
	"github.com/nvytychakdev/vocab-mastery/internal/app/routes"
	"github.com/nvytychakdev/vocab-mastery/internal/app/services"
)

func StartServer() {
	deps := &services.Deps{
		DB:              db.Connect(),
		AuthService:     services.NewAuthService(),
		PasswordService: services.NewPasswordService(),
	}

	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://*", "https://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)
	router.Use(render.SetContentType(render.ContentTypeJSON))

	authHandler := &auth.AuthHandler{Deps: deps}
	dictionaryHandler := &dictionary.DictionaryHandler{Deps: deps}
	mw := vmMiddleware.NewMiddleware(deps)

	router.Mount("/api/v1/auth", routes.AuthRouter(authHandler, mw))
	router.Mount("/api/v1/dictionaries", routes.DictionaryRouter(dictionaryHandler, mw))

	docgen.PrintRoutes(router)

	http.ListenAndServe(":8080", router)
	slog.Info("Started server...")
}
