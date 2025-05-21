package main

import (
	"log/slog"

	"github.com/joho/godotenv"
	"github.com/nvytychakdev/vocab-mastery/internal/app/http"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		slog.Debug("Failed to load environment file, looking for global env...")
	}

	http.StartServer()
}
