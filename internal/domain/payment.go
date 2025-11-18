package domain

import "time"

type PaymentProvider string

const (
	PaymentProviderCOD     PaymentProvider = "COD"
	PaymentProviderVNPAY   PaymentProvider = "VNPAY"
	PaymentProviderMOMO    PaymentProvider = "MOMO"
	PaymentProviderZALOPAY PaymentProvider = "ZALOPAY"
)

type PaymentStatus string

const (
	PaymentStatusPending   PaymentStatus = "Pending"
	PaymentStatusCompleted PaymentStatus = "Completed"
	PaymentStatusFailed    PaymentStatus = "Failed"
)

type Payment struct {
	ID        int             `json:"id" binding:"required"`
	Amount    int64           `json:"amount" binding:"required"`
	Order     *Order          `json:"order" binding:"omitnil"`
	Status    PaymentStatus   `json:"status" binding:"required,oneof=Pending Completed Failed"`
	Provider  PaymentProvider `json:"provider" binding:"required,oneof=COD VNPAY MOMO ZALOPAY"`
	UpdatedAt time.Time       `json:"updatedAt" binding:"required"`
}
