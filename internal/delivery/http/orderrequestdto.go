package http

import (
	"backend/internal/domain"

	"github.com/google/uuid"
)

type ListOrderRequestDto struct {
	PaginationRequestDto
	IDs     []uuid.UUID
	UserIDs []uuid.UUID
	Status  domain.OrderStatus
}

type CreateOrderRequestDto struct {
	Data CreateOrderData
}

type CreateOrderData struct {
	RecipientName string                `json:"recipientName" binding:"required"`
	PhoneNumber   string                `json:"phoneNumber"   binding:"required"`
	Address       string                `json:"address"       binding:"required"`
	Provider      domain.OrderProvider  `json:"provider"      binding:"required"`
	Items         []CreateOrderItemData `json:"items"         binding:"required,dive"`
	UserID        uuid.UUID             `json:"userId"        binding:"required"`
	ReturnURL     string                `json:"returnUrl"`
}

type CreateOrderItemData struct {
	ProductID        uuid.UUID `json:"productId"        binding:"required"`
	ProductVariantID uuid.UUID `json:"productVariantId" binding:"required"`
	Quantity         int       `json:"quantity"         binding:"required"`
}

type UpdateOrderRequestDto struct {
	OrderID uuid.UUID
	Data    UpdateOrderData
}

type UpdateOrderData struct {
	Address string             `json:"address" binding:"required"`
	Status  domain.OrderStatus `json:"status"  binding:"required"`
	IsPaid  bool               `json:"is_paid" binding:"required"`
}

type GetOrderRequestDto struct {
	OrderID uuid.UUID
}

type DeleteOrderRequestDto struct {
	OrderID uuid.UUID
}

type VerifyVNPayIPNRequestDTO struct {
	QueryParams *VerifyVNPayIPNQueryParams
}

type VerifyVNPayIPNQueryParams struct {
	Amount            string `form:"vnp_Amount"            binding:"required"`
	BankTranNo        string `form:"vnp_BankTranNo"        binding:"required"`
	BankCode          string `form:"vnp_BankCode"          binding:"required"`
	CardType          string `form:"vnp_CardType"          binding:"required"`
	OrderInfo         string `form:"vnp_OrderInfo"         binding:"required"`
	PayDate           string `form:"vnp_PayDate"           binding:"required"`
	ResponseCode      string `form:"vnp_ResponseCode"      binding:"required"`
	SecureHash        string `form:"vnp_SecureHash"        binding:"required"`
	TmnCode           string `form:"vnp_TmnCode"           binding:"required"`
	TransactionNo     string `form:"vnp_TransactionNo"     binding:"required"`
	TransactionStatus string `form:"vnp_TransactionStatus" binding:"required"`
	TxnRef            string `form:"vnp_TxnRef"            binding:"required"`
}

type VerifyVNPayIPNResponseDTO struct {
	RspCode string `json:"RspCode" binding:"required"`
	Message string `json:"Message" binding:"required"`
}
