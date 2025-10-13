package error

import (
	"errors"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"

	domainerror "backend/internal/domain/error"
)

func ToDomainError(err error) error {
	if err == nil {
		return nil
	}

	if errors.Is(err, pgx.ErrNoRows) {
		return domainerror.NewNotFoundError("record not found", err)
	}

	var pgErr *pgconn.PgError
	if !errors.As(err, &pgErr) {
		return domainerror.NewInternalError("unexpected database error", err)
	}

	code := pgErr.Code

	if pgerrcode.IsConnectionException(code) {
		return domainerror.NewConnectionError(
			"database connection failed",
			err,
		)
	}

	switch code {
	case pgerrcode.UniqueViolation:
		return domainerror.NewConflictError(
			pgErr.ConstraintName+" already exists",
			err,
		)

	case pgerrcode.ForeignKeyViolation:
		return domainerror.NewBadRequestError(
			"referenced resource does not exist",
			err,
		)

	case pgerrcode.NotNullViolation:
		field := pgErr.ColumnName
		if field == "" {
			field = "required field"
		}
		return domainerror.NewValidationError(
			field+" is required",
			err,
		)

	case pgerrcode.CheckViolation:
		return domainerror.NewValidationError(
			"validation constraint failed: "+pgErr.ConstraintName,
			err,
		)

	case pgerrcode.InvalidTextRepresentation:
		return domainerror.NewValidationError(
			"invalid data format: "+pgErr.Message,
			err,
		)

	case pgerrcode.StringDataRightTruncationDataException:
		field := pgErr.ColumnName
		if field == "" {
			field = "field"
		}
		return domainerror.NewValidationError(
			field+" exceeds maximum length",
			err,
		)

	case pgerrcode.UndefinedTable:
		return domainerror.NewInternalError(
			"database table not found",
			err,
		)

	case pgerrcode.UndefinedColumn:
		return domainerror.NewInternalError(
			"database column not found",
			err,
		)

	case pgerrcode.DeadlockDetected:
		return domainerror.NewConflictError(
			"operation conflict, please retry",
			err,
		)

	case pgerrcode.SerializationFailure:
		return domainerror.NewConflictError(
			"concurrent modification detected, please retry",
			err,
		)

	case pgerrcode.TooManyConnections:
		return domainerror.NewUnavailableError(
			"service temporarily unavailable, please try again",
			err,
		)

	case pgerrcode.QueryCanceled:
		return domainerror.NewInternalError(
			"operation timed out",
			err,
		)

	default:
		return domainerror.NewInternalError(
			"database error occurred",
			err,
		)
	}
}
