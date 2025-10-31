package response

import "time"

type (
	PaymentMethod   string
	PaymentStatus   string
	PaymentProvider string
)

type Payment struct {
	ID                int             `json:"id" binding:"required"`
	Amount            float64         `json:"amount" binding:"required"`
	PaymentMethod     PaymentMethod   `json:"payment_method_id" binding:"required"`
	PaymentStatus     PaymentStatus   `json:"payment_status_id" binding:"required"`
	PaymentProviderID PaymentProvider `json:"payment_provider_id" binding:"omitempty"`
	UpdatedAt         time.Time       `json:"updated_at" binding:"required"`
}
