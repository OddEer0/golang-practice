package domain

import (
	"strconv"
	"strings"
)

const (
	ErrUnauthorizedCode   Code = 401
	ErrForbiddenCode      Code = 403
	ErrNotFoundCode       Code = 404
	ErrRequestTimeoutCode Code = 408
	ErrConflictCode       Code = 409

	ErrInternalCode Code = 500

	ErrUnauthorizedMessage = "Unauthorized"
	ErrForbiddenMessage    = "Forbidden"
	ErrInternalMessage     = "Internal error"
	ErrConflictMessage     = "Entity conflict"
	ErrNotFoundMessage     = "Not found"
)

type Error struct {
	Code
	Message string
}

func (e Error) Error() string {
	err := strings.Builder{}
	err.WriteString("error code: ")
	err.WriteString(strconv.Itoa(int(e.Code)))
	err.WriteString(" message: ")
	err.WriteString(e.Message)
	return err.String()
}

func NewErr(code Code, message string) error {
	return Error{
		Code: code,
		Message: message,
	}
}
