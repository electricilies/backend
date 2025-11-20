package repository

import (
	"backend/internal/domain"

	"github.com/jackc/pgx/v5"
)

type Product interface {
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
	) (*[]domain.Product, error)

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
	) (*domain.Product, error)

	CreateOptions(
		tx pgx.Tx,
		productID int,
		options []struct {
			name   string
			values []string
		},
	) (*[]domain.Option, error)

	CreateOptionValues(
		tx pgx.Tx,
		optionID int,
		values []struct {
			value string
		},
	) (*[]domain.OptionValue, error)

	CreateVariants(
		tx pgx.Tx,
		productID int,
		variants []struct {
			sku            string
			price          int64
			quantity       int
			optionValueIDs []int
		},
	) (*[]domain.ProductVariant, error)

	CreateImages(
		tx pgx.Tx,
		productID int,
		images []struct {
			url   string
			order int
		},
	) (*[]domain.ProductImage, error)

	CreateVariantImages(
		tx pgx.Tx,
		variantID int,
		images []struct {
			url   string
			order int
		},
	) (*[]domain.ProductImage, error)

	Get(
		tx pgx.Tx,
		productID int,
	) (*domain.Product, error)

	Update(
		tx pgx.Tx,
		productID int,
		name *string,
		description *string,
		categoryID *int,
	) (*domain.Product, error)

	Delete(
		tx pgx.Tx,
		productID int,
	) error

	UpdateVariant(
		tx pgx.Tx,
		variantID int,
		price *int64,
		quantity *int,
	) (*domain.ProductVariant, error)

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
