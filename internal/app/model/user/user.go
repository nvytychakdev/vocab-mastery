package user

import (
	"time"

	"github.com/nvytychakdev/vocab-mastery/internal/app/database"
	"github.com/nvytychakdev/vocab-mastery/internal/app/utils"
)

type UserData struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

type User struct {
	UserData
	ID        string    `json:"id,omitempty"`
	CreatedAt time.Time `json:"createdAt"`
}

type UserWithPwd struct {
	User
	Password string `json:"-"`
}

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
	err = database.DBConn.QueryRow(query, email, passwordHash, name).Scan(&userId)
	return userId, err
}

func UserExists(email string) (bool, error) {
	const query = `
		SELECT EXISTS (
			SELECT 1 FROM users WHERE email = $1
		)
	`

	var exists bool
	err := database.DBConn.QueryRow(query, email).Scan(&exists)
	return exists, err
}

func GetUserByID(id string) (*User, error) {
	const query = `
		SELECT id, email, name, created_at
		FROM users
		WHERE id = $1;
	`

	var user User
	err := database.DBConn.QueryRow(query, id).Scan(&user.ID, &user.Email, &user.Name, &user.CreatedAt)
	return &user, err
}

func GetUserWithPawdByEmail(email string) (*UserWithPwd, error) {
	const query = `
		SELECT id, email, password_hash, created_at
		FROM users
		WHERE email = $1;
	`

	var user UserWithPwd
	err := database.DBConn.QueryRow(query, email).Scan(&user.ID, &user.Email, &user.Password, &user.CreatedAt)
	return &user, err
}
