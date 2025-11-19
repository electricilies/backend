package repository

import (
	"backend/internal/domain"

	"github.com/jackc/pgx/v5"
)

type Attribute interface {
	Get(
		tx pgx.Tx,
		id int,
	) (*domain.Attribute, error)

	List(
		tx pgx.Tx,
		ids *[]int,
		search *string,
		deleted string,
		limit int,
		offset int,
	) (*[]domain.Attribute, error)

	Create(
		tx pgx.Tx,
		code string,
		name string,
	) (*domain.Attribute, error)

	Update(
		tx pgx.Tx,
		id int,
		code *string,
		name *string,
	) (*domain.Attribute, error)

	CreateValues(
		tx pgx.Tx,
		attributeID int,
		values []string,
	) (*[]domain.AttributeValue, error)

	UpdateValues(
		tx pgx.Tx,
		attributeValueID int,
		value *string,
	) (*[]domain.AttributeValue, error)

	Delete(
		tx pgx.Tx,
		id int,
	) error

	DeleteValue(
		tx pgx.Tx,
		attributeValueID int,
	) error
}
