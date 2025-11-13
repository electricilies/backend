package request

import "backend/internal/domain/param"

func PaginationParamsToDomain(limit int, offset int) *param.Pagination {
	return &param.Pagination{
		Limit:  &limit,
		Offset: &offset,
	}
}

func DeletedParamToDomain(deleted string) *param.Deleted {
	switch deleted {
	case "all":
		all := param.All
		return &all
	case "only":
		only := param.Only
		return &only
	case "exclude":
		exclude := param.Exclude
		return &exclude
	default:
		return nil
	}
}

func SortParamToDomain(sort string) *param.Sort {
	switch sort {
	case "asc":
		asc := param.Ascending
		return &asc
	case "desc":
		desc := param.Descending
		return &desc
	default:
		return nil
	}
}

func SortPriceParamToDomain(sort string) *param.SortPrice {
	s := SortParamToDomain(sort)
	if s != nil {
		return (*param.SortPrice)(s)
	}
	// TODO: if more sort option handle later
	switch sort {
	default:
		return nil
	}
}

func SortRatingParamToDomain(sort string) *param.SortRating {
	s := SortParamToDomain(sort)
	if s != nil {
		return (*param.SortRating)(s)
	}
	// TODO: if more sort option handle later
	switch sort {
	default:
		return nil
	}
}
