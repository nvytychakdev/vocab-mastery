package db

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nvytychakdev/vocab-mastery/internal/app/model"
)

type DB interface {
	Session() SessionRepo
	User() UserRepo
	UserToken() UserTokenRepo
	Dictionary() DictionaryRepo
	Word() WordRepo
	WordMeaning() WordMeaningRepo
	WordSynonym() WordSynonymRepo
	WordExample() WordExampleRepo
	WordTranslation() WordTranslationRepo
	Translation() TranslationRepo
	Migration() MigrationRepo
	FlashcardDay() FlashcardDayRepo
	FlashcardAttempt() FlashcardAttemptRepo
	FlashcardEngagement() FlashcardEngagementRepo
	FlashcardSession() FlashcardSessionRepo
	UserWordProgress() UserWordProgressRepo
}

type PostgresDB struct {
	conn *pgxpool.Pool
	psql squirrel.StatementBuilderType
}

const MAX_RETRIES = 5

func Connect() DB {
	port, err := strconv.ParseUint(os.Getenv("POSTGRES_PORT"), 10, 16)
	if err != nil {
		slog.Debug("Failed to retrieve postgres port")
	}

	connString := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_HOST"),
		port,
		os.Getenv("POSTGRES_DB"),
	)
	psql := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
	for range MAX_RETRIES {
		conn, err := pgxpool.New(context.Background(), connString)
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
		if strings.ToLower(opts.Sort.Direction) == "desc" {
			orderBy += " DESC"
		}
		builder = builder.OrderBy(orderBy)
	}

	return builder
}
