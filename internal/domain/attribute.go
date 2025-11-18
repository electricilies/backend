package domain

import "time"

type Attribute struct {
	ID        int               `json:"id" binding:"required" example:"123"`
	Code      string            `json:"code" binding:"required" example:"color"`
	Name      string            `json:"name" binding:"required" example:"Color"`
	Values    *[]AttributeValue `json:"values,omitnil"`
	DeletedAt time.Time         `json:"deletedAt,omitempty" binding:"required"`
}

type AttributeValue struct {
	ID    int    `json:"id" binding:"required" example:"1"`
	Value string `json:"value" binding:"required" example:"Red"`
}
