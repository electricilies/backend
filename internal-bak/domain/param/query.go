package param

type Deleted string

const (
	Exclude Deleted = "exclude"
	Only    Deleted = "only"
	All     Deleted = "all"
)

type Sort string

const (
	Ascending  Sort = "asc"
	Descending Sort = "desc"
)

type (
	SortRating Sort
	SortPrice  Sort
)
