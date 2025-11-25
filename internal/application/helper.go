package application

import "backend/internal/delivery/http"

func newPaginationResponseDto[T interface{}](
	data []T,
	totalItems, currentPage, itemsPerPage int,
) *http.PaginationResponseDto[T] {
	return &http.PaginationResponseDto[T]{
		Data: data,
		Meta: http.PaginationMetaResponseDto{
			TotalItems:   totalItems,
			CurrentPage:  currentPage,
			ItemsPerPage: itemsPerPage,
			PageItems:    len(data),
		},
	}
}
