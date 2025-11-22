package repository

func fromPgValidToPtr[T any](value T, valid bool) *T {
	if !valid {
		return nil
	}
	return &value
}
