package error

import "fmt"

type DomainError struct {
	Code    string
	Message string
	Cause   error
}

func (e *DomainError) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Cause)
	}
	return e.Message
}

func (e *DomainError) Unwrap() error {
	return e.Cause
}

type (
	NotFoundError   struct{ DomainError }
	ConflictError   struct{ DomainError }
	ValidationError struct{ DomainError }
	BadRequestError struct{ DomainError }
)

type (
	InternalError    struct{ DomainError }
	ConnectionError  struct{ DomainError }
	UnavailableError struct{ DomainError }
)

// Constructors
func NewNotFoundError(msg string, cause error) *NotFoundError {
	return &NotFoundError{DomainError{Code: "NOT_FOUND", Message: msg, Cause: cause}}
}

func NewConflictError(msg string, cause error) *ConflictError {
	return &ConflictError{DomainError{Code: "CONFLICT", Message: msg, Cause: cause}}
}

func NewValidationError(msg string, cause error) *ValidationError {
	return &ValidationError{DomainError{Code: "VALIDATION_ERROR", Message: msg, Cause: cause}}
}

func NewBadRequestError(msg string, cause error) *BadRequestError {
	return &BadRequestError{DomainError{Code: "BAD_REQUEST", Message: msg, Cause: cause}}
}

func NewInternalError(msg string, cause error) *InternalError {
	return &InternalError{DomainError{Code: "INTERNAL_ERROR", Message: msg, Cause: cause}}
}

func NewConnectionError(msg string, cause error) *ConnectionError {
	return &ConnectionError{DomainError{Code: "CONNECTION_ERROR", Message: msg, Cause: cause}}
}

func NewUnavailableError(msg string, cause error) *UnavailableError {
	return &UnavailableError{DomainError{Code: "SERVICE_UNAVAILABLE", Message: msg, Cause: cause}}
}
