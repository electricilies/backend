package application

type ListCategoryParam struct {
	PaginationParam
	Search *string `json:"search" binding:"omitnil"`
}
