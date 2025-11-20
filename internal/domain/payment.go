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
	ID        int             `json:"id"        binding:"required" validate:"required"`
	Amount    int64           `json:"amount"    binding:"required" validate:"required,gt=0"`
	Order     *Order          `json:"order"     validate:"omitnil"`
	Status    PaymentStatus   `json:"status"    binding:"required" validate:"required,oneof=Pending Completed Failed"`
	Provider  PaymentProvider `json:"provider"  binding:"required" validate:"required,oneof=COD VNPAY MOMO ZALOPAY"`
	UpdatedAt time.Time       `json:"updatedAt" binding:"required" validate:"required"`
}
