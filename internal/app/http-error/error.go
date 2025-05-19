package httpError

import (
	"log/slog"
	"net/http"
)

const (
	// External
	ErrInvalidToken            = 1001
	ErrSessionExpired          = 1002
	ErrTokenRevoked            = 1003
	ErrUserNotFound            = 1004
	ErrPasswordMismatch        = 1005
	ErrUnauthorized            = 1006
	ErrInvalidPayload          = 1007
	ErrUserAlreadyExists       = 1008
	ErrConfirmTokenNotFound    = 1009
	ErrConfirmTokenExpired     = 1010
	ErrConfirmTokenAlreadyUsed = 1011

	// Internal
	ErrInternalServer = 2001
)

var ErrorCodeMessages = map[int16]string{
	ErrInvalidToken:      "Access token is invalid or malformed.",
	ErrSessionExpired:    "Session has expired. Please sign in again.",
	ErrTokenRevoked:      "Token has been revoked or already used. Please sign in again.",
	ErrUserNotFound:      "User does not exist.",
	ErrPasswordMismatch:  "Incorrect password.",
	ErrUnauthorized:      "You are not authorized to perform this action.",
	ErrInvalidPayload:    "Invalid request payload",
	ErrUserAlreadyExists: "User already exists with this email",

	ErrInternalServer: "Internal server error. Please try again later.",
}

type ErrorResponse struct {
	Error error `json:"-"`
	Code  int   `json:"-"`

	StatusText string `json:"status"`
	StatusCode int16  `json:"code,omitempty"`
	ErrorText  string `json:"error,omitempty"`
}

func (e *ErrorResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func NewErrorResponse(code int, statusCode int16, err error) *ErrorResponse {
	slog.Error("Response error", "Err", err, "Code", code, "StatusCode", statusCode)

	return &ErrorResponse{
		Error:      err,
		Code:       code,
		StatusText: http.StatusText(code),
		StatusCode: statusCode,
		ErrorText:  errorCodeMessage(statusCode),
	}
}

func errorCodeMessage(code int16) string {
	return ErrorCodeMessages[code]
}
