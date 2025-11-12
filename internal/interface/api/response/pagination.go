package response

import "backend/internal/domain/params"

type Pagination struct {
	TotalItems   int `json:"totalItems" binding:"required"`
	CurrentPage  int `json:"currentPage" binding:"required"`
	ItemsPerPage int `json:"itemsPerPage" binding:"required"`
}

func PaginationFromDomain(p *params.Metadata) *Pagination {
	return &Pagination{
		TotalItems:   p.TotalRecords,
		CurrentPage:  p.CurrentPage,
		ItemsPerPage: p.ItemsPerPage,
	}
}

type DataPagination struct {
	Data interface{} `json:"data" binding:"required"`
	Meta *Pagination `json:"pagination" binding:"required"`
}

func DataPaginationFromDomain(data interface{}, p *params.Metadata) *DataPagination {
	return &DataPagination{
		Data: data,
		Meta: PaginationFromDomain(p),
	}
}
