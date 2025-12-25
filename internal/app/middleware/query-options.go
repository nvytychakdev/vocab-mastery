package middleware

import (
	"context"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/nvytychakdev/vocab-mastery/internal/app/model"
)

func (mw *Middleware) QueryOptionsContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		dictionaryId := r.URL.Query().Get("dictionaryId")
		offset, _ := strconv.Atoi(r.URL.Query().Get("offset"))
		limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
		sortBy := r.URL.Query().Get("sortBy")
		sortDir := r.URL.Query().Get("dir")

		if limit <= 0 || limit > 1000 {
			limit = 20
		}

		if offset < 0 {
			offset = 0
		}

		pgn := &model.Pagination{Offset: offset, Limit: limit}

		fltr := &model.Filters{}
		if len(dictionaryId) > 0 {
			dictionaryUUID, err := uuid.Parse(dictionaryId)
			if err == nil {
				fltr.DictionaryID = &dictionaryUUID
			}
		}

		var srt *model.Sort
		if len(sortBy) > 0 && len(sortDir) > 0 {
			srt = &model.Sort{Field: sortBy, Direction: sortDir}
		}

		ctx := context.WithValue(r.Context(), QUERY_OPTIONS_KEY, &model.QueryOptions{Filters: fltr, Pagination: pgn, Sort: srt})
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetQueryOptionsContext(r *http.Request) *model.QueryOptions {
	queryOptions, ok := r.Context().Value(QUERY_OPTIONS_KEY).(*model.QueryOptions)

	if !ok {
		return &model.QueryOptions{
			Filters:    &model.Filters{},
			Pagination: &model.Pagination{Offset: 0, Limit: 20},
			Sort:       &model.Sort{},
		}
	}

	return queryOptions
}
