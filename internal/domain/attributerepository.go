package domain

import (
	"github.com/jackc/pgx/v5"
)

type AttributeRepository interface {
	Count(
		tx pgx.Tx,
		ids *[]int,
		deleted string,
		limit int,
		offset int,
	) (*int, error)

	List(
		tx pgx.Tx,
		ids *[]int,
		search *string,
		deleted string,
		limit int,
		offset int,
	) (*[]Attribute, error)

	Create(
		tx pgx.Tx,
		code string,
		name string,
	) (*Attribute, error)

	Get(
		tx pgx.Tx,
		id int,
	) (*Attribute, error)

	Update(
		tx pgx.Tx,
		id int,
		code *string,
		name *string,
	) (*Attribute, error)

	CreateValues(
		tx pgx.Tx,
		attributeID int,
		values []string,
	) (*[]AttributeValue, error)

	UpdateValue(
		tx pgx.Tx,
		attributeValueID int,
		value *string,
	) (*AttributeValue, error)

	Delete(
		tx pgx.Tx,
		id int,
	) error

	DeleteValue(
		tx pgx.Tx,
		attributeValueID int,
	) error
}
