package domain

import (
	"fmt"
)

type Error = error

const (
	ErrNotFound = "not_found"
	ErrConflict = "conflict"
)

type Err struct {
	Message string
	Cause   error
}

func (e *Err) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Cause)
	}
	return e.Message
}

func (e *Err) Unwrap() error {
	return e.Cause
}

type (
	NotFoundError     struct{ Err }
	ConflictError     struct{ Err }
	ValidationError   struct{ Err }
	BadRequestError   struct{ Err }
	UnauthorizedError struct{ Err }
	ForbiddenError    struct{ Err }
)

type (
	InternalError    struct{ Err }
	ConnectionError  struct{ Err }
	UnavailableError struct{ Err }
)

func NewNotFoundError(msg string, cause error) *NotFoundError {
	return &NotFoundError{Err{Message: msg, Cause: cause}}
}

func NewConflictError(msg string, cause error) *ConflictError {
	return &ConflictError{Err{Message: msg, Cause: cause}}
}

func NewValidationError(msg string, cause error) *ValidationError {
	return &ValidationError{Err{Message: msg, Cause: cause}}
}

func NewBadRequestError(msg string, cause error) *BadRequestError {
	return &BadRequestError{Err{Message: msg, Cause: cause}}
}

func NewInternalError(msg string, cause error) *InternalError {
	return &InternalError{Err{Message: msg, Cause: cause}}
}

func NewConnectionError(msg string, cause error) *ConnectionError {
	return &ConnectionError{Err{Message: msg, Cause: cause}}
}

func NewUnavailableError(msg string, cause error) *UnavailableError {
	return &UnavailableError{Err{Message: msg, Cause: cause}}
}

func NewUnauthorizedError(msg string, cause error) *UnauthorizedError {
	return &UnauthorizedError{Err{Message: msg, Cause: cause}}
}

func NewForbiddenError(msg string, cause error) *ForbiddenError {
	return &ForbiddenError{Err{Message: msg, Cause: cause}}
}
