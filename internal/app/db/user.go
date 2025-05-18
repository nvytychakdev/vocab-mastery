package db

import (
	"github.com/nvytychakdev/vocab-mastery/internal/app/model"
	"github.com/nvytychakdev/vocab-mastery/internal/app/utils"
)

func CreateUser(email string, password string, name string) (string, error) {
	passwordHash, err := utils.HashPassword(password)
	if err != nil {
		return "", err
	}

	const query = `
		INSERT INTO users (email, password_hash, name) 
		VALUES ($1, $2, $3) 
		RETURNING id;
	`
	var userId string
	err = DBConn.QueryRow(query, email, passwordHash, name).Scan(&userId)
	return userId, err
}

func UserExists(email string) (bool, error) {
	const query = `
		SELECT EXISTS (
			SELECT 1 FROM users WHERE email = $1
		)
	`

	var exists bool
	err := DBConn.QueryRow(query, email).Scan(&exists)
	return exists, err
}

func GetUserByID(id string) (*model.User, error) {
	const query = `
		SELECT id, email, name, created_at
		FROM users
		WHERE id = $1;
	`

	var user model.User
	err := DBConn.QueryRow(query, id).Scan(&user.ID, &user.Email, &user.Name, &user.CreatedAt)
	return &user, err
}

func GetUserWithPawdByEmail(email string) (*model.UserWithPwd, error) {
	const query = `
		SELECT id, email, password_hash, created_at
		FROM users
		WHERE email = $1;
	`

	var user model.UserWithPwd
	err := DBConn.QueryRow(query, email).Scan(&user.ID, &user.Email, &user.Password, &user.CreatedAt)
	return &user, err
}
