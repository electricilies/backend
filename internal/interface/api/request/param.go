package request

import "backend/internal/domain/param"

const ErrorInvalidDeletedParam = errors.New("invalid deleted param")
const ErrorInvalidSortParam = errors.New("invalid sort param")
const ErrorInvalidSortRatingParam = errors.New("invalid sort rating param")

func PaginationParamsToDomain(limit, offset int) param.Pagination {
	return param.Pagination{
		Limit:  limit,
		Offset: offset,
	}
}

func DeletedParamToDomain(deleted string) (*param.Deleted, error) {
	switch deleted {
	case "all":
		all := param.All
		return &all, nil
	case "only":
		only := param.Only
		return &only, nil
	case "exclude":
		exclude := param.Exclude
		return &exclude, nil
	default:
		return nil,  ErrorInvalidDeletedParam
	}
}

func sortParamToDomain(sort string) (*param.Sort, error) {
	switch sort {
	case "asc":
		asc := param.Ascending
		return &asc, nil
	case "desc":
		desc := param.Descending
		return &desc, nil
	default:
		return nil, ErrorInvalidSortParam
	}
}

func SortPriceParamToDomain(sort string) (*param.SortPrice, error) {
	sortPrice, err := sortParamToDomain(sort)
	if err != nil {
		return nil, ErrorInvalidSortParam
	}
	return sortPrice, nil
}

func SortRatingParamToDomain(sort string) (*param.SortRating, error) {
	sortRating, err := sortParamToDomain(sort)
	if err != nil {
		return nil, ErrorInvalidSortRatingParam
	}
	return sortRating, nil
}
