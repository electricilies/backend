package service

type PaginationParam struct {
	Page  *int
	Limit *int
}

type PaginationMeta struct {
	TotalItems   int `json:"totalItems" binding:"required"`
	CurrentPage  int `json:"currentPage" binding:"required"`
	ItemsPerPage int `json:"itemsPerPage" binding:"required"`
	PageItems    int `json:"pageItems"`
}

type Pagination[T interface{}] struct {
	Data []T            `json:"data" binding:"required"`
	Meta PaginationMeta `json:"meta" binding:"required"`
}
