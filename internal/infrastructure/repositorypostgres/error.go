package repositorypostgres

import (
	"context"
	"errors"
	"net"
	"net/http"

	"backend/internal/domain"

	"github.com/Nerzal/gocloak/v13"
	"github.com/hashicorp/go-multierror"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

func toDomainError(err error) error {
	if err == nil {
		return nil
	}

	if errors.Is(err, pgx.ErrNoRows) {
		return multierror.Append(domain.ErrNotFound, errors.New("record not found"), err)
	}

	var pgErr *pgconn.PgError
	if !errors.As(err, &pgErr) {
		return multierror.Append(domain.ErrServiceError, errors.New("unexpected postgres error"), err)
	}

	switch pgErr.Code {

	case pgerrcode.NotNullViolation:
		field := pgErr.ColumnName
		if field == "" {
			field = "field"
		}
		return multierror.Append(domain.ErrInvalid, errors.New(field+" is required"), err)

	case pgerrcode.CheckViolation:
		return multierror.Append(domain.ErrInvalid, errors.New("check constraint failed: "+pgErr.ConstraintName), err)

	case pgerrcode.InvalidTextRepresentation:
		return multierror.Append(domain.ErrInvalid, errors.New("invalid text representation"), err)

	case pgerrcode.StringDataRightTruncationDataException:
		field := pgErr.ColumnName
		if field == "" {
			field = "field"
		}
		return multierror.Append(domain.ErrInvalid, errors.New(field+" too long"), err)

	// --- CONFLICT ---
	case pgerrcode.UniqueViolation:
		return multierror.Append(domain.ErrExists, errors.New(pgErr.ConstraintName+" already exists"), err)

	case pgerrcode.ForeignKeyViolation:
		return multierror.Append(domain.ErrInvalid, errors.New("referenced row not found"), err)

	case pgerrcode.DeadlockDetected:
		return multierror.Append(domain.ErrConflict, errors.New("deadlock detected"), err)

	case pgerrcode.SerializationFailure:
		return multierror.Append(domain.ErrConflict, errors.New("serialization failed"), err)

	case
		pgerrcode.ConnectionFailure,
		pgerrcode.TooManyConnections, pgerrcode.AdminShutdown,
		pgerrcode.CrashShutdown,
		pgerrcode.CannotConnectNow:
		return multierror.Append(domain.ErrUnavailable, errors.New("database connection failed"), err)

	default:
		return multierror.Append(domain.ErrServiceError, errors.New("unhandled postgres error"), err)
	}
}

// FIXME: wtf keycloak here?
func ToDomainErrorFromGoCloak(err error) error {
	if err == nil {
		return nil
	}

	var apiErr gocloak.APIError
	if errors.As(err, &apiErr) {
		switch apiErr.Code {

		case http.StatusBadRequest:
			return multierror.Append(domain.ErrInvalid, errors.New(apiErr.Message), err)

		case http.StatusForbidden:
			return multierror.Append(domain.ErrForbidden, errors.New("access denied"), err)

		case http.StatusNotFound:
			return multierror.Append(domain.ErrNotFound, errors.New("resource not found"), err)

		case http.StatusConflict:
			return multierror.Append(domain.ErrExists, errors.New("resource conflict"), err)

		case http.StatusServiceUnavailable, http.StatusGatewayTimeout:
			return multierror.Append(domain.ErrUnavailable, errors.New("keycloak unavailable"), err)

		default:
			return multierror.Append(domain.ErrServiceError, errors.New("unknown keycloak error"), err)
		}
	}

	if errors.Is(err, context.DeadlineExceeded) {
		return multierror.Append(domain.ErrTimeout, errors.New("keycloak timeout"), err)
	}

	var netErr net.Error
	if errors.As(err, &netErr) {
		if netErr.Timeout() {
			return multierror.Append(domain.ErrTimeout, errors.New("keycloak network timeout"), err)
		}
		return multierror.Append(domain.ErrUnavailable, errors.New("keycloak network error"), err)
	}

	return multierror.Append(domain.ErrUnknown, errors.New("unexpected keycloak error"), err)
}
