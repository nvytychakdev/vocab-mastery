package session

import (
	"time"

	"github.com/nvytychakdev/vocab-mastery/internal/app/database"
)

type Session struct {
	ID             string    `json:"id,omitempty"`
	UserID         string    `json:"userId"`
	RefreshTokenID string    `json:"jti,omitempty"`
	ExpiresAt      time.Time `json:"expiresAt,omitempty"`
	CreatedAt      time.Time `json:"createdAt,omitempty"`
}

func CraeteSession(userId string, jti string) (string, error) {
	expiresAt := time.Now().Add(90 * 24 * time.Hour)
	const query = `
		INSERT INTO sessions (user_id, jti, expires_at) 
		VALUES ($1, $2, $3) 
		RETURNING id;
	`

	var sessionId string
	err := database.DBConn.QueryRow(query, userId, jti, expiresAt).Scan(&sessionId)
	return sessionId, err
}

func SessionExists(id string) (bool, error) {
	const query = `
		SELECT EXISTS (
			SELECT 1 FROM sessions WHERE id = $1
		)
	`

	var exists bool
	err := database.DBConn.QueryRow(query, id).Scan(&exists)
	return exists, err
}

func UpdateSessionJti(id string, jti string) error {
	const query = `
		UPDATE sessions 
		SET jti = $2
		WHERE id = $1;
	`

	_, err := database.DBConn.Exec(query, id, jti)
	return err
}

func GetSessionByID(id string) (*Session, error) {
	const query = `
		SELECT id, jti, user_id, expires_at, created_at
		FROM sessions 
		WHERE id = $1;
	`

	var session Session
	err := database.DBConn.QueryRow(query, id).Scan(&session.ID, &session.RefreshTokenID, &session.UserID, &session.ExpiresAt, &session.CreatedAt)
	return &session, err
}

func DeleteSessionByID(id string) error {
	const query = `
		DELETE FROM sessions
		WHERE id = $1;
	`

	_, err := database.DBConn.Exec(query, id)
	return err
}
