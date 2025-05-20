package db

import (
	"time"

	"github.com/nvytychakdev/vocab-mastery/internal/app/model"
)

func CraeteSession(userId string, jti string) (string, error) {
	expiresAt := time.Now().Add(90 * 24 * time.Hour)
	const query = `
		INSERT INTO sessions (user_id, jti, expires_at) 
		VALUES ($1, $2, $3) 
		RETURNING id;
	`

	var sessionId string
	err := DBConn.QueryRow(query, userId, jti, expiresAt).Scan(&sessionId)
	return sessionId, err
}

func SessionExists(id string) (bool, error) {
	const query = `
		SELECT EXISTS (
			SELECT 1 FROM sessions WHERE id = $1
		)
	`

	var exists bool
	err := DBConn.QueryRow(query, id).Scan(&exists)
	return exists, err
}

func UpdateSessionJti(id string, jti string) error {
	const query = `
		UPDATE sessions 
		SET jti = $2, refreshed_at = now()
		WHERE id = $1;
	`

	_, err := DBConn.Exec(query, id, jti)
	return err
}

func GetSessionByID(id string) (*model.Session, error) {
	const query = `
		SELECT id, jti, user_id, expires_at, created_at
		FROM sessions 
		WHERE id = $1;
	`

	var session model.Session
	err := DBConn.QueryRow(query, id).Scan(&session.ID, &session.RefreshTokenID, &session.UserID, &session.ExpiresAt, &session.CreatedAt)
	return &session, err
}

func DeleteSessionByID(id string) error {
	const query = `
		DELETE FROM sessions
		WHERE id = $1;
	`

	_, err := DBConn.Exec(query, id)
	return err
}
