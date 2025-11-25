package repositorypostgres

func fromPgValidToPtr[T any](value T, valid bool) *T {
	if !valid {
		return nil
	}
	return &value
}

// func fromPgValidToNonPtr[T any](value T, valid bool, defaultValue T) T {
// 	if !valid {
// 		return defaultValue
// 	}
// 	return value
// }
