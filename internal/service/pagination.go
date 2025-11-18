package service

type PaginationParam struct {
	Page  int `json:"page" example:"1"`
	Limit int `json:"limit" example:"10"`
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
