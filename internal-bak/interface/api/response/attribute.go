package response

type Attribute struct {
	ID              string           `json:"id" binding:"required" example:"123"`
	Code            string           `json:"code" binding:"required" example:"color"`
	Name            string           `json:"name" binding:"required" example:"Color"`
	AttributeValues []AttributeValue `json:"attributeValues" binding:"required"`
}

type AttributeValue struct {
	ID    string `json:"id" binding:"required" example:"1"`
	Value string `json:"value" binding:"required" example:"Red"`
}
