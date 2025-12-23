package db

import (
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jackc/pgx"
	"github.com/nvytychakdev/vocab-mastery/internal/app/model"
)

type SessionRepo interface {
	Create(userID uuid.UUID, jti string) (uuid.UUID, error)
	Exists(id uuid.UUID) (bool, error)
	UpdateJti(id uuid.UUID, jti string) error
	GetByID(id uuid.UUID) (*model.Session, error)
	DeleteByID(id uuid.UUID) error
}

type sessionRepo struct {
	conn *pgx.Conn
	psql sq.StatementBuilderType
}

func (db *PostgresDB) Session() SessionRepo {
	return &sessionRepo{conn: db.conn, psql: db.psql}
}

func (p *sessionRepo) Create(userId uuid.UUID, jti string) (uuid.UUID, error) {
	expiresAt := time.Now().Add(90 * 24 * time.Hour)

	query, args, err := p.psql.
		Insert("sessions").
		Columns("user_id", "jti", "expires_at").
		Values(userId, jti, expiresAt).
		Suffix("RETURNING \"id\"").ToSql()

	if err != nil {
		return uuid.Nil, err
	}

	var sessionId uuid.UUID
	err = p.conn.QueryRow(query, args...).Scan(&sessionId)
	return sessionId, err
}

func (p *sessionRepo) Exists(id uuid.UUID) (bool, error) {
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

	var exists bool
	err = p.conn.QueryRow(query, args...).Scan(&exists)
	return exists, err
}

func (p *sessionRepo) UpdateJti(id uuid.UUID, jti string) error {
	query, args, err := p.psql.
		Update("sessions").Set("jti", jti).Set("refreshed_at", "now()").
		Where(sq.Eq{"id": id}).ToSql()

	if err != nil {
		return err
	}

	_, err = p.conn.Exec(query, args...)
	return err
}

func (p *sessionRepo) GetByID(id uuid.UUID) (*model.Session, error) {
	query, args, err := p.psql.
		Select("id", "jti", "user_id", "expires_at", "created_at").
		From("sessions").Where(sq.Eq{"id": id}).ToSql()

	if err != nil {
		return nil, err
	}

	var session model.Session
	err = p.conn.QueryRow(query, args...).Scan(&session.ID, &session.RefreshTokenID, &session.UserID, &session.ExpiresAt, &session.CreatedAt)
	return &session, err
}

func (p *sessionRepo) DeleteByID(id uuid.UUID) error {

	query, args, err := p.psql.
		Delete("sessions").
		Where(sq.Eq{"id": id}).ToSql()

	if err != nil {
		return err
	}

	_, err = p.conn.Exec(query, args...)
	return err
}
