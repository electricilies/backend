package request

type CreateAttribute struct {
	Code  string `json:"code" binding:"required"`
	Name  string `json:"name" binding:"required"`
	Value string `json:"value" binding:"required"`
}

type UpdateAttribute struct {
	Name string `json:"name" binding:"required"`
}

type CreateAttributeValue struct {
	Value string `json:"value" binding:"required"`
}
