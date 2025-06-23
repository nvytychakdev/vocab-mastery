package db

import (
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jackc/pgx"
)

type UserTokenRepo interface {
	Create(userId string, tokenType string) (string, string, error)
	FindNonExpired(token string, tokenType string) (string, *time.Time, error)
	SetUsed(token string) error
}

type userTokenRepo struct {
	conn *pgx.Conn
	psql sq.StatementBuilderType
}

func (db *PostgresDB) UserToken() UserTokenRepo {
	return &userTokenRepo{conn: db.conn, psql: db.psql}
}

func (p *userTokenRepo) Create(userId string, tokenType string) (string, string, error) {
	token, expiresAt := generateEmailConfirmToken()

	query, args, err := p.psql.Insert("user_tokens").
		Columns("user_id", "token", "type", "expires_at").
		Values(userId, token, tokenType, expiresAt).Suffix("RETURNING \"id\"").ToSql()

	if err != nil {
		return "", "", nil
	}

	var userTokenId string
	err = p.conn.QueryRow(query, args...).Scan(&userTokenId)
	return userTokenId, token, err
}

func (p *userTokenRepo) FindNonExpired(token string, tokenType string) (string, *time.Time, error) {
	var userId string
	var usedAt *time.Time

	query, args, err := p.psql.
		Select("userId", "used_at").From("user_tokens").
		Where(sq.And{
			sq.Eq{"token": token},
			sq.Eq{"type": tokenType},
			sq.Eq{"expires_at": "now()"},
		}).ToSql()

	if err != nil {
		return "", nil, err
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
