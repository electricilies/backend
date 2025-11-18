package domain

import "errors"

var (
	ErrNotFound      = errors.New("not found")
	ErrInvalid       = errors.New("invalid data")
	ErrExists        = errors.New("already exists")
	ErrForbidden     = errors.New("forbidden")
	ErrUnauthorized  = errors.New("unauthorized")
	ErrConflict      = errors.New("conflict")
	ErrUnavailable   = errors.New("service unavailable")
	ErrTimeout       = errors.New("timeout")
	ErrConnection    = errors.New("connection error")
	ErrCanceled      = errors.New("operation canceled")
	ErrInternalDB    = errors.New("internal database error")
	ErrInvalidFormat = errors.New("invalid format")
	ErrServiceError  = errors.New("service error")
	ErrUnknown       = errors.New("unknown error")
)
