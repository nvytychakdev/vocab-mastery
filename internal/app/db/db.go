package db

import (
	"log/slog"
	"os"
	"strconv"

	"github.com/jackc/pgx"
)

type DB interface {
	SessionRepository
	UserRepository
	UserTokenRepository
}

type PostgresDB struct {
	conn *pgx.Conn
}

var (
	Instance DB
)

func Connect() {
	port, err := strconv.ParseUint(os.Getenv("POSTGRES_PORT"), 10, 16)
	if err != nil {
		slog.Debug("Failed to retrieve postgres port")
	}

	conn, err := pgx.Connect(pgx.ConnConfig{
		Host:     os.Getenv("POSTGRES_HOST"),
		Port:     uint16(port),
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		Database: os.Getenv("POSTGRES_DB"),
	})
	if err != nil {
		slog.Error("Not able to connect to the database", "Error", err)
		return
	}
	Instance = &PostgresDB{conn}
}
