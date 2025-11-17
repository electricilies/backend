package mapper

import (
	"math/big"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

func ToPostgresTextType(v *string) pgtype.Text {
	if v == nil {
		return pgtype.Text{
			Valid: false,
		}
	}
	return pgtype.Text{
		String: *v,
		Valid:  true,
	}
}

func ToPostgresTypeInt(v *int) pgtype.Int4 {
	if v == nil {
		return pgtype.Int4{
			Valid: false,
		}
	}
	return pgtype.Int4{
		Int32: int32(*v),
		Valid: true,
	}
}

func ToPostgresTypeFloat(v *float32) pgtype.Float4 {
	if v == nil {
		return pgtype.Float4{
			Valid: false,
		}
	}
	return pgtype.Float4{
		Float32: *v,
		Valid:   true,
	}
}

func ToPostgresTypeNumeric(v *int64) pgtype.Numeric {
	if v == nil {
		return pgtype.Numeric{
			Valid: false,
		}
	}
	return pgtype.Numeric{
		Int:   big.NewInt(*v),
		Valid: true,
	}
}

func ToInt32Slice(v *[]int) []int32 {
	if v == nil {
		return nil
	}
	result := make([]int32, len(*v))
	for i, v := range *v {
		result[i] = int32(v)
	}
	return result
}

func ToPostgresTimeStampType(v *time.Time) pgtype.Timestamp {
	if v == nil {
		return pgtype.Timestamp{
			Valid: false,
		}
	}
	return pgtype.Timestamp{
		Time:  *v,
		Valid: true,
	}
}
