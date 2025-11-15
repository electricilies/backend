package response

import "backend/internal/domain/param"

type Pagination struct {
	TotalItems   int `json:"totalItems" binding:"required"`
	CurrentPage  int `json:"currentPage" binding:"required"`
	ItemsPerPage int `json:"itemsPerPage" binding:"required"`
	PageItems    int `json:"pageItems"`
}

func PaginationFromDomain(p *param.PaginationMetadata) *Pagination {
	return &Pagination{
		TotalItems:   *p.TotalRecords,
		CurrentPage:  *p.CurrentPage,
		ItemsPerPage: *p.ItemsPerPage,
		PageItems:    *p.PageItems,
	}
}

type DataPagination struct {
	Data interface{} `json:"data" binding:"required"`
	Meta *Pagination `json:"pagination" binding:"required"`
}

func DataPaginationFromDomain(data interface{}, p *param.PaginationMetadata) *DataPagination {
	return &DataPagination{
		Data: data,
		Meta: PaginationFromDomain(p),
	}
}
