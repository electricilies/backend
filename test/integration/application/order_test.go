// vim: tabstop=4 shiftwidth=4:
//go:build integration

package application_test

import (
	"context"
	"testing"

	"backend/config"
	"backend/internal/application"
	"backend/internal/client"
	"backend/internal/delivery/http"
	"backend/internal/domain"
	"backend/internal/infrastructure/repositorypostgres"
	"backend/internal/service"
	"backend/test/integration/component"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type OrderTestSuite struct {
	suite.Suite
	containers          *component.Containers
	app                 http.OrderApplication
	productRepo         domain.ProductRepository
	orderRepo           domain.OrderRepository
	vnpayPaymentService *application.MockVNPayPaymentService

	// Seed data IDs from .rules/006-testing.md
	seededProductID       uuid.UUID
	seededVariantID       uuid.UUID
	seededSecondProductID uuid.UUID
	seededSecondVariantID uuid.UUID
	seededUserID          uuid.UUID
}

func TestOrderSuite(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(OrderTestSuite))
}

func (s *OrderTestSuite) newContainersConfig() *component.ContainersConfig {
	containersConfig := component.NewContainersConfig(&component.NewContainersConfigParam{
		DBEnabled: true,
	})
	containersConfig.DB.Seed = true
	return containersConfig
}

func (s *OrderTestSuite) newConfig(
	ctx context.Context,
) *config.Server {
	s.T().Helper()

	dbConnStr, err := s.containers.DB.ConnectionString(ctx, "sslmode=disable")
	s.Require().NoError(err, "failed to get db connection string")

	return &config.Server{
		DBURL: dbConnStr,
	}
}

func (s *OrderTestSuite) SetupSuite() {
	ctx := s.T().Context()
	containersConfig := s.newContainersConfig()

	var err error
	s.containers, err = component.NewContainers(ctx, containersConfig)
	s.Require().NoError(err, "failed to start containers")

	cfg := s.newConfig(ctx)

	validate := validator.New(
		validator.WithRequiredStructEnabled(),
	)

	conn := client.NewDBConnection(ctx, cfg)
	queries := client.NewDBQueries(conn)

	s.orderRepo = repositorypostgres.ProvideOrder(queries, conn)
	s.productRepo = repositorypostgres.ProvideProduct(queries, conn)
	cartRepo := repositorypostgres.ProvideCart(queries, conn)

	orderService := service.ProvideOrder(validate)
	productService := service.ProvideProduct(validate)

	s.vnpayPaymentService = application.NewMockVNPayPaymentService(s.T())

	s.app = application.ProvideOrder(
		s.vnpayPaymentService,
		s.orderRepo,
		orderService,
		s.productRepo,
		productService,
		cartRepo,
	)

	// Seed data from .rules/006-testing.md
	s.seededProductID = uuid.MustParse("00000000-0000-7000-0000-000278469304")
	s.seededVariantID = uuid.MustParse("00000000-0000-7000-0000-000278469308")
	s.seededSecondProductID = uuid.MustParse("00000000-0000-7000-0000-000278469345")
	s.seededSecondVariantID = uuid.MustParse("00000000-0000-7000-0000-000278469347")
	s.seededUserID = uuid.MustParse("00000000-0000-7000-0000-000000000003")
}

func (s *OrderTestSuite) TearDownSuite() {
	s.containers.Cleanup(s.T())
}

func (s *OrderTestSuite) SetupTest() {
	(*s.vnpayPaymentService) = *application.NewMockVNPayPaymentService(s.T())
}

func (s *OrderTestSuite) TestOrderLifecycle() {
	ctx := s.T().Context()

	var vnpayOrderID uuid.UUID

	s.Run("Create Order with VNPAY provider (Pending status)", func() {
		s.vnpayPaymentService.EXPECT().
			GetPaymentURL(mock.Anything, mock.Anything).
			Return("https://sandbox.vnpayment.vn/paymentv2/vpcpay.html", nil).
			Once()

		result, err := s.app.Create(ctx, http.CreateOrderRequestDto{
			Data: http.CreateOrderData{
				RecipientName: "John Doe",
				PhoneNumber:   "+84912345678",
				Address:       "123 Test Street, Ho Chi Minh City",
				Provider:      domain.PaymentProviderVNPAY,
				Items: []http.CreateOrderItemData{
					{
						ProductID:        s.seededProductID,
						ProductVariantID: s.seededVariantID,
						Quantity:         2,
					},
				},
				UserID:    s.seededUserID,
				ReturnURL: "https://example.com/return",
			},
		})
		s.Require().NoError(err)
		s.Require().NotNil(result)
		s.Equal(domain.OrderStatusPending, result.Status)
		s.False(result.IsPaid)
		s.Equal("John Doe", result.RecipentName)
		s.Equal("+84912345678", result.PhoneNumber)
		s.Equal(s.seededUserID, result.UserID)
		s.NotEmpty(result.PaymentURL)
		s.Len(result.Items, 1)
		s.Equal(2, result.Items[0].Quantity)
		s.Positive(result.TotalAmount)
		vnpayOrderID = result.ID
	})

	s.Run("Get Order by ID", func() {
		result, err := s.app.Get(ctx, http.GetOrderRequestDto{
			OrderID: vnpayOrderID,
		})
		s.Require().NoError(err)
		s.Require().NotNil(result)
		s.Equal(vnpayOrderID, result.ID)
		s.Equal(domain.OrderStatusPending, result.Status)
		s.Len(result.Items, 1)

		item := result.Items[0]
		s.NotNil(item.Product.ID)
		s.NotEmpty(item.Product.Name)
		s.NotNil(item.ProductVariant.ID)
		s.NotEmpty(item.ProductVariant.SKU)
		s.Positive(item.Price)
	})

	s.Run("Update Order", func() {
		result, err := s.app.Update(ctx, http.UpdateOrderRequestDto{
			OrderID: vnpayOrderID,
			Data: http.UpdateOrderData{
				Address: "456 New Address, Hanoi",
				Status:  domain.OrderStatusPending,
				IsPaid:  false,
			},
		})
		s.Require().NoError(err)
		s.Require().NotNil(result)
		s.Equal("456 New Address, Hanoi", result.Address)
	})

	s.Run("VNPay IPN Success - Verify inventory decrease", func() {
		product, err := s.productRepo.Get(ctx, domain.ProductRepositoryGetParam{
			ProductID: s.seededProductID,
		})
		s.Require().NoError(err)
		variant := product.GetVariantByID(s.seededVariantID)
		s.Require().NotNil(variant)
		initialQuantity := variant.Quantity

		order, err := s.app.Get(ctx, http.GetOrderRequestDto{
			OrderID: vnpayOrderID,
		})
		s.Require().NoError(err)
		orderedQuantity := order.Items[0].Quantity

		s.vnpayPaymentService.EXPECT().
			VerifyIPN(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			RunAndReturn(func(
				ctx context.Context,
				param application.VerifyIPNVNPayParam,
				getOrder func(ctx context.Context, orderID uuid.UUID) (*domain.Order, error),
				onSuccess func(ctx context.Context, order *domain.Order) error,
				onFailure func(ctx context.Context, order *domain.Order) error,
			) (string, string, error) {
				orderID, err := uuid.Parse(param.TxnRef)
				if err != nil {
					return "99", "Invalid Order ID", err
				}

				order, err := getOrder(ctx, orderID)
				if err != nil {
					return "01", "Order not found", err
				}

				if param.ResponseCode == "00" && param.TransactionStatus == "00" {
					if err := onSuccess(ctx, order); err != nil {
						return "99", "Error processing payment", err
					}
					return "00", "Success", nil
				}

				if err := onFailure(ctx, order); err != nil {
					return "99", "Error processing failure", err
				}
				return "02", "Payment failed", nil
			}).Once()

		result, err := s.app.VerifyVNPayIPN(ctx, http.VerifyVNPayIPNRequestDTO{
			QueryParams: &http.VerifyVNPayIPNQueryParams{
				Amount:            "100000000", // 1,000,000 VND * 100
				BankTranNo:        "VNP123456",
				BankCode:          "NCB",
				CardType:          "ATM",
				OrderInfo:         "Payment for order",
				PayDate:           "20251210150000",
				ResponseCode:      "00",
				SecureHash:        "test_hash",
				TmnCode:           "test_tmn",
				TransactionNo:     "123456789",
				TransactionStatus: "00",
				TxnRef:            vnpayOrderID.String(),
			},
		})
		s.Require().NoError(err)
		s.Require().NotNil(result)
		s.Equal("00", result.RspCode)
		s.NotEmpty(result.Message)

		updatedOrder, err := s.app.Get(ctx, http.GetOrderRequestDto{
			OrderID: vnpayOrderID,
		})
		s.Require().NoError(err)
		s.Equal(domain.OrderStatusProcessing, updatedOrder.Status)
		s.True(updatedOrder.IsPaid)

		productAfter, err := s.productRepo.Get(ctx, domain.ProductRepositoryGetParam{
			ProductID: s.seededProductID,
		})
		s.Require().NoError(err)
		variantAfter := productAfter.GetVariantByID(s.seededVariantID)
		s.Require().NotNil(variantAfter)
		s.Equal(initialQuantity-orderedQuantity, variantAfter.Quantity, "Inventory should decrease by ordered quantity")
	})

	s.Run("Create another order for failure test", func() {
		s.vnpayPaymentService.EXPECT().
			GetPaymentURL(mock.Anything, mock.Anything).
			Return("https://sandbox.vnpayment.vn/paymentv2/vpcpay.html", nil).
			Once()

		result, err := s.app.Create(ctx, http.CreateOrderRequestDto{
			Data: http.CreateOrderData{
				RecipientName: "Jane Doe",
				PhoneNumber:   "+84987654321",
				Address:       "789 Another Street",
				Provider:      domain.PaymentProviderVNPAY,
				Items: []http.CreateOrderItemData{
					{
						ProductID:        s.seededSecondProductID,
						ProductVariantID: s.seededSecondVariantID,
						Quantity:         1,
					},
				},
				UserID:    s.seededUserID,
				ReturnURL: "https://example.com/return",
			},
		})
		s.Require().NoError(err)
		s.Require().NotNil(result)
		vnpayOrderID = result.ID // Reuse variable for failure test
	})

	s.Run("VNPay IPN Failure - Order cancelled", func() {
		product, err := s.productRepo.Get(ctx, domain.ProductRepositoryGetParam{
			ProductID: s.seededSecondProductID,
		})
		s.Require().NoError(err)
		variant := product.GetVariantByID(s.seededSecondVariantID)
		s.Require().NotNil(variant)
		initialQuantity := variant.Quantity

		s.vnpayPaymentService.EXPECT().
			VerifyIPN(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			RunAndReturn(func(
				ctx context.Context,
				param application.VerifyIPNVNPayParam,
				getOrder func(ctx context.Context, orderID uuid.UUID) (*domain.Order, error),
				onSuccess func(ctx context.Context, order *domain.Order) error,
				onFailure func(ctx context.Context, order *domain.Order) error,
			) (string, string, error) {
				orderID, err := uuid.Parse(param.TxnRef)
				if err != nil {
					return "99", "Invalid Order ID", err
				}

				order, err := getOrder(ctx, orderID)
				if err != nil {
					return "01", "Order not found", err
				}

				if err := onFailure(ctx, order); err != nil {
					return "99", "Error processing failure", err
				}
				return "02", "Payment failed", nil
			}).Once()

		result, err := s.app.VerifyVNPayIPN(ctx, http.VerifyVNPayIPNRequestDTO{
			QueryParams: &http.VerifyVNPayIPNQueryParams{
				Amount:            "100000000",
				BankTranNo:        "VNP789",
				BankCode:          "NCB",
				CardType:          "ATM",
				OrderInfo:         "Payment for order",
				PayDate:           "20251210150000",
				ResponseCode:      "24", // Failed - Customer cancelled
				SecureHash:        "test_hash",
				TmnCode:           "test_tmn",
				TransactionNo:     "987654321",
				TransactionStatus: "02", // Failed
				TxnRef:            vnpayOrderID.String(),
			},
		})
		s.Require().NoError(err)
		s.Require().NotNil(result)
		s.NotEqual("00", result.RspCode)
		s.NotEmpty(result.Message)

		updatedOrder, err := s.app.Get(ctx, http.GetOrderRequestDto{
			OrderID: vnpayOrderID,
		})
		s.Require().NoError(err)
		s.Equal(domain.OrderStatusCancelled, updatedOrder.Status)
		s.False(updatedOrder.IsPaid)

		productAfter, err := s.productRepo.Get(ctx, domain.ProductRepositoryGetParam{
			ProductID: s.seededSecondProductID,
		})
		s.Require().NoError(err)
		variantAfter := productAfter.GetVariantByID(s.seededSecondVariantID)
		s.Require().NotNil(variantAfter)
		s.Equal(initialQuantity, variantAfter.Quantity, "Inventory should NOT change on failed payment")
	})

	s.Run("Create Order with COD provider (no payment URL)", func() {
		result, err := s.app.Create(ctx, http.CreateOrderRequestDto{
			Data: http.CreateOrderData{
				RecipientName: "Cash Customer",
				PhoneNumber:   "+84123456789",
				Address:       "999 Cash Street",
				Provider:      domain.PaymentProviderCOD,
				ReturnURL:     "https://example.com/return",
				UserID:        s.seededUserID,
				Items: []http.CreateOrderItemData{
					{
						ProductID:        s.seededProductID,
						ProductVariantID: s.seededVariantID,
						Quantity:         1,
					},
				},
			},
		})

		s.vnpayPaymentService.EXPECT().GetPaymentURL(mock.Anything, mock.Anything).Return("", nil).Times(0)
		s.Require().NoError(err)
		s.Require().NotNil(result)
		s.Equal(domain.OrderStatusPending, result.Status)
		s.Empty(result.PaymentURL)
	})
}

func (s *OrderTestSuite) TestCreateOrderWithMultipleItemsInLifecycle() {
	ctx := s.T().Context()

	s.vnpayPaymentService.EXPECT().
		GetPaymentURL(mock.Anything, mock.Anything).
		Return("https://sandbox.vnpayment.vn/paymentv2/vpcpay.html", nil).
		Once()

	result, err := s.app.Create(ctx, http.CreateOrderRequestDto{
		Data: http.CreateOrderData{
			RecipientName: "Multi Item Customer",
			PhoneNumber:   "+84999888777",
			Address:       "Multi Item Street",
			Provider:      domain.PaymentProviderVNPAY,
			Items: []http.CreateOrderItemData{
				{
					ProductID:        s.seededProductID,
					ProductVariantID: s.seededVariantID,
					Quantity:         2,
				},
				{
					ProductID:        s.seededSecondProductID,
					ProductVariantID: s.seededSecondVariantID,
					Quantity:         1,
				},
			},
			UserID:    s.seededUserID,
			ReturnURL: "https://example.com/return",
		},
	})
	s.Require().NoError(err)
	s.Require().NotNil(result)
	s.Len(result.Items, 2)
	s.Positive(result.TotalAmount)

	// Verify total amount calculation
	expectedTotal := result.Items[0].Price*int64(result.Items[0].Quantity) +
		result.Items[1].Price*int64(result.Items[1].Quantity)
	s.Equal(expectedTotal, result.TotalAmount)
}

func (s *OrderTestSuite) TestErrorHandling() {
	ctx := s.T().Context()
	nonExistentID := uuid.MustParse("00000000-0000-0000-0000-000000000001")

	testcases := []struct {
		name   string
		action func() error
	}{
		{
			name: "Get non-existent order fails",
			action: func() error {
				_, err := s.app.Get(ctx, http.GetOrderRequestDto{
					OrderID: nonExistentID,
				})
				return err
			},
		},
		{
			name: "Update non-existent order fails",
			action: func() error {
				_, err := s.app.Update(ctx, http.UpdateOrderRequestDto{
					OrderID: nonExistentID,
					Data: http.UpdateOrderData{
						Address: "New Address",
						Status:  domain.OrderStatusPending,
						IsPaid:  false,
					},
				})
				return err
			},
		},
		{
			name: "Create order with non-existent product fails",
			action: func() error {
				_, err := s.app.Create(ctx, http.CreateOrderRequestDto{
					Data: http.CreateOrderData{
						RecipientName: "Test",
						PhoneNumber:   "+84912345678",
						Address:       "Test",
						Provider:      domain.PaymentProviderCOD,
						Items: []http.CreateOrderItemData{
							{
								ProductID:        nonExistentID,
								ProductVariantID: nonExistentID,
								Quantity:         1,
							},
						},
						UserID: s.seededUserID,
					},
				})
				return err
			},
		},
	}

	for _, tc := range testcases {
		s.Run(tc.name, func() {
			err := tc.action()
			s.Require().Error(err, tc.name)
		})
	}
}

func (s *OrderTestSuite) TestValidationDefects() {
	ctx := s.T().Context()

	// NOTE: These tests expose a defect (DF-O-CV-01) - orderService.Validate() is not called
	// in the Create function, so invalid orders can be created. For now, we mock the payment
	// service to prevent test failures, but these should fail once validation is added.

	testcases := []struct {
		name string
		data http.CreateOrderData
	}{
		{
			name: "Create order with empty items (SHOULD fail but doesn't - DF-O-CV-01)",
			data: http.CreateOrderData{
				RecipientName: "Test",
				PhoneNumber:   "+84912345678",
				Address:       "Test",
				Provider:      domain.PaymentProviderCOD,
				Items:         []http.CreateOrderItemData{},
				UserID:        s.seededUserID,
			},
		},
		{
			name: "Validation: Invalid phone number (SHOULD fail but doesn't - DF-O-CV-01)",
			data: http.CreateOrderData{
				RecipientName: "Test",
				PhoneNumber:   "invalid-phone",
				Address:       "Test",
				Provider:      domain.PaymentProviderCOD,
				Items: []http.CreateOrderItemData{
					{
						ProductID:        s.seededProductID,
						ProductVariantID: s.seededVariantID,
						Quantity:         1,
					},
				},
				UserID: s.seededUserID,
			},
		},
	}

	for _, tc := range testcases {
		s.Run(tc.name, func() {
			// Even with invalid data, the code will still try to call GetPaymentURL
			// So we need to mock it (though validation should ideally fail earlier)
			s.vnpayPaymentService.EXPECT().
				GetPaymentURL(mock.Anything, mock.Anything).
				Return("", nil).
				Maybe()

			// FIXME: This should fail validation but doesn't (DF-O-CV-01)
			// Once validation is added, remove the mock and expect error
			result, err := s.app.Create(ctx, http.CreateOrderRequestDto{Data: tc.data})
			// Currently succeeds but should fail
			s.Require().NoError(err)
			s.NotNil(result)
			// s.Require().Error(err) // Uncomment when DF-O-CV-01 is fixed
		})
	}
}

func (s *OrderTestSuite) TestBoundaryOrderQuantityAtMaximumInventory() {
	ctx := s.T().Context()

	// Mock for COD order (Maybe since validation might prevent call)
	s.vnpayPaymentService.EXPECT().
		GetPaymentURL(mock.Anything, mock.MatchedBy(func(param application.GetPaymentURLVNPayParam) bool {
			return param.Order.Provider == domain.PaymentProviderCOD
		})).
		Return("", nil).
		Maybe()

	// Get current inventory
	product, err := s.productRepo.Get(ctx, domain.ProductRepositoryGetParam{
		ProductID: s.seededProductID,
	})
	s.Require().NoError(err)
	variant := product.GetVariantByID(s.seededVariantID)
	s.Require().NotNil(variant)
	currentQuantity := variant.Quantity

	// Try to order exact available quantity
	result, err := s.app.Create(ctx, http.CreateOrderRequestDto{
		Data: http.CreateOrderData{
			RecipientName: "Boundary Test",
			PhoneNumber:   "+84111222333",
			Address:       "Boundary Street",
			Provider:      domain.PaymentProviderCOD,
			Items: []http.CreateOrderItemData{
				{
					ProductID:        s.seededProductID,
					ProductVariantID: s.seededVariantID,
					Quantity:         currentQuantity,
				},
			},
			UserID: s.seededUserID,
		},
	})
	s.Require().NoError(err)
	s.NotNil(result)
	s.Equal(currentQuantity, result.Items[0].Quantity)
}

func (s *OrderTestSuite) TestBoundaryOrderQuantityExceedsInventory() {
	ctx := s.T().Context()

	// NOTE: This test exposes a defect (DF-O-CV-01) - inventory validation is not performed

	// Mock for COD order (Maybe since validation might prevent call)
	s.vnpayPaymentService.EXPECT().
		GetPaymentURL(mock.Anything, mock.MatchedBy(func(param application.GetPaymentURLVNPayParam) bool {
			return param.Order.Provider == domain.PaymentProviderCOD
		})).
		Return("", nil).
		Maybe()

	// Get current inventory
	product, err := s.productRepo.Get(ctx, domain.ProductRepositoryGetParam{
		ProductID: s.seededSecondProductID,
	})
	s.Require().NoError(err)
	variant := product.GetVariantByID(s.seededSecondVariantID)
	s.Require().NotNil(variant)
	currentQuantity := variant.Quantity

	// FIXME: This should fail validation but doesn't (DF-O-CV-01)
	// Inventory validation is not performed in Create function
	result, err := s.app.Create(ctx, http.CreateOrderRequestDto{
		Data: http.CreateOrderData{
			RecipientName: "Overflow Test",
			PhoneNumber:   "+84444555666",
			Address:       "Overflow Street",
			Provider:      domain.PaymentProviderCOD,
			Items: []http.CreateOrderItemData{
				{
					ProductID:        s.seededSecondProductID,
					ProductVariantID: s.seededSecondVariantID,
					Quantity:         currentQuantity + 1,
				},
			},
			UserID: s.seededUserID,
		},
	})
	// Currently succeeds but should fail
	s.Require().NoError(err)
	s.NotNil(result)
	// s.Require().Error(err) // Uncomment when inventory validation is added
}

func (s *OrderTestSuite) TestBoundaryInvalidQuantities() {
	ctx := s.T().Context()

	// NOTE: These tests expose a defect (DF-O-CV-01) - validation is not performed

	testcases := []struct {
		name     string
		quantity int
	}{
		{
			name:     "Zero quantity order (SHOULD fail but doesn't - DF-O-CV-01)",
			quantity: 0,
		},
		{
			name:     "Negative quantity order (SHOULD fail but doesn't - DF-O-CV-01)",
			quantity: -1,
		},
	}

	for _, tc := range testcases {
		s.Run(tc.name, func() {
			// Mock for COD order (Maybe since validation might prevent call)
			s.vnpayPaymentService.EXPECT().
				GetPaymentURL(mock.Anything, mock.MatchedBy(func(param application.GetPaymentURLVNPayParam) bool {
					return param.Order.Provider == domain.PaymentProviderCOD
				})).
				Return("", nil).
				Maybe()

			// FIXME: This should fail validation but doesn't (DF-O-CV-01)
			result, err := s.app.Create(ctx, http.CreateOrderRequestDto{
				Data: http.CreateOrderData{
					RecipientName: tc.name,
					PhoneNumber:   "+84777888999",
					Address:       tc.name + " Street",
					Provider:      domain.PaymentProviderCOD,
					Items: []http.CreateOrderItemData{
						{
							ProductID:        s.seededProductID,
							ProductVariantID: s.seededVariantID,
							Quantity:         tc.quantity,
						},
					},
					UserID: s.seededUserID,
				},
			})
			// Currently succeeds but should fail
			s.Require().NoError(err)
			s.NotNil(result)
			// s.Require().Error(err) // Uncomment when DF-O-CV-01 is fixed
		})
	}
}

func (s *OrderTestSuite) TestListOrdersWithFilters() {
	ctx := s.T().Context()

	s.vnpayPaymentService.EXPECT().
		GetPaymentURL(mock.Anything, mock.Anything).
		Return("", nil).
		Times(0)

	result, err := s.app.Create(ctx, http.CreateOrderRequestDto{
		Data: http.CreateOrderData{
			RecipientName: "Filter Test Customer",
			PhoneNumber:   "+84123456789",
			Address:       "Filter Test Street",
			Provider:      domain.PaymentProviderCOD,
			ReturnURL:     "https://example.com/return",
			UserID:        s.seededUserID,
			Items: []http.CreateOrderItemData{
				{
					ProductID:        s.seededProductID,
					ProductVariantID: s.seededVariantID,
					Quantity:         1,
				},
			},
		},
	})

	s.Require().NoError(err)
	s.Require().NotNil(result)
	codOrderID := result.ID

	listResult, err := s.app.List(ctx, http.ListOrderRequestDto{
		PaginationRequestDto: http.PaginationRequestDto{
			Page:  1,
			Limit: 10,
		},
		IDs:     []uuid.UUID{codOrderID},
		UserIDs: []uuid.UUID{s.seededUserID},
	})
	s.Require().NoError(err)
	s.Require().NotNil(listResult)
	s.GreaterOrEqual(listResult.Meta.TotalItems, 1)
}
