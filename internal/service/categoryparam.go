package service

type ListCategoryParam struct {
	PaginationParam
	Search *string `json:"search"`
}

type GetCategoryParam struct {
	CategoryID int
}

type CreateCategoryData struct {
	Name string `json:"name" binding:"required"`
}

type CreateCategoryParam struct {
	Data CreateCategoryData
}

type UpdateCategoryData struct {
	Name *string `json:"name" binding:"omitnil"`
}

type UpdateCategoryParam struct {
	CategoryID int
	Data       UpdateCategoryData
}
