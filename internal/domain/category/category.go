package category

import (
	"time"
)

type Category struct {
	id        int
	name      string
	createdAt time.Time
	updatedAt time.Time
	deletedAt time.Time
}

func NewCategory(name string) *Category {
	return &Category{
		name:      name,
		createdAt: time.Now(),
		updatedAt: time.Now(),
	}
}

func (c *Category) GetID() int {
	return c.id
}

func (c *Category) GetName() string {
	return c.name
}

func (c *Category) GetCreatedAt() time.Time {
	return c.createdAt
}

func (c *Category) GetUpdatedAt() time.Time {
	return c.updatedAt
}

func (c *Category) GetDeletedAt() time.Time {
	return c.deletedAt
}

func (c *Category) SetName(name string) {
	c.name = name
	c.updatedAt = time.Now()
}

func (c *Category) SetDeletedAt(deletedAt time.Time) {
	c.deletedAt = deletedAt
	c.updatedAt = time.Now()
}
