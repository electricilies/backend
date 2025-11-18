package service

type CreateCategoryParam struct {
	Data CreateCategoryData
}

type CreateCategoryData struct {
	Name string `json:"name" binding:"required"`
}

type GetCategoryParam struct {
	CategoryID int
}

type UpdateCategoryParam struct {
	CategoryID int
	Data       UpdateCategoryData
}

type UpdateCategoryData struct {
	Name *string `json:"name" binding:"omitnil"`
}

type ListCategoryParam struct {
	PaginationParam
	Search *string `json:"search"`
}
