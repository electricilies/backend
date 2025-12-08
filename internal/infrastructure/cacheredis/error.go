package cacheredis

import (
	"backend/internal/domain"

	"github.com/hashicorp/go-multierror"
)

func toDomainError(
	err error,
) error {
	return multierror.Append(domain.ErrServiceError, err)
}
