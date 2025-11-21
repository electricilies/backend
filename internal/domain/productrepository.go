package domain

import (
	"github.com/jackc/pgx/v5"
)

type ProductRepository interface {
	List(
		tx pgx.Tx,
		ids *[]int,
		search *string,
		min_price *int64,
		max_price *int64,
		rating *float64,
		category_ids *[]int,
		deleted string,
		sort_rating *string,
		sort_price *string,
		limit int,
		offset int,
	) (*[]Product, error)

	Count(
		tx pgx.Tx,
		ids *[]int,
		min_price *int64,
		max_price *int64,
		rating *float64,
		category_ids *[]int,
		deleted string,
		sort_rating *string,
		sort_price *string,
		limit int,
		offset int,
	) (*int, error)

	Create(
		tx pgx.Tx,
		name string,
		description string,
		attributeValueIDs []int,
		categoryID int,
	) (*Product, error)

	CreateOptions(
		tx pgx.Tx,
		productID int,
		options []struct {
			name   string
			values []string
		},
	) (*[]Option, error)

	CreateOptionValues(
		tx pgx.Tx,
		optionID int,
		values []struct {
			value string
		},
	) (*[]OptionValue, error)

	CreateVariants(
		tx pgx.Tx,
		productID int,
		variants []struct {
			sku            string
			price          int64
			quantity       int
			optionValueIDs []int
		},
	) (*[]ProductVariant, error)

	CreateImages(
		tx pgx.Tx,
		productID int,
		images []struct {
			url   string
			order int
		},
	) (*[]ProductImage, error)

	CreateVariantImages(
		tx pgx.Tx,
		variantID int,
		images []struct {
			url   string
			order int
		},
	) (*[]ProductImage, error)

	Get(
		tx pgx.Tx,
		productID int,
	) (*Product, error)

	Update(
		tx pgx.Tx,
		productID int,
		name *string,
		description *string,
		categoryID *int,
	) (*Product, error)

	Delete(
		tx pgx.Tx,
		productID int,
	) error

	UpdateVariant(
		tx pgx.Tx,
		variantID int,
		price *int64,
		quantity *int,
	) (*ProductVariant, error)

	UpdateOptions(
		tx pgx.Tx,
		options []struct {
			id   int
			name string
		},
	) error

	UpdateOptionValues(
		tx pgx.Tx,
		optionValues []struct {
			id    int
			value string
		},
	) error
}
