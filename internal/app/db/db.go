package db

import (
	"log/slog"
	"os"
	"strconv"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx"
	"github.com/nvytychakdev/vocab-mastery/internal/app/model"
)

type DB interface {
	Session() SessionRepo
	User() UserRepo
	UserToken() UserTokenRepo
	Dictionary() DictionaryRepo
	Word() WordRepo
	Translation() TranslationRepo
}

type PostgresDB struct {
	conn *pgx.Conn
	psql squirrel.StatementBuilderType
}

const MAX_RETRIES = 5

func Connect() DB {
	port, err := strconv.ParseUint(os.Getenv("POSTGRES_PORT"), 10, 16)
	if err != nil {
		slog.Debug("Failed to retrieve postgres port")
	}

	psql := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
	for range MAX_RETRIES {
		conn, err := pgx.Connect(pgx.ConnConfig{
			Host:     os.Getenv("POSTGRES_HOST"),
			Port:     uint16(port),
			User:     os.Getenv("POSTGRES_USER"),
			Password: os.Getenv("POSTGRES_PASSWORD"),
			Database: os.Getenv("POSTGRES_DB"),
		})
		if err == nil {
			slog.Info("Database connection established")
			return &PostgresDB{conn, psql}
		}

		slog.Error("Not able to connect to the database, retry again.", "Error", err)
		time.Sleep(5 * time.Second)
	}

	return nil
}

func ApplyQueryOptions(builder squirrel.SelectBuilder, opts *model.QueryOptions) squirrel.SelectBuilder {

	if opts.Pagination != nil {
		builder = builder.Offset(uint64(opts.Pagination.Offset)).Limit(uint64(opts.Pagination.Limit))
	}

	if opts.Sort != nil {
		orderBy := opts.Sort.Field
		if opts.Sort.Direction == "desc" {
			orderBy += " DESC"
		}
		builder = builder.OrderBy(orderBy)
	}

	return builder
}
