package response

import "time"

type Order struct {
    ID            int         `json:"id" binding:"required"`
    UserID        string      `json:"userId" binding:"required"`
    OrderStatusID int         `json:"orderStatusId" binding:"required"`
    PaymentID     int         `json:"paymentId" binding:"required"`
    CreatedAt     time.Time   `json:"createdAt" binding:"required"`
    UpdatedAt     time.Time   `json:"updatedAt" binding:"required"`
    Items         []OrderItem `json:"items" binding:"required"`
}

type OrderItem struct {
    ID             int            `json:"id" binding:"required"`
    OrderID        int            `json:"orderId" binding:"required"`
    Product        Product        `json:"product" binding:"required"`
    ProductVariant ProductVariant `json:"productVariant" binding:"required"`
    Quantity       int            `json:"quantity" binding:"required"`
    Price          float64        `json:"price" binding:"required"`
}
