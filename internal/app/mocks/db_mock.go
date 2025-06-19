package mocks

import (
	"time"

	"github.com/nvytychakdev/vocab-mastery/internal/app/model"
)

type MockDB struct {
	DeleteSessionByIDFn func(sessionId string) error
}

func (m *MockDB) DeleteSessionByID(sessionId string) error {
	if m.DeleteSessionByIDFn != nil {
		return m.DeleteSessionByIDFn(sessionId)
	}

	return nil
}

func (m *MockDB) CreateSession(userId string, jti string) (string, error) { return "", nil }

func (m *MockDB) SessionExists(id string) (bool, error) { return false, nil }

func (m *MockDB) UpdateSessionJti(id string, jti string) error { return nil }

func (m *MockDB) GetSessionByID(id string) (*model.Session, error) { return nil, nil }

func (m *MockDB) CreateUser(email string, password string, name string) (string, error) {
	return "", nil
}

func (p *MockDB) CreateUserOAuth(email string, name string, provider string, providerId string, pictureUrl string, emailVerified bool) (string, error) {
	return "", nil
}

func (m *MockDB) UserExists(email string) (bool, error) { return false, nil }

func (m *MockDB) GetUserByID(id string) (*model.User, error) { return nil, nil }

func (m *MockDB) GetUserByEmail(email string) (*model.User, error) { return nil, nil }

func (m *MockDB) GetUserWithPawdByEmail(email string) (*model.UserWithPwd, error) { return nil, nil }

func (m *MockDB) SetUserEmailConfirmed(id string) error { return nil }

func (m *MockDB) CreateUserToken(userId string, tokenType string) (string, string, error) {
	return "", "", nil
}

func (m *MockDB) GetNonExpiredUserToken(token string, tokenType string) (string, *time.Time, error) {
	return "", nil, nil
}

func (m *MockDB) SetUserTokenUsed(token string) error { return nil }
