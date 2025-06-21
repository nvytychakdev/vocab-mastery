package model

type Pagination struct {
	Offset int
	Limit  int
}

type PaginationResponse struct {
	Total  int `json:"total"`
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}
