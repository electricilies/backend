package request

type CreateReturnRequest struct {
	Reason      string `json:"reason" binding:"required"`
	OrderItemID int    `json:"order_item_id" binding:"required"`
	UserID      string `json:"user_id" binding:"required"`
}

type UpdateReturnStatus struct {
	StatusID int `json:"status_id" binding:"required"`
}
