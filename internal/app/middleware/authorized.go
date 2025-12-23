package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/go-chi/render"
	"github.com/google/uuid"
	httpError "github.com/nvytychakdev/vocab-mastery/internal/app/http-error"
)

func (m *Middleware) Authorized(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")

		tokenPrefix := "Bearer "
		if authHeader == "" || !strings.HasPrefix(authHeader, tokenPrefix) {
			render.Render(w, r, httpError.NewErrorResponse(http.StatusUnauthorized, httpError.ErrInvalidToken, nil))
			return
		}

		tokenString := strings.TrimPrefix(authHeader, tokenPrefix)

		token, claims, err := m.Deps.AuthService.ParseToken(tokenString)
		if err != nil || !token.Valid {
			render.Render(w, r, httpError.NewErrorResponse(http.StatusUnauthorized, httpError.ErrInvalidToken, nil))
			return
		}

		ctx := context.WithValue(r.Context(), USER_ID_KEY, claims.UserId)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetAuthorizedUserId(r *http.Request) uuid.UUID {
	return r.Context().Value(USER_ID_KEY).(uuid.UUID)
}
