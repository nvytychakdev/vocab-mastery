package db

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jackc/pgx"
	"github.com/nvytychakdev/vocab-mastery/internal/app/model"
)

type UserRepo interface {
	Create(email string, password string, name string) (uuid.UUID, error)
	CreateOAuth(email string, name string, provider string, providerId string, pictureUrl string, emailVerified bool) (uuid.UUID, error)
	Exists(email string) (bool, error)
	ExistsByProvider(email string, provider string) (bool, error)
	GetByID(id uuid.UUID) (*model.User, error)
	GetByEmail(email string) (*model.User, error)
	GetByEmailWithPwd(email string) (*model.UserWithPwd, error)
	SetEmailConfirmed(id uuid.UUID) error
}

type userRepo struct {
	conn *pgx.Conn
	psql sq.StatementBuilderType
}

func (db *PostgresDB) User() UserRepo {
	return &userRepo{conn: db.conn, psql: db.psql}
}

func (p *userRepo) Create(email string, passwordHash string, name string) (uuid.UUID, error) {
	query, args, err := p.psql.Insert("users").
		Columns("email", "password_hash", "name", "auth_provider").
		Values(email, passwordHash, name, "local").
		Suffix("RETURNING \"id\"").ToSql()

	if err != nil {
		return uuid.Nil, err
	}

	var userId uuid.UUID
	err = p.conn.QueryRow(query, args...).Scan(&userId)
	return userId, err
}

func (p *userRepo) CreateOAuth(email string, name string, provider string, providerId string, pictureUrl string, emailVerified bool) (uuid.UUID, error) {
	query, args, err := p.psql.Insert("users").
		Columns("email", "name", "auth_provider", "auth_provider_user_id", "picture_url", "is_email_confirmed").
		Values(email, name, provider, providerId, pictureUrl, emailVerified).
		Suffix("RETURNING \"id\"").ToSql()

	if err != nil {
		return uuid.Nil, err
	}

	var userId uuid.UUID
	err = p.conn.QueryRow(query, args...).Scan(&userId)
	return userId, err
}

func (p *userRepo) Exists(email string) (bool, error) {
	query, args, err := p.psql.Select("1").
		Prefix("SELECT EXISTS (").
		From("users").Where(sq.Eq{"email": email}).
		Suffix(")").ToSql()

	if err != nil {
		return false, err
	}

	var exists bool
	err = p.conn.QueryRow(query, args...).Scan(&exists)
	return exists, err
}

func (p *userRepo) ExistsByProvider(email string, provider string) (bool, error) {
	query, args, err := p.psql.Select("1").
		Prefix("SELECT EXISTS (").
		From("users").
		Where(
			sq.And{
				sq.Eq{"email": email},
				sq.Eq{"auth_provider": provider},
			},
		).
		Suffix(")").ToSql()

	if err != nil {
		return false, err
	}

	var exists bool
	err = p.conn.QueryRow(query, args...).Scan(&exists)
	return exists, err
}

func (p *userRepo) GetByID(id uuid.UUID) (*model.User, error) {
	query, args, err := p.psql.
		Select("id", "email", "name", "created_at", "picture_url").
		From("users").Where(sq.Eq{"id": id}).ToSql()

	if err != nil {
		return nil, err
	}

	var user model.User
	err = p.conn.QueryRow(query, args...).Scan(&user.ID, &user.Email, &user.Name, &user.CreatedAt, &user.PictureUrl)
	return &user, err
}

func (p *userRepo) GetByEmail(email string) (*model.User, error) {
	query, args, err := p.psql.
		Select("id", "email", "name", "created_at", "picture_url").
		From("users").Where(sq.Eq{"email": email}).ToSql()

	if err != nil {
		return nil, err
	}

	var user model.User
	err = p.conn.QueryRow(query, args...).Scan(&user.ID, &user.Email, &user.Name, &user.CreatedAt, &user.PictureUrl)
	return &user, err
}

func (p *userRepo) GetByEmailWithPwd(email string) (*model.UserWithPwd, error) {
	query, args, err := p.psql.
		Select("id", "email", "password_hash", "is_email_confirmed", "created_at").
		From("users").Where(sq.Eq{"email": email}).ToSql()

	if err != nil {
		return nil, err
	}

	var user model.UserWithPwd
	err = p.conn.QueryRow(query, args...).Scan(&user.ID, &user.Email, &user.Password, &user.IsEmailConfirmed, &user.CreatedAt)
	return &user, err
}

func (p *userRepo) SetEmailConfirmed(id uuid.UUID) error {
	query, args, err := p.psql.
		Update("users").
		Set("is_email_confirmed", "TRUE").
		Where(sq.Eq{"id": id}).ToSql()

	if err != nil {
		return err
	}

	_, err = p.conn.Exec(query, args...)
	return err
}
