package application

type ListProductParam struct {
	PaginationParam
	CategoryIDs *[]int  `binding:"omitempty"`
	MinPrice    *int64  `binding:"omitempty"`
	MaxPrice    *int64  `binding:"omitempty"`
	SortPrice   *string `binding:"omitempty,oneof=asc desc"`
	SortRating  *string `binding:"omitempty,oneof=asc desc"`
	Search      *string `binding:"omitempty"`
	Deleted     *string `binding:"omitempty,oneof=exclude only all"`
}
