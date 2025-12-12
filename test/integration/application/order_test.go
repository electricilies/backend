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

	// Seed data IDs from .rules/011-integrationtest.md

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
	err = domain.RegisterOrderValidates(validate)
	s.Require().NoError(err)

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

	// Seed data from .rules/011-integrationtest.md

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

func (s *OrderTestSuite) TestCODOrderLifecycle() {
	ctx := s.T().Context()
	var codOrderID uuid.UUID

	s.Run("Create COD order", func() {
		result, err := s.app.Create(ctx, http.CreateOrderRequestDto{
			Data: http.CreateOrderData{
				RecipientName: "Cash Customer",
				PhoneNumber:   "+84123456789",
				Address:       "123 Cash Street, Ho Chi Minh City",
				Provider:      domain.PaymentProviderCOD,
				ReturnURL:     "https://example.com/return",
				UserID:        s.seededUserID,
				Items: []http.CreateOrderItemData{
					{
						ProductID:        s.seededProductID,
						ProductVariantID: s.seededVariantID,
						Quantity:         2,
					},
				},
			},
		})
		s.Require().NoError(err)
		s.Require().NotNil(result)
		s.Equal(domain.OrderStatusPending, result.Status)
		s.False(result.IsPaid)
		s.Empty(result.PaymentURL, "COD orders should not have payment URL")
		s.Equal("Cash Customer", result.RecipentName)
		s.Equal("+84123456789", result.PhoneNumber)
		s.Len(result.Items, 1)
		s.Equal(2, result.Items[0].Quantity)
		s.Positive(result.TotalAmount)
		codOrderID = result.ID
	})

	s.Run("Get COD order", func() {
		result, err := s.app.Get(ctx, http.GetOrderRequestDto{
			OrderID: codOrderID,
		})
		s.Require().NoError(err)
		s.Require().NotNil(result)
		s.Equal(codOrderID, result.ID)
		s.Equal(domain.OrderStatusPending, result.Status)
		s.False(result.IsPaid)
		s.Len(result.Items, 1)

		item := result.Items[0]
		s.NotNil(item)
		s.NotEmpty(item.Product.Name)
		s.NotEmpty(item.ProductVariant.SKU)
		s.Positive(item.Price)
	})

	s.Run("Update COD order address", func() {
		result, err := s.app.Update(ctx, http.UpdateOrderRequestDto{
			OrderID: codOrderID,
			Data: http.UpdateOrderData{
				Address: "456 New Cash Street, Hanoi",
				Status:  domain.OrderStatusShipping,
				IsPaid:  false,
			},
		})
		s.Require().NoError(err)
		s.Require().NotNil(result)
		s.Equal("456 New Cash Street, Hanoi", result.Address)
		s.Equal(domain.OrderStatusShipping, result.Status)
	})

	s.Run("List orders with filters", func() {
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

		found := false
		for _, order := range listResult.Data {
			if order.ID == codOrderID {
				found = true
				s.Equal(domain.PaymentProviderCOD, order.Provider)
				break
			}
		}
		s.True(found, "COD order should be in the list")
	})
}

func (s *OrderTestSuite) TestVNPayOrderLifecycle() {
	ctx := s.T().Context()
	var vnpayOrderID uuid.UUID

	s.Run("Create VNPAY order with single item", func() {
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
		s.NotEmpty(result.PaymentURL, "VNPAY orders must have payment URL")
		s.Equal("John Doe", result.RecipentName)
		s.Len(result.Items, 1)
		s.Positive(result.TotalAmount)
		vnpayOrderID = result.ID
	})

	s.Run("Get VNPAY order", func() {
		result, err := s.app.Get(ctx, http.GetOrderRequestDto{
			OrderID: vnpayOrderID,
		})
		s.Require().NoError(err)
		s.Require().NotNil(result)
		s.Equal(vnpayOrderID, result.ID)
		s.Equal(domain.OrderStatusPending, result.Status)
		s.Equal(domain.PaymentProviderVNPAY, result.Provider)
	})

	s.Run("Update VNPAY order", func() {
		result, err := s.app.Update(ctx, http.UpdateOrderRequestDto{
			OrderID: vnpayOrderID,
			Data: http.UpdateOrderData{
				Address: "456 Updated Address, Hanoi",
				Status:  domain.OrderStatusPending,
				IsPaid:  false,
			},
		})
		s.Require().NoError(err)
		s.Require().NotNil(result)
		s.Equal("456 Updated Address, Hanoi", result.Address)
	})

	s.Run("VNPAY IPN Success - Verify payment and inventory decrease", func() {
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
				Amount:            "100000000",
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
}

func (s *OrderTestSuite) TestVNPayOrderWithMultipleItems() {
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

	expectedTotal := result.Items[0].Price*int64(result.Items[0].Quantity) +
		result.Items[1].Price*int64(result.Items[1].Quantity)
	s.Equal(expectedTotal, result.TotalAmount)
}

func (s *OrderTestSuite) TestVNPayIPNFailure() {
	ctx := s.T().Context()
	var vnpayOrderID uuid.UUID

	s.Run("Create VNPAY order for failure test", func() {
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
		vnpayOrderID = result.ID
	})

	s.Run("VNPay IPN Failure - Order cancelled, inventory unchanged", func() {
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
				ResponseCode:      "24",
				SecureHash:        "test_hash",
				TmnCode:           "test_tmn",
				TransactionNo:     "987654321",
				TransactionStatus: "02",
				TxnRef:            vnpayOrderID.String(),
			},
		})
		s.Require().NoError(err)
		s.Require().NotNil(result)
		s.NotEqual("00", result.RspCode)

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
}

func (s *OrderTestSuite) TestGetNonExistentOrder() {
	ctx := s.T().Context()
	nonExistentID := uuid.MustParse("00000000-0000-0000-0000-000000000001")

	_, err := s.app.Get(ctx, http.GetOrderRequestDto{
		OrderID: nonExistentID,
	})
	s.Error(err)
}

func (s *OrderTestSuite) TestUpdateNonExistentOrder() {
	ctx := s.T().Context()
	nonExistentID := uuid.MustParse("00000000-0000-0000-0000-000000000001")

	_, err := s.app.Update(ctx, http.UpdateOrderRequestDto{
		OrderID: nonExistentID,
		Data: http.UpdateOrderData{
			Address: "New Address",
			Status:  domain.OrderStatusPending,
			IsPaid:  false,
		},
	})
	s.Error(err)
}

func (s *OrderTestSuite) TestCreateOrderWithNonExistentProduct() {
	ctx := s.T().Context()
	nonExistentID := uuid.MustParse("00000000-0000-0000-0000-000000000001")

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
	s.Error(err)
}

func (s *OrderTestSuite) TestOrderValidationDefectEmptyItems() {
	ctx := s.T().Context()

	data := http.CreateOrderData{
		RecipientName: "Test",
		PhoneNumber:   "+84912345678",
		Address:       "Test",
		Provider:      domain.PaymentProviderCOD,
		Items:         []http.CreateOrderItemData{},
		UserID:        s.seededUserID,
	}

	result, err := s.app.Create(ctx, http.CreateOrderRequestDto{Data: data})

	s.Require().Error(err)
	s.Nil(result)
}

func (s *OrderTestSuite) TestOrderValidationDefectInvalidPhoneNumber() {
	ctx := s.T().Context()

	data := http.CreateOrderData{
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
	}

	result, err := s.app.Create(ctx, http.CreateOrderRequestDto{Data: data})

	s.Require().Error(err)
	s.Nil(result)
}
