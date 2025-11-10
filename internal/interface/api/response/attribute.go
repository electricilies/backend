package response

type Attribute struct {
	ID              string           `json:"id" binding:"required"`
	Code            string           `json:"code" binding:"required"`
	Name            string           `json:"name" binding:"required"`
	AttributeValues []AttributeValue `json:"attributeValues" binding:"required"`
}

type AttributeValue struct {
	ID    string `json:"id" binding:"required"`
	Value string `json:"value" binding:"required"`
}

type AttributesPagination struct {
	Meta Pagination  `json:"meta" binding:"required"`
	Data []Attribute `json:"data" binding:"required"`
}
