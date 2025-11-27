package repositorypostgres

import (
	"math/big"

	"github.com/jackc/pgx/v5/pgtype"
)

func fromPgValidToPtr[T any](value T, valid bool) *T {
	if !valid {
		return nil
	}
	return &value
}

func fromPgValidToNonPtr[T any](value T, valid bool, defaultValue T) T {
	if !valid {
		return defaultValue
	}
	return value
}

func numericToInt64(n pgtype.Numeric) int64 {
	if !n.Valid {
		return 0
	}
	var result big.Int
	n.Int.Mul(n.Int, big.NewInt(0).Exp(big.NewInt(10), big.NewInt(int64(n.Exp)), nil))
	result.Set(n.Int)
	return result.Int64()
}

func int64ToNumeric(value int64) pgtype.Numeric {
	return pgtype.Numeric{
		Int:   big.NewInt(value),
		Exp:   0,
		Valid: true,
	}
}
