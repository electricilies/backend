package response

import "time"

type PaymentMethod struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Payment struct {
	ID                int        `json:"id"`
	Amount            float64    `json:"amount"`
	PaymentMethodID   int        `json:"payment_method_id"`
	PaymentStatusID   int        `json:"payment_status_id"`
	PaymentProviderID int        `json:"payment_provider_id"`
	UpdatedAt         *time.Time `json:"updated_at"`
}
