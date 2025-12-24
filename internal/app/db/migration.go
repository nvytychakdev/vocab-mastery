package db

import (
	"log/slog"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
)

type MigrationRepo interface {
	Check(migrationId int) (bool, error)
	Run(migrationId int, migration func() error) error
}

type migrationRepo struct {
	conn *pgxpool.Pool
	psql sq.StatementBuilderType
}

func (db *PostgresDB) Migration() MigrationRepo {
	return &migrationRepo{conn: db.conn, psql: db.psql}
}

func (r *migrationRepo) Check(version int) (bool, error) {
	slog.Info("[DATA_MIGRATION] Check migration...", "version", version)
	return false, nil
}

func (r *migrationRepo) Run(version int, migration func() error) error {
	slog.Info("[DATA_MIGRATION] Run migration...", "version", version)
	err := migration()
	if err != nil {
		slog.Error("[DATA_MIGRATION] Run error", "version", version)
		return err
	}

	// query, args, err := r.psql.
	// 	Insert("data_migrations").
	// 	Columns("version", "dirty").
	// 	Values(version, false).
	// 	Suffix("RETURNING \"version\"").ToSql()

	// if err != nil {
	// 	slog.Error("[DATA_MIGRATION] Version update error", "version", version)
	// 	return err
	// }

	// var dictionaryId string
	// err = r.conn.QueryRow(query, args...).Scan(&dictionaryId)
	return err
}
