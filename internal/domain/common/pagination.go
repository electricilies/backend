package common

type PaginationMetadata struct {
	TotalRecords int
	CurrentPage  int
	ItemsPerPage int
	PageItems    int
}

type Pagination[T any] struct {
	Metadata PaginationMetadata
	Items    []T
}
