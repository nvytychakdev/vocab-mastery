package auth

import (
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang-jwt/jwt/v5"
	"github.com/nvytychakdev/vocab-mastery/internal/app/handler/auth"
	"github.com/nvytychakdev/vocab-mastery/internal/app/mocks"
	"github.com/nvytychakdev/vocab-mastery/internal/app/services"
)

func TestSignOut_Success(t *testing.T) {
	reqBody := `{"refreshToken":"test-token"}`
	req := httptest.NewRequest(http.MethodPost, "/auth/sign-out", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	mockAuth := &mocks.MockAuthService{
		ParseTokenFn: func(token string) (*jwt.Token, *services.TokenClaims, error) {
			return &jwt.Token{
				Valid: true,
			}, &services.TokenClaims{SessionId: "Test"}, nil
		},
	}
	mockDB := &mocks.MockDB{}

	handler := &auth.AuthHandler{
		Deps: &services.Deps{
			DB:          mockDB,
			AuthService: mockAuth,
		},
	}

	handler.SignOut(w, req)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, res.StatusCode)
	}

	resBody, _ := io.ReadAll(res.Body)
	if !strings.Contains(string(resBody), `"ok":true`) {
		t.Errorf("expected response body to contain \n\n'ok:true', \n\ngot\n\n %s", resBody)
	}
}

func TestSignOut_InvalidBody(t *testing.T) {
	reqBody := `{"wrongBody":"test-token"}`
	req := httptest.NewRequest(http.MethodPost, "/auth/sign-out", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	mockAuth := &mocks.MockAuthService{
		ParseTokenFn: func(token string) (*jwt.Token, *services.TokenClaims, error) {
			return &jwt.Token{
				Valid: true,
			}, &services.TokenClaims{SessionId: "Test"}, nil
		},
	}
	mockDB := &mocks.MockDB{}

	handler := &auth.AuthHandler{
		Deps: &services.Deps{
			DB:          mockDB,
			AuthService: mockAuth,
		},
	}

	handler.SignOut(w, req)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusBadRequest {
		t.Errorf("expected status %d, got %d", http.StatusBadRequest, res.StatusCode)
	}

	resBody, _ := io.ReadAll(res.Body)
	resExpected := `"status":"Bad Request","code":1007,"error":"Invalid request payload"`
	if !strings.Contains(string(resBody), resExpected) {
		t.Errorf("expected response body to contain \n\n'%s', \n\ngot\n\n %s", resExpected, resBody)
	}
}

func TestSignOut_ParseError(t *testing.T) {
	reqBody := `{"refreshToken":"test-token"}`
	req := httptest.NewRequest(http.MethodPost, "/auth/sign-out", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	mockAuth := &mocks.MockAuthService{
		ParseTokenFn: func(token string) (*jwt.Token, *services.TokenClaims, error) {
			return nil, nil, errors.New("error parsing the token")
		},
	}
	mockDB := &mocks.MockDB{}

	handler := &auth.AuthHandler{
		Deps: &services.Deps{
			DB:          mockDB,
			AuthService: mockAuth,
		},
	}

	handler.SignOut(w, req)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusUnauthorized {
		t.Errorf("expected status %d, got %d", http.StatusUnauthorized, res.StatusCode)
	}

	resBody, _ := io.ReadAll(res.Body)
	resExpected := `"status":"Unauthorized","code":1006,"error":"You are not authorized to perform this action."`
	if !strings.Contains(string(resBody), resExpected) {
		t.Errorf("expected response body to contain \n\n'%s', \n\ngot\n\n %s", resExpected, resBody)
	}
}

func TestSignOut_InvalidToken(t *testing.T) {
	reqBody := `{"refreshToken":"test-token"}`
	req := httptest.NewRequest(http.MethodPost, "/auth/sign-out", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	mockAuth := &mocks.MockAuthService{
		ParseTokenFn: func(token string) (*jwt.Token, *services.TokenClaims, error) {
			return &jwt.Token{Valid: false}, nil, nil
		},
	}
	mockDB := &mocks.MockDB{}

	handler := &auth.AuthHandler{
		Deps: &services.Deps{
			DB:          mockDB,
			AuthService: mockAuth,
		},
	}

	handler.SignOut(w, req)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusUnauthorized {
		t.Errorf("expected status %d, got %d", http.StatusUnauthorized, res.StatusCode)
	}

	resBody, _ := io.ReadAll(res.Body)
	resExpected := `"status":"Unauthorized","code":1006,"error":"You are not authorized to perform this action."`
	if !strings.Contains(string(resBody), resExpected) {
		t.Errorf("expected response body to contain \n\n'%s', \n\ngot\n\n %s", resExpected, resBody)
	}
}

func TestSignOut_DeleteSessionError(t *testing.T) {
	reqBody := `{"refreshToken":"test-token"}`
	req := httptest.NewRequest(http.MethodPost, "/auth/sign-out", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	mockAuth := &mocks.MockAuthService{
		ParseTokenFn: func(token string) (*jwt.Token, *services.TokenClaims, error) {
			return &jwt.Token{Valid: true}, &services.TokenClaims{SessionId: "Test"}, nil
		},
	}
	mockDB := &mocks.MockDB{
		DeleteSessionByIDFn: func(sessionId string) error {
			return errors.New("Not able to delete session")
		},
	}

	handler := &auth.AuthHandler{
		Deps: &services.Deps{
			DB:          mockDB,
			AuthService: mockAuth,
		},
	}

	handler.SignOut(w, req)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusInternalServerError {
		t.Errorf("expected status %d, got %d", http.StatusInternalServerError, res.StatusCode)
	}

	resBody, _ := io.ReadAll(res.Body)
	resExpected := `"status":"Internal Server Error","code":2001,"error":"Internal server error. Please try again later."`
	if !strings.Contains(string(resBody), resExpected) {
		t.Errorf("expected response body to contain \n\n'%s', \n\ngot\n\n %s", resExpected, resBody)
	}
}
