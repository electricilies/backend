package param

type Pagination struct {
	Limit  int
	Offset int
}

type Metadata struct {
	TotalRecords int
	CurrentPage  int
	ItemsPerPage int
}
