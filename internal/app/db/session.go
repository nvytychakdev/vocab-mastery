package db

import (
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx"
	"github.com/nvytychakdev/vocab-mastery/internal/app/model"
)

type SessionRepo interface {
	Create(userID string, jti string) (string, error)
	Exists(id string) (bool, error)
	UpdateJti(id string, jti string) error
	GetByID(id string) (*model.Session, error)
	DeleteByID(id string) error
}

type sessionRepo struct {
	conn *pgx.Conn
	psql sq.StatementBuilderType
}

func (db *PostgresDB) Session() SessionRepo {
	return &sessionRepo{conn: db.conn, psql: db.psql}
}

func (p *sessionRepo) Create(userId string, jti string) (string, error) {
	expiresAt := time.Now().Add(90 * 24 * time.Hour)
	const query = `
		INSERT INTO sessions (user_id, jti, expires_at) 
		VALUES ($1, $2, $3) 
		RETURNING id;
	`

	var sessionId string
	err := p.conn.QueryRow(query, userId, jti, expiresAt).Scan(&sessionId)
	return sessionId, err
}

func (p *sessionRepo) Exists(id string) (bool, error) {
	const query = `
		SELECT EXISTS (
			SELECT 1 FROM sessions WHERE id = $1
		)
	`

	var exists bool
	err := p.conn.QueryRow(query, id).Scan(&exists)
	return exists, err
}

func (p *sessionRepo) UpdateJti(id string, jti string) error {
	const query = `
		UPDATE sessions 
		SET jti = $2, refreshed_at = now()
		WHERE id = $1;
	`

	_, err := p.conn.Exec(query, id, jti)
	return err
}

func (p *sessionRepo) GetByID(id string) (*model.Session, error) {
	const query = `
		SELECT id, jti, user_id, expires_at, created_at
		FROM sessions 
		WHERE id = $1;
	`

	var session model.Session
	err := p.conn.QueryRow(query, id).Scan(&session.ID, &session.RefreshTokenID, &session.UserID, &session.ExpiresAt, &session.CreatedAt)
	return &session, err
}

func (p *sessionRepo) DeleteByID(id string) error {
	const query = `
		DELETE FROM sessions
		WHERE id = $1;
	`

	_, err := p.conn.Exec(query, id)
	return err
}
