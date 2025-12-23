package db

import (
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jackc/pgx"
)

type UserTokenRepo interface {
	Create(userId uuid.UUID, tokenType string) (uuid.UUID, string, error)
	FindNonExpired(token string, tokenType string) (uuid.UUID, *time.Time, error)
	SetUsed(token string) error
}

type userTokenRepo struct {
	conn *pgx.Conn
	psql sq.StatementBuilderType
}

func (db *PostgresDB) UserToken() UserTokenRepo {
	return &userTokenRepo{conn: db.conn, psql: db.psql}
}

func (p *userTokenRepo) Create(userId uuid.UUID, tokenType string) (uuid.UUID, string, error) {
	token, expiresAt := generateEmailConfirmToken()

	query, args, err := p.psql.Insert("user_tokens").
		Columns("user_id", "token", "type", "expires_at").
		Values(userId, token, tokenType, expiresAt).Suffix("RETURNING \"id\"").ToSql()

	if err != nil {
		return uuid.Nil, "", nil
	}

	var userTokenId uuid.UUID
	err = p.conn.QueryRow(query, args...).Scan(&userTokenId)
	return userTokenId, token, err
}

func (p *userTokenRepo) FindNonExpired(token string, tokenType string) (uuid.UUID, *time.Time, error) {
	var userId uuid.UUID
	var usedAt *time.Time

	query, args, err := p.psql.
		Select("user_id", "used_at").From("user_tokens").
		Where(sq.And{
			sq.Eq{"token": token},
			sq.Eq{"type": tokenType},
			sq.Gt{"expires_at": "now()"},
		}).ToSql()

	if err != nil {
		return uuid.Nil, nil, err
	}

	err = p.conn.QueryRow(query, args...).Scan(&userId, &usedAt)
	return userId, usedAt, err
}

func (p *userTokenRepo) SetUsed(token string) error {

	query, args, err := p.psql.
		Update("user_tokens").Set("used_at", time.Now()).
		Where(sq.Eq{"token": token}).ToSql()

	if err != nil {
		return err
	}

	_, err = p.conn.Exec(query, args...)
	return err
}

func generateEmailConfirmToken() (string, time.Time) {
	return uuid.NewString(), time.Now().Add(1 * time.Hour)
}
