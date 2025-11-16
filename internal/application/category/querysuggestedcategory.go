package category

import "backend/internal/domain/category"

type QuerySuggestedCategory interface {
	ListSuggestedCategories(category.Category) (*[]string, error)
}
