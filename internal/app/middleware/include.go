package middleware

import (
	"context"
	"net/http"
	"strings"
)

type IncludeSet map[string]bool

func (mw *Middleware) IncludeContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		includeSet := make(IncludeSet)
		includeRaw := r.URL.Query().Get("include")

		for _, item := range strings.Split(includeRaw, ",") {
			include := strings.TrimSpace(item)
			if item != "" {
				includeSet[include] = true
			}
		}

		ctx := context.WithValue(r.Context(), INCLUDE_KEY, includeSet)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetIncludeContext(r *http.Request) IncludeSet {
	include, ok := r.Context().Value(INCLUDE_KEY).(IncludeSet)
	if !ok {
		return nil
	}

	return include
}
