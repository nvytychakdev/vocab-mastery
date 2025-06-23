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

	query, args, err := p.psql.
		Insert("sessions").
		Columns("user_id", "jti", "expires_at").
		Values(userId, jti, expiresAt).
		Suffix("RETURNING \"id\"").ToSql()

	if err != nil {
		return "", err
	}

	var sessionId string
	err = p.conn.QueryRow(query, args...).Scan(&sessionId)
	return sessionId, err
}

func (p *sessionRepo) Exists(id string) (bool, error) {
	query, args, err := p.psql.Select("1").
		Prefix("SELECT EXISTS (").
		From("sessions").
		Where(
			sq.Eq{"id": id},
		).
		Suffix(")").ToSql()

	if err != nil {
		return false, err
	}

	// const query = `
	// 	SELECT EXISTS (
	// 		SELECT 1 FROM sessions WHERE id = $1
	// 	)
	// `

	var exists bool
	err = p.conn.QueryRow(query, args...).Scan(&exists)
	return exists, err
}

func (p *sessionRepo) UpdateJti(id string, jti string) error {
	query, args, err := p.psql.
		Update("sessions").Set("jti", jti).Set("refreshed_at", "now()").
		Where(sq.Eq{"id": id}).ToSql()

	if err != nil {
		return err
	}

	// const query = `
	// 	UPDATE sessions
	// 	SET jti = $2, refreshed_at = now()
	// 	WHERE id = $1;
	// `

	_, err = p.conn.Exec(query, args...)
	return err
}

func (p *sessionRepo) GetByID(id string) (*model.Session, error) {
	query, args, err := p.psql.
		Select("id", "jti", "user_id", "expires_at", "created_at").
		From("sessions").Where(sq.Eq{"id": id}).ToSql()

	if err != nil {
		return nil, err
	}

	// const query = `
	// 	SELECT id, jti, user_id, expires_at, created_at
	// 	FROM sessions
	// 	WHERE id = $1;
	// `

	var session model.Session
	err = p.conn.QueryRow(query, args...).Scan(&session.ID, &session.RefreshTokenID, &session.UserID, &session.ExpiresAt, &session.CreatedAt)
	return &session, err
}

func (p *sessionRepo) DeleteByID(id string) error {

	query, args, err := p.psql.
		Delete("sessions").
		Where(sq.Eq{"id": id}).ToSql()

	if err != nil {
		return err
	}

	// const query = `
	// 	DELETE FROM sessions
	// 	WHERE id = $1;
	// `

	_, err = p.conn.Exec(query, args...)
	return err
}
