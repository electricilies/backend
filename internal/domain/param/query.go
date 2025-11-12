package param

type Deleted string

const (
	Exclude Deleted = "exclude"
	Only    Deleted = "only"
	All     Deleted = "all"
)
