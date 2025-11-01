package response

import "time"

type RefundStatus string

type Refund struct {
    ID            int           `json:"id" binding:"required"`
    Status        RefundStatus  `json:"statusId" binding:"required"`
    Payment       Payment       `json:"payment" binding:"required"`
    ReturnRequest ReturnRequest `json:"returnRequest" binding:"required"`
    CreatedAt     time.Time     `json:"createdAt" binding:"required"`
    UpdatedAt     time.Time     `json:"updatedAt" binding:"required"`
}
