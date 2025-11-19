package domain

import "errors"

var (
	ErrConflict     = errors.New("conflict")
	ErrExists       = errors.New("already exists")
	ErrForbidden    = errors.New("forbidden")
	ErrInvalid      = errors.New("invalid data")
	ErrNotFound     = errors.New("not found")
	ErrServiceError = errors.New("service error")
	ErrTimeout      = errors.New("timeout")
	ErrUnavailable  = errors.New("service unavailable")
	ErrUnknown      = errors.New("unknown error")
)
