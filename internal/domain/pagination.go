package domain

type Pagination struct {
	TotalItems   int `json:"totalItems" binding:"required"`
	CurrentPage  int `json:"currentPage" binding:"required"`
	ItemsPerPage int `json:"itemsPerPage" binding:"required"`
	PageItems    int `json:"pageItems"`
}

type DataPagination struct {
	Data interface{} `json:"data" binding:"required"`
	Meta *Pagination `json:"pagination" binding:"required"`
}
