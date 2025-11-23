package slice

func SlicePointerToSlice[T any](input []*T) []T {
	result := make([]T, 0, len(input))
	for _, v := range input {
		if v != nil {
			result = append(result, *v)
		}
	}
	return result
}

func SlicePtrToPtrSlice[T any](input []*T) *[]T {
	result := SlicePointerToSlice(input)
	return &result
}
