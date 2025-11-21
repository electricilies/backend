package application

func newPagintion[T any](data []T, totalItems, currentPage, itemsPerPage int) *Pagination[T] {
	return &Pagination[T]{
		Data: data,
		Meta: PaginationMeta{
			TotalItems:   totalItems,
			CurrentPage:  currentPage,
			ItemsPerPage: itemsPerPage,
			PageItems:    len(data),
		},
	}
}
