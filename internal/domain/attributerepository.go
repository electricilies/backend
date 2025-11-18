package domain

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type AttributeRepository interface {
	Get(
		tx pgx.Tx,
		ID int,
	) (*Attribute, error)

	List(
		tx pgx.Tx,
		IDs *[]int,
		Search *string,
		Deleted string,
		Limit int,
		Offset int,
	) (*[]Attribute, error)

	Create(
		tx pgx.Tx,
		Code string,
		Name string,
	) (*Attribute, error)

	Update(
		tx pgx.Tx,
		ID int,
		Code *string,
		Name *string,
	) (*Attribute, error)

	CreateValues(
		tx pgx.Tx,
		attributeID int,
		values []string,
	) (*[]AttributeValue, error)

	UpdateValues(
		tx pgx.Tx,
		attributeValueID int,
		value *string,
	) (*[]AttributeValue, error)
}
