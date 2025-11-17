package repository

import "reflect"

func mapPgtypeToDomainPtr[T interface{}](value T, valid bool) *T {
	if !valid {
		return nil
	}
	return &value
}

func domainPtrToPgtype[T interface{}](value *T) (T, bool) {
	if value == nil {
		var zero T
		return zero, false
	}
	return *value, true
}

func domainPtrSliceToPgtype[T, U interface{}](value *[]T, mapper ...func(T) U) ([]U, bool) {
	if len(mapper) > 1 {
		panic("only one mapper function is allowed")
	}
	if value == nil {
		var zero []U
		return zero, false
	}
	var t T
	var u U
	if reflect.TypeOf(t) == reflect.TypeOf(u) && len(mapper) == 0 {
		return any(*value).([]U), true
	}
	mapped := make([]U, len(*value))
	for i, v := range *value {
		mapped[i] = mapper[0](v)
	}
	return mapped, true
}
