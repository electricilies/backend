package helper

func AppendNotExistsToSlice[T comparable](slice []T, elements ...T) []T {
	existing := make(map[T]struct{}, len(slice))
	for _, item := range slice {
		existing[item] = struct{}{}
	}
	for _, elem := range elements {
		if _, found := existing[elem]; !found {
			slice = append(slice, elem)
			existing[elem] = struct{}{}
		}
	}
	return slice
}

// https://go.dev/wiki/SliceTricks#filtering-without-allocating
func RemoveElementsFromSlice[T comparable](origin []T, toRemove []T) []T {
	set := make(map[T]struct{}, len(toRemove))
	for _, elem := range toRemove {
		set[elem] = struct{}{}
	}
	new := origin[:0]
	for _, item := range origin {
		if _, found := set[item]; !found {
			new = append(new, item)
		}
	}
	clear(origin[len(new):])
	return new
}

func FilterSlice[T any](origin []T, predicate func(T) bool) []T {
	new := origin[:0]
	for _, item := range origin {
		if predicate(item) {
			new = append(new, item)
		}
	}
	clear(origin[len(new):])
	return new
}

func FilterSliceFromSlice[T any, U comparable](origin []T, toRemove []U, mapper func(T) U) []T {
	set := make(map[U]struct{}, len(toRemove))
	for _, elem := range toRemove {
		set[elem] = struct{}{}
	}
	newSlice := origin[:0]
	for _, item := range origin {
		mappedValue := mapper(item)
		if _, found := set[mappedValue]; !found {
			newSlice = append(newSlice, item)
		}
	}
	clear(origin[len(newSlice):])
	return newSlice
}

// CheckAllExist checks if all elements in toCheck exist in the origin slice
// Returns true if all exist, false otherwise
func CheckAllExist[T any, U comparable](origin []T, toCheck []U, mapper func(T) U) bool {
	existingSet := make(map[U]struct{}, len(origin))
	for _, item := range origin {
		existingSet[mapper(item)] = struct{}{}
	}
	for _, elem := range toCheck {
		if _, found := existingSet[elem]; !found {
			return false
		}
	}
	return true
}

// CheckAnyExist checks if any elements in toCheck exist in the origin slice
// Returns true if at least one exists, false if none exist
func CheckAnyExist[T any, U comparable](origin []T, toCheck []U, mapper func(T) U) bool {
	existingSet := make(map[U]struct{}, len(origin))
	for _, item := range origin {
		existingSet[mapper(item)] = struct{}{}
	}
	for _, elem := range toCheck {
		if _, found := existingSet[elem]; found {
			return true
		}
	}
	return false
}

// FindNonExistent returns a slice of elements from toCheck that don't exist in origin
func FindNonExistent[T any, U comparable](origin []T, toCheck []U, mapper func(T) U) []U {
	existingSet := make(map[U]struct{}, len(origin))
	for _, item := range origin {
		existingSet[mapper(item)] = struct{}{}
	}
	var nonExistent []U
	for _, elem := range toCheck {
		if _, found := existingSet[elem]; !found {
			nonExistent = append(nonExistent, elem)
		}
	}
	return nonExistent
}

// FindExisting returns a slice of elements from toCheck that exist in origin
func FindExisting[T any, U comparable](origin []T, toCheck []U, mapper func(T) U) []U {
	existingSet := make(map[U]struct{}, len(origin))
	for _, item := range origin {
		existingSet[mapper(item)] = struct{}{}
	}
	var existing []U
	for _, elem := range toCheck {
		if _, found := existingSet[elem]; found {
			existing = append(existing, elem)
		}
	}
	return existing
}
