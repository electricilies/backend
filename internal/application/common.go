package application

type DeleteImageURL struct {
	URL string `json:"url" binding:"required,url"`
}

type UploadImageURL struct {
	URL string `json:"url" binding:"required,url"`
	Key string `json:"key" binding:"required"`
}

type PaginationParam struct {
	Page  *int `binding:"omitnil,min=1,max=50"`
	Limit *int `binding:"omitnil,min=1,max=100"`
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

type DeletedParam string

const (
	DeletedExcludeParam DeletedParam = "exclude"
	DeletedOnlyParam    DeletedParam = "only"
	DeletedAllParam     DeletedParam = "all"
)
