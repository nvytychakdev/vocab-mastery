package model

import "github.com/google/uuid"

type QueryOptions struct {
	Filters    *Filters
	Pagination *Pagination
	Sort       *Sort
}

type Sort struct {
	Field     string
	Direction string
}

type Pagination struct {
	Offset int
	Limit  int
}

type PaginationResponse struct {
	Total  int `json:"total"`
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

type Filters struct {
	DictionaryID *uuid.UUID
}
