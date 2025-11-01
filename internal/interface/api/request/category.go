package request

type CreateCategory struct {
    Name string `json:"name" binding:"required"`
}

type UpdateCategory struct {
    Name string `json:"name" binding:"required"`
}
