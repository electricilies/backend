package param

type Pagination struct {
	Limit  int
	Offset int
}

type PaginationMetadata struct {
	TotalRecords int
	CurrentPage  int
	ItemsPerPage int
	PageItems    int
}
