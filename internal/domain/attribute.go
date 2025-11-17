package domain

import "time"

type Attribute struct {
	ID              int              `json:"id" binding:"required" example:"123"`
	Code            string           `json:"code" binding:"required" example:"color"`
	Name            string           `json:"name" binding:"required" example:"Color"`
	AttributeValues []AttributeValue `json:"attributeValues" binding:"required"`
	DeletedAt       *time.Time       `json:"deletedAt" binding:"required"`
}

type AttributeValue struct {
	ID    int    `json:"id" binding:"required" example:"1"`
	Value string `json:"value" binding:"required" example:"Red"`
}
