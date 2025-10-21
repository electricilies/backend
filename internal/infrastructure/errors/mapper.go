package errors

import (
	"backend/internal/domain"
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"

	"github.com/Nerzal/gocloak"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

func ToDomainErrorFromPostgres(err error) error {
	if err == nil {
		return nil
	}

	if errors.Is(err, pgx.ErrNoRows) {
		return domain.NewNotFoundError("record not found", err)
	}

	var pgErr *pgconn.PgError
	if !errors.As(err, &pgErr) {
		return domain.NewInternalError("unexpected database error", err)
	}

	code := pgErr.Code

	if pgerrcode.IsConnectionException(code) {
		return domain.NewConnectionError(
			"database connection failed",
			err,
		)
	}

	switch code {
	case pgerrcode.UniqueViolation:
		return domain.NewConflictError(
			pgErr.ConstraintName+" already exists",
			err,
		)

	case pgerrcode.ForeignKeyViolation:
		return domain.NewBadRequestError(
			"referenced resource does not exist",
			err,
		)

	case pgerrcode.NotNullViolation:
		field := pgErr.ColumnName
		if field == "" {
			field = "required field"
		}
		return domain.NewValidationError(
			field+" is required",
			err,
		)

	case pgerrcode.CheckViolation:
		return domain.NewValidationError(
			"validation constraint failed: "+pgErr.ConstraintName,
			err,
		)

	case pgerrcode.InvalidTextRepresentation:
		return domain.NewValidationError(
			"invalid data format: "+pgErr.Message,
			err,
		)

	case pgerrcode.StringDataRightTruncationDataException:
		field := pgErr.ColumnName
		if field == "" {
			field = "field"
		}
		return domain.NewValidationError(
			field+" exceeds maximum length",
			err,
		)

	case pgerrcode.UndefinedTable:
		return domain.NewInternalError(
			"database table not found",
			err,
		)

	case pgerrcode.UndefinedColumn:
		return domain.NewInternalError(
			"database column not found",
			err,
		)

	case pgerrcode.DeadlockDetected:
		return domain.NewConflictError(
			"operation conflict, please retry",
			err,
		)

	case pgerrcode.SerializationFailure:
		return domain.NewConflictError(
			"concurrent modification detected, please retry",
			err,
		)

	case pgerrcode.TooManyConnections:
		return domain.NewUnavailableError(
			"service temporarily unavailable, please try again",
			err,
		)

	case pgerrcode.QueryCanceled:
		return domain.NewInternalError(
			"operation timed out",
			err,
		)

	default:
		return domain.NewInternalError(
			"database error occurred",
			err,
		)
	}
}

func ToDomainErrorFromGoCloak(err error) error {
	if err == nil {
		return nil
	}

	var apiErr gocloak.APIError
	if errors.As(err, &apiErr) {
		switch apiErr.Code {
		case http.StatusBadRequest:
			return domain.NewValidationError(apiErr.Message, err)

		case http.StatusUnauthorized:
			return domain.NewUnauthorizedError("invalid credentials", err)

		case http.StatusForbidden:
			return domain.NewForbiddenError("access denied", err)

		case http.StatusNotFound:
			return domain.NewNotFoundError("resource not found", err)

		case http.StatusConflict:
			return domain.NewConflictError("resource already exists", err)

		case http.StatusGatewayTimeout, http.StatusServiceUnavailable:
			return domain.NewUnavailableError("keycloak service unavailable", err)

		default:
			return domain.NewInternalError(fmt.Sprintf("keycloak error: %s", apiErr.Message), err)
		}
	}

	if errors.Is(err, context.DeadlineExceeded) {
		return domain.NewUnavailableError("keycloak request timed out", err)
	}

	var netErr net.Error
	if errors.As(err, &netErr) {
		if netErr.Timeout() {
			return domain.NewUnavailableError("keycloak network timeout", err)
		}
		return domain.NewConnectionError("keycloak connection failed", err)
	}

	return domain.NewInternalError("unexpected keycloak error", err)
}
