// vim: tabstop=4:
package domain_test

import (
	"testing"
	"time"

	"backend/internal/domain"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
)

type OrderTestSuite struct {
	suite.Suite
	validate *validator.Validate
}

func (s *OrderTestSuite) SetupSuite() {
	s.validate = validator.New(validator.WithRequiredStructEnabled())
	err := domain.RegisterOrderValidates(s.validate)
	s.Require().NoError(err)
}

func (s *OrderTestSuite) TestNewOrderItemBoundaryValues() {
	s.T().Parallel()
	testcases := []struct {
		name      string
		quantity  int
		price     int64
		expectOk  bool
		expectErr bool
	}{
		{
			name:      "quantity 0 (invalid)",
			quantity:  0,
			price:     1000,
			expectOk:  false,
			expectErr: true,
		},
		{
			name:      "quantity 1 (min)",
			quantity:  1,
			price:     1000,
			expectOk:  true,
			expectErr: false,
		},
		{
			name:      "quantity 2 (min + 1)",
			quantity:  2,
			price:     1000,
			expectOk:  true,
			expectErr: false,
		},
		{
			name:      "quantity 50",
			quantity:  50,
			price:     1000,
			expectOk:  true,
			expectErr: false,
		},
		{
			name:      "quantity 99 (max)",
			quantity:  99,
			price:     1000,
			expectOk:  true,
			expectErr: false,
		},
		{
			name:      "quantity 100 (max + 1, invalid)",
			quantity:  100,
			price:     1000,
			expectOk:  false,
			expectErr: true,
		},
		{
			name:      "quantity -1 (negative)",
			quantity:  -1,
			price:     1000,
			expectOk:  false,
			expectErr: true,
		},
		{
			name:      "price 0 (invalid)",
			quantity:  1,
			price:     0,
			expectOk:  false,
			expectErr: true,
		},
		{
			name:      "price 1 (min)",
			quantity:  1,
			price:     1,
			expectOk:  true,
			expectErr: false,
		},
		{
			name:      "price -1 (negative)",
			quantity:  1,
			price:     -1,
			expectOk:  false,
			expectErr: true,
		},
		{
			name:      "valid order item",
			quantity:  5,
			price:     10000,
			expectOk:  true,
			expectErr: false,
		},
	}

	productID := uuid.New()
	variantID := uuid.New()

	for _, tc := range testcases {
		s.Run(tc.name, func() {
			orderItem, err := domain.NewOrderItem(productID, variantID, tc.quantity, tc.price)

			s.NoError(err, tc.name)
			s.NotNil(orderItem, tc.name)
			s.Equal(tc.quantity, orderItem.Quantity, tc.name)
			s.Equal(tc.price, orderItem.Price, tc.name)
			s.Equal(productID, orderItem.ProductID, tc.name)
			s.Equal(variantID, orderItem.ProductVariantID, tc.name)
			s.NotNil(orderItem.ID, tc.name)

			validationErr := s.validate.Struct(orderItem)
			if tc.expectErr {
				s.Error(validationErr, tc.name)
			} else {
				s.NoError(validationErr, tc.name)
			}
		})
	}
}

func (s *OrderTestSuite) TestNewOrderTotalAmountCalculation() {
	s.T().Parallel()
	testcases := []struct {
		name  string
		items []struct {
			quantity int
			price    int64
		}
		expectedTotal     int64
		expectValidation  bool
		validationMessage string
	}{
		{
			name: "empty items (invalid)",
			items: []struct {
				quantity int
				price    int64
			}{},
			expectedTotal:     0,
			expectValidation:  false,
			validationMessage: "items must have at least 1 item",
		},
		{
			name: "single item",
			items: []struct {
				quantity int
				price    int64
			}{
				{quantity: 2, price: 100},
			},
			expectedTotal:    200,
			expectValidation: true,
		},
		{
			name: "multiple items",
			items: []struct {
				quantity int
				price    int64
			}{
				{quantity: 2, price: 100},
				{quantity: 3, price: 200},
			},
			expectedTotal:    800,
			expectValidation: true,
		},
		{
			name: "single item large quantity",
			items: []struct {
				quantity int
				price    int64
			}{
				{quantity: 10, price: 5000},
			},
			expectedTotal:    50000,
			expectValidation: true,
		},
		{
			name: "multiple items with different prices",
			items: []struct {
				quantity int
				price    int64
			}{
				{quantity: 1, price: 100},
				{quantity: 2, price: 250},
				{quantity: 3, price: 500},
			},
			expectedTotal:    2100,
			expectValidation: true,
		},
	}

	userID := uuid.New()

	for _, tc := range testcases {
		s.Run(tc.name, func() {
			var orderItems []domain.OrderItem
			for _, item := range tc.items {
				orderItem, err := domain.NewOrderItem(uuid.New(), uuid.New(), item.quantity, item.price)
				s.Require().NoError(err)
				orderItems = append(orderItems, *orderItem)
			}

			order, err := domain.NewOrder(
				userID,
				"John Doe",
				"+84901234567",
				"123 Street",
				domain.PaymentProviderCOD,
				orderItems,
			)

			s.NoError(err, tc.name)
			s.NotNil(order, tc.name)
			s.Equal(tc.expectedTotal, order.TotalAmount, tc.name)
			s.Equal(userID, order.UserID, tc.name)
			s.Equal(domain.OrderStatusPending, order.Status, tc.name)
			s.False(order.IsPaid, tc.name)
			s.NotNil(order.ID, tc.name)
			s.NotNil(order.CreatedAt, tc.name)
			s.NotNil(order.UpdatedAt, tc.name)

			validationErr := s.validate.Struct(order)
			if tc.expectValidation {
				s.NoError(validationErr, tc.name)
			} else {
				s.Error(validationErr, tc.name)
			}
		})
	}
}

func (s *OrderTestSuite) TestOrderUpdate() {
	s.T().Parallel()
	testcases := []struct {
		name                string
		initialAddress      string
		initialStatus       domain.OrderStatus
		initialIsPaid       bool
		updateAddress       string
		updateStatus        domain.OrderStatus
		updateIsPaid        bool
		expectUpdatedAtDiff bool
	}{
		{
			name:                "update address only",
			initialAddress:      "123 Old Street",
			initialStatus:       domain.OrderStatusPending,
			initialIsPaid:       false,
			updateAddress:       "456 New Street",
			updateStatus:        domain.OrderStatusPending,
			updateIsPaid:        false,
			expectUpdatedAtDiff: true,
		},
		{
			name:                "update status only",
			initialAddress:      "123 Street",
			initialStatus:       domain.OrderStatusPending,
			initialIsPaid:       false,
			updateAddress:       "123 Street",
			updateStatus:        domain.OrderStatusProcessing,
			updateIsPaid:        false,
			expectUpdatedAtDiff: true,
		},
		{
			name:                "update isPaid only",
			initialAddress:      "123 Street",
			initialStatus:       domain.OrderStatusPending,
			initialIsPaid:       false,
			updateAddress:       "123 Street",
			updateStatus:        domain.OrderStatusPending,
			updateIsPaid:        true,
			expectUpdatedAtDiff: true,
		},
		{
			name:                "update all fields",
			initialAddress:      "123 Old Street",
			initialStatus:       domain.OrderStatusPending,
			initialIsPaid:       false,
			updateAddress:       "456 New Street",
			updateStatus:        domain.OrderStatusDelivered,
			updateIsPaid:        true,
			expectUpdatedAtDiff: true,
		},
		{
			name:                "no changes",
			initialAddress:      "123 Street",
			initialStatus:       domain.OrderStatusPending,
			initialIsPaid:       false,
			updateAddress:       "123 Street",
			updateStatus:        domain.OrderStatusPending,
			updateIsPaid:        false,
			expectUpdatedAtDiff: false,
		},
	}

	for _, tc := range testcases {
		s.Run(tc.name, func() {
			orderItem, err := domain.NewOrderItem(uuid.New(), uuid.New(), 1, 1000)
			s.Require().NoError(err)

			order, err := domain.NewOrder(
				uuid.New(),
				"John Doe",
				"+84901234567",
				tc.initialAddress,
				domain.PaymentProviderCOD,
				[]domain.OrderItem{*orderItem},
			)
			s.Require().NoError(err)

			order.Status = tc.initialStatus
			order.IsPaid = tc.initialIsPaid
			originalUpdatedAt := order.UpdatedAt

			time.Sleep(10 * time.Millisecond)

			order.Update(tc.updateAddress, tc.updateStatus, tc.updateIsPaid)

			s.Equal(tc.updateAddress, order.Address, tc.name)
			s.Equal(tc.updateStatus, order.Status, tc.name)
			s.Equal(tc.updateIsPaid, order.IsPaid, tc.name)

			if tc.expectUpdatedAtDiff {
				s.True(order.UpdatedAt.After(originalUpdatedAt), "UpdatedAt should be updated")
			} else {
				s.Equal(originalUpdatedAt, order.UpdatedAt, "UpdatedAt should not change")
			}
		})
	}
}

func (s *OrderTestSuite) TestOrderStatusTransitions() {
	s.T().Parallel()
	testcases := []struct {
		name          string
		initialStatus domain.OrderStatus
		newStatus     domain.OrderStatus
	}{
		{
			name:          "pending to processing",
			initialStatus: domain.OrderStatusPending,
			newStatus:     domain.OrderStatusProcessing,
		},
		{
			name:          "processing to shipping",
			initialStatus: domain.OrderStatusProcessing,
			newStatus:     domain.OrderStatusShipping,
		},
		{
			name:          "shipping to delivered",
			initialStatus: domain.OrderStatusShipping,
			newStatus:     domain.OrderStatusDelivered,
		},
		{
			name:          "pending to cancelled",
			initialStatus: domain.OrderStatusPending,
			newStatus:     domain.OrderStatusCancelled,
		},
	}

	for _, tc := range testcases {
		s.Run(tc.name, func() {
			orderItem, err := domain.NewOrderItem(uuid.New(), uuid.New(), 1, 1000)
			s.Require().NoError(err)

			order, err := domain.NewOrder(
				uuid.New(),
				"John Doe",
				"+84901234567",
				"123 Street",
				domain.PaymentProviderCOD,
				[]domain.OrderItem{*orderItem},
			)
			s.Require().NoError(err)

			order.Status = tc.initialStatus
			originalUpdatedAt := order.UpdatedAt

			time.Sleep(10 * time.Millisecond)

			order.Update(order.Address, tc.newStatus, order.IsPaid)

			s.Equal(tc.newStatus, order.Status, tc.name)
			s.True(order.UpdatedAt.After(originalUpdatedAt), "UpdatedAt should be updated")
		})
	}
}

func TestOrder(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(OrderTestSuite))
}
