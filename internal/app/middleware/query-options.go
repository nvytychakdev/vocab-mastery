package middleware

import (
	"context"
	"net/http"
	"strconv"

	"github.com/nvytychakdev/vocab-mastery/internal/app/model"
)

func (mw *Middleware) QueryOptionsContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		offset, _ := strconv.Atoi(r.URL.Query().Get("offset"))
		limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
		sortBy := r.URL.Query().Get("sortBy")
		sortDir := r.URL.Query().Get("dir")

		if limit <= 0 || limit > 100 {
			limit = 20
		}

		if offset < 0 {
			offset = 0
		}

		pgn := &model.Pagination{Offset: offset, Limit: limit}

		var srt *model.Sort
		if len(sortBy) > 0 && len(sortDir) > 0 {
			srt = &model.Sort{Field: sortBy, Direction: sortDir}
		}

		ctx := context.WithValue(r.Context(), QUERY_OPTIONS_KEY, &model.QueryOptions{Pagination: pgn, Sort: srt})
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetQueryOptionsContext(r *http.Request) *model.QueryOptions {
	queryOptions, ok := r.Context().Value(QUERY_OPTIONS_KEY).(*model.QueryOptions)

	if !ok {
		return &model.QueryOptions{
			Pagination: &model.Pagination{Offset: 0, Limit: 20},
			Sort:       &model.Sort{},
		}
	}

	return queryOptions
}
