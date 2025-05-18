package http

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/nvytychakdev/vocab-mastery/internal/app/handler/auth"
)

func StartServer() {
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)
	router.Use(render.SetContentType(render.ContentTypeJSON))

	router.Mount("/api/v1/auth", auth.AuthRouter())
	http.ListenAndServe(":8080", router)
	slog.Info("Started server...")
}
