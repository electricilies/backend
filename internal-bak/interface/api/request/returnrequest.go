package request

type ReturnRequestStatus string

type CreateReturnRequest struct {
	Reason      string `json:"reason" binding:"required"`
	OrderItemID int    `json:"orderItemId" binding:"required"`
	UserID      string `json:"userId" binding:"required"`
}

type UpdateReturnRequestStatus struct {
	Status ReturnRequestStatus `json:"returnRequestStatus" binding:"required"`
}
