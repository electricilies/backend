package response

import "backend/internal/domain/pagination"

type Pagination struct {
	TotalItems   int `json:"totalItems" binding:"required"`
	CurrentPage  int `json:"currentPage" binding:"required"`
	ItemsPerPage int `json:"itemsPerPage" binding:"required"`
}

func PaginationFromDomain(p *pagination.Metadata) *Pagination {
	return &Pagination{
		TotalItems:   p.TotalRecords,
		CurrentPage:  p.CurrentPage,
		ItemsPerPage: p.ItemsPerPage,
	}
}
