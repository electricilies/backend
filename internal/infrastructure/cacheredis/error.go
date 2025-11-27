package cacheredis

import (
	"errors"

	"backend/internal/domain"

	"github.com/hashicorp/go-multierror"
	"github.com/redis/go-redis/v9"
)

var ErrClientNil = errors.New("redis client is nil")

func toDomainError(
	err error,
) error {
	if err == nil {
		return nil
	}
	switch err {
	case ErrClientNil:
		return multierror.Append(domain.ErrServiceError, err)
	case redis.Nil:
		return multierror.Append(domain.ErrServiceError, errors.New("cache miss"), err)
	default:
		return multierror.Append(domain.ErrServiceError, err)
	}
}
