package domain

import "fmt"

type Err struct {
	Code    string
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
	NotFoundError   struct{ Err }
	ConflictError   struct{ Err }
	ValidationError struct{ Err }
	BadRequestError struct{ Err }
)

type (
	InternalError    struct{ Err }
	ConnectionError  struct{ Err }
	UnavailableError struct{ Err }
)

func NewNotFoundError(msg string, cause error) *NotFoundError {
	return &NotFoundError{Err{Code: "NOT_FOUND", Message: msg, Cause: cause}}
}

func NewConflictError(msg string, cause error) *ConflictError {
	return &ConflictError{Err{Code: "CONFLICT", Message: msg, Cause: cause}}
}

func NewValidationError(msg string, cause error) *ValidationError {
	return &ValidationError{Err{Code: "VALIDATION_ERROR", Message: msg, Cause: cause}}
}

func NewBadRequestError(msg string, cause error) *BadRequestError {
	return &BadRequestError{Err{Code: "BAD_REQUEST", Message: msg, Cause: cause}}
}

func NewInternalError(msg string, cause error) *InternalError {
	return &InternalError{Err{Code: "INTERNAL_ERROR", Message: msg, Cause: cause}}
}

func NewConnectionError(msg string, cause error) *ConnectionError {
	return &ConnectionError{Err{Code: "CONNECTION_ERROR", Message: msg, Cause: cause}}
}

func NewUnavailableError(msg string, cause error) *UnavailableError {
	return &UnavailableError{Err{Code: "SERVICE_UNAVAILABLE", Message: msg, Cause: cause}}
}
