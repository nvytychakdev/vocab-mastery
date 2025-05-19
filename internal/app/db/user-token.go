package db

import (
	"time"

	"github.com/google/uuid"
)

func CreateUserToken(userId string, tokenType string) (string, error) {
	token, expiresAt := generateEmailConfirmToken()

	const query = `
		INSERT INTO user_tokens (user_id, token, type, expires_at) 
		VALUES ($1, $2, $3, $4) 
		RETURNING id;
	`
	var userTokenId string
	err := DBConn.QueryRow(query, userId, token, tokenType, expiresAt).Scan(&userTokenId)
	return userTokenId, err
}

func GetNonExpiredUserToken(token string, tokenType string) (string, *time.Time, error) {
	var userId string
	var usedAt *time.Time

	query := `
		SELECT user_id, used_at
		FROM user_tokens
		WHERE token = $1 AND type = $2 and expires_at > now()
	`

	err := DBConn.QueryRow(query, token, tokenType).Scan(&userId, &usedAt)
	return userId, usedAt, err
}

func SetUserTokenUsed(token string) error {
	const query = `
		UPDATE user_tokens 
		SET used_at = $2
		WHERE token = $1;
	`

	_, err := DBConn.Exec(query, token, time.Now())
	return err
}

func generateEmailConfirmToken() (string, time.Time) {
	return uuid.NewString(), time.Now().Add(1 * time.Hour)
}
