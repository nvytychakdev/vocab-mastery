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
	migrationsData "github.com/nvytychakdev/vocab-mastery/internal/app/db/migrations-data"
	"github.com/nvytychakdev/vocab-mastery/internal/app/handler/auth"
	"github.com/nvytychakdev/vocab-mastery/internal/app/handler/dictionary"
	"github.com/nvytychakdev/vocab-mastery/internal/app/handler/flashcard"
	"github.com/nvytychakdev/vocab-mastery/internal/app/handler/language"
	"github.com/nvytychakdev/vocab-mastery/internal/app/handler/word"
	vmMiddleware "github.com/nvytychakdev/vocab-mastery/internal/app/middleware"
	"github.com/nvytychakdev/vocab-mastery/internal/app/routes"
	"github.com/nvytychakdev/vocab-mastery/internal/app/services"
)

func StartServer() {
	db := db.Connect()
	deps := &services.Deps{
		DB:                      db,
		AuthService:             services.NewAuthService(),
		PasswordService:         services.NewPasswordService(),
		FlashcardSessionService: services.NewFlashcardSessionService(db),
		FlashcardCardService:    services.NewFlashcardCardService(db),
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
	langHandler := &language.LanguageHandler{Deps: deps}
	wordHandler := &word.WordHandler{Deps: deps}
	flashcardHandler := &flashcard.FlashcardHandler{Deps: deps}
	mw := vmMiddleware.NewMiddleware(deps)

	migrationRepo := deps.DB.Migration()
	migrationsData.RunLatest(migrationRepo)

	router.Mount("/api/v1/auth", routes.AuthRouter(authHandler, mw))
	router.Mount("/api/v1/dictionaries", routes.DictionaryRouter(dictionaryHandler, wordHandler, mw))
	router.Mount("/api/v1/words", routes.WordRouter(wordHandler, mw))
	router.Mount("/api/v1/language", routes.LanguageRoutes(langHandler, mw))
	router.Mount("/api/v1/flashcards", routes.FlashcardRouter(flashcardHandler, mw))

	docgen.PrintRoutes(router)
	slog.Info("Started server...")
	http.ListenAndServe(":8080", router)
}
