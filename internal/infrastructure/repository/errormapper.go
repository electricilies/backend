package repository

import (
	"context"
	"errors"
	"net"
	"net/http"

	"backend/internal/domain"

	"github.com/Nerzal/gocloak/v13"
	awshttp "github.com/aws/aws-sdk-go-v2/aws/transport/http"
	"github.com/aws/smithy-go"
	"github.com/hashicorp/go-multierror"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

func ToDomainErrorFromPostgres(err error) error {
	if err == nil {
		return nil
	}

	if errors.Is(err, pgx.ErrNoRows) {
		return WrapMutilError(domain.ErrNotFound, "record not found", err)
	}

	var pgErr *pgconn.PgError
	if !errors.As(err, &pgErr) {
		return WrapMutilError(domain.ErrInternalDB, "unexpected postgres error", err)
	}

	switch pgErr.Code {

	case pgerrcode.NotNullViolation:
		field := pgErr.ColumnName
		if field == "" {
			field = "field"
		}
		return WrapMutilError(domain.ErrInvalid, field+" is required", err)

	case pgerrcode.CheckViolation:
		return WrapMutilError(domain.ErrInvalid, "check constraint failed: "+pgErr.ConstraintName, err)

	case pgerrcode.InvalidTextRepresentation:
		return WrapMutilError(domain.ErrInvalidFormat, "invalid text representation", err)

	case pgerrcode.StringDataRightTruncationDataException:
		field := pgErr.ColumnName
		if field == "" {
			field = "field"
		}
		return WrapMutilError(domain.ErrInvalid, field+" too long", err)

	// --- CONFLICT ---
	case pgerrcode.UniqueViolation:
		return WrapMutilError(domain.ErrExists, pgErr.ConstraintName+" already exists", err)

	case pgerrcode.ForeignKeyViolation:
		return WrapMutilError(domain.ErrInvalid, "referenced row not found", err)

	case pgerrcode.DeadlockDetected:
		return WrapMutilError(domain.ErrConflict, "deadlock detected", err)

	case pgerrcode.SerializationFailure:
		return WrapMutilError(domain.ErrConflict, "serialization failed", err)

	case pgerrcode.TooManyConnections:
		return WrapMutilError(domain.ErrUnavailable, "too many connections", err)

	case pgerrcode.AdminShutdown,
		pgerrcode.CrashShutdown,
		pgerrcode.CannotConnectNow:
		return WrapMutilError(domain.ErrUnavailable, "database not available", err)

	case pgerrcode.ConnectionFailure:
		return WrapMutilError(domain.ErrConnection, "database connection failed", err)

	default:
		return WrapMutilError(domain.ErrInternalDB, "unhandled postgres error", err)
	}
}

func ToDomainErrorFromS3(err error) error {
	if err == nil {
		return nil
	}

	var apiErr smithy.APIError
	if errors.As(err, &apiErr) {
		switch apiErr.ErrorCode() {

		case "AccessDenied":
			return WrapMutilError(domain.ErrForbidden, "s3 access denied", err)

		case "NoSuchBucket", "NoSuchKey":
			return WrapMutilError(domain.ErrNotFound, "s3 resource not found", err)

		case "InvalidBucketName", "InvalidObjectState":
			return WrapMutilError(domain.ErrInvalidFormat, "invalid s3 request", err)

		case "BucketAlreadyExists", "BucketAlreadyOwnedByYou":
			return WrapMutilError(domain.ErrExists, "bucket conflict", err)

		case "ServiceUnavailable", "SlowDown":
			return WrapMutilError(domain.ErrUnavailable, "s3 service unavailable", err)

		case "RequestTimeout":
			return WrapMutilError(domain.ErrTimeout, "s3 request timeout", err)

		default:
			return WrapMutilError(domain.ErrServiceError, "unknown s3 error", err)
		}
	}

	var httpErr *awshttp.ResponseError

	if errors.As(err, &httpErr) {
		switch httpErr.HTTPStatusCode() {
		case 400:
			return WrapMutilError(domain.ErrInvalid, "s3 bad request", err)
		case 403:
			return WrapMutilError(domain.ErrForbidden, "s3 forbidden", err)
		case 404:
			return WrapMutilError(domain.ErrNotFound, "s3 not found", err)
		case 409:
			return WrapMutilError(domain.ErrConflict, "s3 conflict", err)
		case 408:
			return WrapMutilError(domain.ErrTimeout, "s3 timeout", err)
		case 503:
			return WrapMutilError(domain.ErrUnavailable, "s3 unavailable", err)
		default:
			return WrapMutilError(domain.ErrServiceError, "s3 http error", err)
		}
	}
	var netErr net.Error
	if errors.As(err, &netErr) {
		if netErr.Timeout() {
			return WrapMutilError(domain.ErrTimeout, "s3 network timeout", err)
		}
		return WrapMutilError(domain.ErrConnection, "s3 network failure", err)
	}

	return WrapMutilError(domain.ErrUnknown, "unexpected s3 error", err)
}

func ToDomainErrorFromGoCloak(err error) error {
	if err == nil {
		return nil
	}

	var apiErr gocloak.APIError
	if errors.As(err, &apiErr) {
		switch apiErr.Code {

		case http.StatusBadRequest:
			return WrapMutilError(domain.ErrInvalid, apiErr.Message, err)

		case http.StatusForbidden:
			return WrapMutilError(domain.ErrForbidden, "access denied", err)

		case http.StatusNotFound:
			return WrapMutilError(domain.ErrNotFound, "resource not found", err)

		case http.StatusConflict:
			return WrapMutilError(domain.ErrExists, "resource conflict", err)

		case http.StatusServiceUnavailable, http.StatusGatewayTimeout:
			return WrapMutilError(domain.ErrUnavailable, "keycloak unavailable", err)

		default:
			return WrapMutilError(domain.ErrServiceError, "unknown keycloak error", err)
		}
	}

	if errors.Is(err, context.DeadlineExceeded) {
		return WrapMutilError(domain.ErrTimeout, "keycloak timeout", err)
	}

	var netErr net.Error
	if errors.As(err, &netErr) {
		if netErr.Timeout() {
			return WrapMutilError(domain.ErrTimeout, "keycloak network timeout", err)
		}
		return WrapMutilError(domain.ErrConnection, "keycloak network error", err)
	}

	return WrapMutilError(domain.ErrUnknown, "unexpected keycloak error", err)
}

func WrapMutilError(domainErr error, msg string, original error) error {
	return multierror.Append(
		nil,
		domainErr,
		errors.New(msg),
		original,
	)
}
