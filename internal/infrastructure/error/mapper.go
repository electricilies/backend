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

	if code == pgerrcode.UniqueViolation {
		return domainerror.NewConflictError(
			pgErr.ConstraintName+" already exists",
			err,
		)
	}

	if code == pgerrcode.ForeignKeyViolation {
		return domainerror.NewBadRequestError(
			"referenced resource does not exist",
			err,
		)
	}

	if code == pgerrcode.NotNullViolation {
		field := pgErr.ColumnName
		if field == "" {
			field = "required field"
		}
		return domainerror.NewValidationError(
			field+" is required",
			err,
		)
	}

	if code == pgerrcode.CheckViolation {
		return domainerror.NewValidationError(
			"validation constraint failed: "+pgErr.ConstraintName,
			err,
		)
	}

	if code == pgerrcode.InvalidTextRepresentation {
		return domainerror.NewValidationError(
			"invalid data format: "+pgErr.Message,
			err,
		)
	}

	if code == pgerrcode.StringDataRightTruncationDataException {
		field := pgErr.ColumnName
		if field == "" {
			field = "field"
		}
		return domainerror.NewValidationError(
			field+" exceeds maximum length",
			err,
		)
	}

	if code == pgerrcode.UndefinedTable {
		return domainerror.NewInternalError(
			"database table not found",
			err,
		)
	}

	if code == pgerrcode.UndefinedColumn {
		return domainerror.NewInternalError(
			"database column not found",
			err,
		)
	}

	if pgerrcode.IsConnectionException(code) {
		return domainerror.NewConnectionError(
			"database connection failed",
			err,
		)
	}

	if code == pgerrcode.DeadlockDetected {
		return domainerror.NewConflictError(
			"operation conflict, please retry",
			err,
		)
	}

	if code == pgerrcode.SerializationFailure {
		return domainerror.NewConflictError(
			"concurrent modification detected, please retry",
			err,
		)
	}

	if code == pgerrcode.TooManyConnections {
		return domainerror.NewUnavailableError(
			"service temporarily unavailable, please try again",
			err,
		)
	}

	if code == pgerrcode.QueryCanceled {
		return domainerror.NewInternalError(
			"operation timed out",
			err,
		)
	}

	return domainerror.NewInternalError(
		"database error occurred",
		err,
	)
}
