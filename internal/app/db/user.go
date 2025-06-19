package db

import (
	"github.com/nvytychakdev/vocab-mastery/internal/app/model"
)

type UserRepository interface {
	CreateUser(email string, password string, name string) (string, error)
	CreateUserOAuth(email string, name string, provider string, providerId string, pictureUrl string, emailVerified bool) (string, error)
	UserExists(email string) (bool, error)
	GetUserByID(id string) (*model.User, error)
	GetUserByEmail(email string) (*model.User, error)
	GetUserWithPawdByEmail(email string) (*model.UserWithPwd, error)
	SetUserEmailConfirmed(id string) error
}

func (p *PostgresDB) CreateUser(email string, passwordHash string, name string) (string, error) {
	const query = `
		INSERT INTO users (email, password_hash, name, auth_provider) 
		VALUES ($1, $2, $3, $4) 
		RETURNING id;
	`
	var userId string
	err := p.conn.QueryRow(query, email, passwordHash, name, "local").Scan(&userId)
	return userId, err
}

func (p *PostgresDB) CreateUserOAuth(email string, name string, provider string, providerId string, pictureUrl string, emailVerified bool) (string, error) {
	const query = `
		INSERT INTO users (email, name, auth_provider, auth_provider_user_id, picture_url, is_email_confirmed) 
		VALUES ($1, $2, $3, $4, $5, $6) 
		RETURNING id;
	`
	var userId string
	err := p.conn.QueryRow(query, email, name, provider, providerId, pictureUrl, emailVerified).Scan(&userId)
	return userId, err
}

func (p *PostgresDB) UserExists(email string) (bool, error) {
	const query = `
		SELECT EXISTS (
			SELECT 1 FROM users WHERE email = $1
		)
	`

	var exists bool
	err := p.conn.QueryRow(query, email).Scan(&exists)
	return exists, err
}

func (p *PostgresDB) GetUserByID(id string) (*model.User, error) {
	const query = `
		SELECT id, email, name, created_at, picture_url
		FROM users
		WHERE id = $1;
	`

	var user model.User
	err := p.conn.QueryRow(query, id).Scan(&user.ID, &user.Email, &user.Name, &user.CreatedAt, &user.PictureUrl)
	return &user, err
}

func (p *PostgresDB) GetUserByEmail(email string) (*model.User, error) {
	const query = `
		SELECT id, email, name, created_at, picture_url
		FROM users
		WHERE email = $1;
	`

	var user model.User
	err := p.conn.QueryRow(query, email).Scan(&user.ID, &user.Email, &user.Name, &user.CreatedAt, &user.PictureUrl)
	return &user, err
}

func (p *PostgresDB) GetUserWithPawdByEmail(email string) (*model.UserWithPwd, error) {
	const query = `
		SELECT id, email, password_hash, is_email_confirmed, created_at
		FROM users
		WHERE email = $1;
	`

	var user model.UserWithPwd
	err := p.conn.QueryRow(query, email).Scan(&user.ID, &user.Email, &user.Password, &user.IsEmailConfirmed, &user.CreatedAt)
	return &user, err
}

func (p *PostgresDB) SetUserEmailConfirmed(id string) error {
	const query = `
		UPDATE users 
		SET is_email_confirmed = TRUE 
		WHERE id = $1;
	`

	_, err := p.conn.Exec(query, id)
	return err
}
