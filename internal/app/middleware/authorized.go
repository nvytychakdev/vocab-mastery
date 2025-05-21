package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/go-chi/render"
	httpError "github.com/nvytychakdev/vocab-mastery/internal/app/http-error"
	"github.com/nvytychakdev/vocab-mastery/internal/app/services"
)

type contextKey string

const USER_ID_KEY contextKey = "userId"
const SESSION_ID_KEY contextKey = "userId"

type Middleware struct {
	Deps *services.Deps
}

func NewMiddleware(deps *services.Deps) *Middleware {
	return &Middleware{Deps: deps}
}

func (m *Middleware) Authorized(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")

		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			render.Render(w, r, httpError.NewErrorResponse(http.StatusUnauthorized, httpError.ErrInvalidToken, nil))
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		token, claims, err := m.Deps.TokenService.ParseToken(tokenString)
		if err != nil || !token.Valid {
			render.Render(w, r, httpError.NewErrorResponse(http.StatusUnauthorized, httpError.ErrInvalidToken, nil))
			return
		}

		ctx := context.WithValue(r.Context(), USER_ID_KEY, claims.UserId)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
