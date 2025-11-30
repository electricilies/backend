// vim: tabstop=4:

package domain_test

import (
	"testing"

	"backend/internal/domain"

	"github.com/stretchr/testify/assert"
)

func TestValidateOrderTotalAmount(t *testing.T) {
	testcases := []struct {
		name     string
		order    domain.Order
		expectOk bool
	}{
		{
			name: "valid total amount",
			order: domain.Order{
				Items: []domain.OrderItem{
					{Price: 100},
					{Price: 200},
				},
				TotalAmount: 300,
			},
			expectOk: true,
		},
		{
			name: "invalid total amount",
			order: domain.Order{
				Items: []domain.OrderItem{
					{Price: 100},
					{Price: 200},
				},
				TotalAmount: 400,
			},
			expectOk: false,
		},
	}
	for _, tc := range testcases {
		ok := domain.ValidateOrderTotalAmount(&tc.order)
		assert.Equal(t, tc.expectOk, ok, tc.name)

	}
}
