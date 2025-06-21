package middleware

import (
	"context"
	"net/http"
	"strconv"

	"github.com/nvytychakdev/vocab-mastery/internal/app/model"
)

func (mw *Middleware) PaginationContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		offset, _ := strconv.Atoi(r.URL.Query().Get("offset"))
		limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))

		if limit <= 0 || limit > 100 {
			limit = 20
		}

		if offset < 0 {
			offset = 0
		}

		p := &model.Pagination{Offset: offset, Limit: limit}
		ctx := context.WithValue(r.Context(), PAGINATION_KEY, p)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetPaginationContext(r *http.Request) *model.Pagination {
	pagination, ok := r.Context().Value(PAGINATION_KEY).(*model.Pagination)
	if !ok {
		return &model.Pagination{Offset: 0, Limit: 20}
	}
	return &model.Pagination{Offset: pagination.Offset, Limit: pagination.Limit}
}
