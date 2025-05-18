package httpError

import (
	"net/http"
)

type ErrorResponse struct {
	Error error `json:"-"`
	Code  int   `json:"-"`

	StatusText string `json:"status"`
	StatusCode int64  `json:"code,omitempty"`
	ErrorText  string `json:"error,omitempty"`
}

func (e *ErrorResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func NewErrorResponse(err error, code int) *ErrorResponse {
	return &ErrorResponse{
		Error:      err,
		Code:       code,
		StatusText: http.StatusText(code),
		ErrorText:  err.Error(),
	}
}
