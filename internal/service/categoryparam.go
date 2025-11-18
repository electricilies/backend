package service

type CreateCategoryParam struct {
	Data CreateCategoryData `binding:"required"`
}

type CreateCategoryData struct {
	Name string `json:"name" binding:"required"`
}

type GetCategoryParam struct {
	CategoryID int `binding:"required"`
}

type UpdateCategoryParam struct {
	CategoryID int `binding:"required"`
	Data       UpdateCategoryData `binding:"required"`
}

type UpdateCategoryData struct {
	Name *string `json:"name" binding:"omitnil"`
}

type ListCategoryParam struct {
	PaginationParam
	Search *string `json:"search" binding:"omitnil"`
}
