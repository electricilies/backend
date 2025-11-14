package helper

// https://github.com/kubernetes/utils/blob/master/ptr/ptr.go

func To[T any](v T) *T {
	return &v
}

func Deref[T any](ptr *T, def T) T {
	if ptr != nil {
		return *ptr
	}
	return def
}
