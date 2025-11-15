package response

import "time"

type PaymentProvider string

const (
	PaymentProviderCOD     PaymentProvider = "COD"
	PaymentProviderVNPAY   PaymentProvider = "VNPAY"
	PaymentProviderMOMO    PaymentProvider = "MOMO"
	PaymentProviderZALOPAY PaymentProvider = "ZALO"
)

type PaymentStatus string

const (
	PaymentStatusPending   PaymentStatus = "PENDING"
	PaymentStatusCompleted PaymentStatus = "COMPLETED"
	PaymentStatusFailed    PaymentStatus = "FAILED"
)

type Payment struct {
	ID              int             `json:"id" binding:"required"`
	Amount          int64           `json:"amount" binding:"required"`
	PaymentStatus   PaymentStatus   `json:"paymentStatus" binding:"required"`
	PaymentProvider PaymentProvider `json:"paymentProvider" binding:"omitempty"`
	UpdatedAt       time.Time       `json:"updatedAt" binding:"required"`
}
