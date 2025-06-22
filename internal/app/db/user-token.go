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

	const query = `
		INSERT INTO user_tokens (user_id, token, type, expires_at) 
		VALUES ($1, $2, $3, $4) 
		RETURNING id;
	`
	var userTokenId string
	err := p.conn.QueryRow(query, userId, token, tokenType, expiresAt).Scan(&userTokenId)
	return userTokenId, token, err
}

func (p *userTokenRepo) FindNonExpired(token string, tokenType string) (string, *time.Time, error) {
	var userId string
	var usedAt *time.Time

	query := `
		SELECT user_id, used_at
		FROM user_tokens
		WHERE token = $1 AND type = $2 and expires_at > now()
	`

	err := p.conn.QueryRow(query, token, tokenType).Scan(&userId, &usedAt)
	return userId, usedAt, err
}

func (p *userTokenRepo) SetUsed(token string) error {
	const query = `
		UPDATE user_tokens 
		SET used_at = $2
		WHERE token = $1;
	`

	_, err := p.conn.Exec(query, token, time.Now())
	return err
}

func generateEmailConfirmToken() (string, time.Time) {
	return uuid.NewString(), time.Now().Add(1 * time.Hour)
}
