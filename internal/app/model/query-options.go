package model

type QueryOptions struct {
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
