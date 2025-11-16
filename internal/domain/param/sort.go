package param

type Sort string

const (
	SortAscending  Sort = "asc"
	SortDescending Sort = "desc"
)

type (
	SortRating Sort
	SortPrice  Sort
)
