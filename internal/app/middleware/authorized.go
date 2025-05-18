package middleware

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/go-chi/render"
	"github.com/nvytychakdev/vocab-mastery/internal/app/auth"
	httpError "github.com/nvytychakdev/vocab-mastery/internal/app/http-error"
)

type contextKey string

const USER_ID_KEY contextKey = "userId"
const SESSION_ID_KEY contextKey = "userId"

func Authorized(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")

		if authHeader == "" {
			render.Render(w, r, httpError.NewErrorResponse(errors.New("access token is missing"), http.StatusUnauthorized))
			return
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			render.Render(w, r, httpError.NewErrorResponse(errors.New("invalid token"), http.StatusUnauthorized))
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		token, claims, err := auth.ParseToken(tokenString)
		if err != nil || !token.Valid {
			render.Render(w, r, httpError.NewErrorResponse(errors.New("invalid token"), http.StatusUnauthorized))
			return
		}

		ctx := context.WithValue(r.Context(), USER_ID_KEY, claims.UserId)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
