package response

import "time"

type Review struct {
    ID        int       `json:"id" binding:"required"`
    Rate      int       `json:"rate" binding:"required"`
    Content   string    `json:"content" binding:"omitnil"`
    ImageURL  string    `json:"imageUrl" binding:"omitnil"`
    User      User      `json:"user" binding:"required"`
    CreatedAt time.Time `json:"createdAt" binding:"required"`
    UpdatedAt time.Time `json:"updatedAt" binding:"required"`
}
