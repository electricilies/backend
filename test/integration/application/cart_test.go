// vim: tabstop=4 shiftwidth=4:
//go:build integration

package application_test

import (
	"context"
	"strings"
	"testing"

	"backend/config"
	"backend/internal/application"
	"backend/internal/client"
	"backend/internal/delivery/http"
	"backend/internal/domain"
	"backend/internal/infrastructure/cacheredis"
	"backend/internal/infrastructure/repositorypostgres"
	"backend/internal/service"
	"backend/test/integration/component"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
)

type CartTestSuite struct {
	suite.Suite
	containers  *component.Containers
	app         http.CartApplication
	productRepo domain.ProductRepository

	// Seed data IDs from .rules/006-testing.md
	seededProductID       uuid.UUID
	seededVariantID       uuid.UUID
	seededUserID          uuid.UUID
	seededCartID          uuid.UUID
	seededCartItemID      uuid.UUID
	seededSecondVariantID uuid.UUID
}

func TestCartSuite(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(CartTestSuite))
}

func (s *CartTestSuite) newContainersConfig() *component.ContainersConfig {
	containersConfig := component.NewContainersConfig(&component.NewContainersConfigParam{
		DBEnabled:    true,
		RedisEnabled: true,
	})
	containersConfig.DB.Seed = true
	return containersConfig
}

func (s *CartTestSuite) newConfig(
	ctx context.Context,
) *config.Server {
	s.T().Helper()

	dbConnStr, err := s.containers.DB.ConnectionString(ctx, "sslmode=disable")
	s.Require().NoError(err, "failed to get db connection string")

	redisConnStr, err := s.containers.Redis.ConnectionString(ctx)
	s.Require().NoError(err, "failed to get redis connection string")
	return &config.Server{
		DBURL:     dbConnStr,
		RedisAddr: strings.TrimPrefix(redisConnStr, "redis://"),
	}
}

func (s *CartTestSuite) SetupSuite() {
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

	cartRepo := repositorypostgres.ProvideCart(queries, conn)
	s.productRepo = repositorypostgres.ProvideProduct(queries, conn)

	cartService := service.ProvideCart(validate)

	redisClient := client.NewRedis(ctx, cfg)
	cartCache := cacheredis.ProvideCart(redisClient)

	s.app = application.ProvideCart(cartRepo, cartService, cartCache, s.productRepo)

	s.seededProductID = uuid.MustParse("00000000-0000-7000-0000-000278469304")
	s.seededVariantID = uuid.MustParse("00000000-0000-7000-0000-000278469308")
	s.seededSecondVariantID = uuid.MustParse("00000000-0000-7000-0000-000278469306")
	s.seededUserID = uuid.MustParse("00000000-0000-7000-0000-000000000003") // customer user
	s.seededCartID = uuid.MustParse("00000000-0000-7000-0000-000000000001")
	s.seededCartItemID = uuid.MustParse("00000000-0000-7000-0000-000000000001")
}

func (s *CartTestSuite) TearDownSuite() {
	s.containers.Cleanup(s.T())
}

func (s *CartTestSuite) TestCartLifecycle() {
	ctx := s.T().Context()

	var newUserID uuid.UUID
	var newCartID uuid.UUID
	var newItemID uuid.UUID

	s.Run("Create Cart for new user", func() {
		newUserID = uuid.New()
		result, err := s.app.Create(ctx, http.CreateCartRequestDto{
			Data: http.CreateCartData{
				UserID: newUserID,
			},
		})
		s.Require().NoError(err)
		s.Require().NotNil(result)
		s.Equal(newUserID, result.UserID)
		s.Empty(result.Items)
		s.NotNil(result.ID)
		s.NotNil(result.UpdatedAt)
		newCartID = result.ID
	})

	s.Run("Get Cart by ID (cache miss)", func() {
		result, err := s.app.Get(ctx, http.GetCartRequestDto{
			CartID: newCartID,
		})
		s.Require().NoError(err)
		s.Require().NotNil(result)
		s.Equal(newCartID, result.ID)
		s.Equal(newUserID, result.UserID)
		s.Empty(result.Items)
	})

	s.Run("Get Cart by ID (cache hit)", func() {
		result1, err := s.app.Get(ctx, http.GetCartRequestDto{
			CartID: newCartID,
		})
		s.Require().NoError(err)
		s.Require().NotNil(result1)

		result2, err := s.app.Get(ctx, http.GetCartRequestDto{
			CartID: newCartID,
		})
		s.Require().NoError(err)
		s.Require().NotNil(result2)
		s.Equal(result1.ID, result2.ID)
		s.Equal(result1.UserID, result2.UserID)
	})

	s.Run("Get Cart by User ID", func() {
		result, err := s.app.GetByUser(ctx, http.GetCartByUserRequestDto{
			UserID: newUserID,
		})
		s.Require().NoError(err)
		s.Require().NotNil(result)
		s.Equal(newCartID, result.ID)
		s.Equal(newUserID, result.UserID)
	})

	s.Run("Add Item to Cart", func() {
		result, err := s.app.CreateItem(ctx, http.CreateCartItemRequestDto{
			UserID: newUserID,
			CartID: newCartID,
			Data: http.CreateCartItemData{
				ProductID:        s.seededProductID,
				ProductVariantID: s.seededVariantID,
				Quantity:         2,
			},
		})
		s.Require().NoError(err)
		s.Require().NotNil(result)
		s.Equal(2, result.Quantity)
		s.Equal(s.seededProductID, result.Product.ID)
		s.Equal(s.seededVariantID, result.ProductVariant.ID)
		s.NotEmpty(result.Product.Name)
		s.NotEmpty(result.ProductVariant.SKU)
		s.Positive(result.ProductVariant.Price)
		newItemID = result.ID
	})

	s.Run("Get Cart after adding item", func() {
		result, err := s.app.Get(ctx, http.GetCartRequestDto{
			CartID: newCartID,
		})
		s.Require().NoError(err)
		s.Require().NotNil(result)
		s.Len(result.Items, 1)
		s.Equal(newItemID, result.Items[0].ID)
		s.Equal(2, result.Items[0].Quantity)
		s.Equal(s.seededProductID, result.Items[0].Product.ID)
	})

	s.Run("Add same item again (upsert - quantity increases)", func() {
		result, err := s.app.CreateItem(ctx, http.CreateCartItemRequestDto{
			UserID: newUserID,
			CartID: newCartID,
			Data: http.CreateCartItemData{
				ProductID:        s.seededProductID,
				ProductVariantID: s.seededVariantID,
				Quantity:         3,
			},
		})
		s.Require().NoError(err)
		s.Require().NotNil(result)
		s.Equal(5, result.Quantity)
		s.Equal(newItemID, result.ID)
	})

	s.Run("Update Item quantity", func() {
		result, err := s.app.UpdateItem(ctx, http.UpdateCartItemRequestDto{
			UserID: newUserID,
			CartID: newCartID,
			ItemID: newItemID,
			Data: http.UpdateCartItemData{
				Quantity: 10,
			},
		})
		s.Require().NoError(err)
		s.Require().NotNil(result)
		s.Equal(10, result.Quantity)
		s.Equal(newItemID, result.ID)
	})

	s.Run("Cache is invalidated after update", func() {
		result, err := s.app.Get(ctx, http.GetCartRequestDto{
			CartID: newCartID,
		})
		s.Require().NoError(err)
		s.Require().NotNil(result)
		s.Len(result.Items, 1)
		s.Equal(10, result.Items[0].Quantity)
	})

	s.Run("Delete Item from Cart", func() {
		err := s.app.DeleteItem(ctx, http.DeleteCartItemRequestDto{
			UserID: newUserID,
			CartID: newCartID,
			ItemID: newItemID,
		})
		s.NoError(err)

		cart, err := s.app.Get(ctx, http.GetCartRequestDto{
			CartID: newCartID,
		})
		s.Require().NoError(err)
		s.Empty(cart.Items)
	})
}

func (s *CartTestSuite) TestCartSecurityCases() {
	ctx := s.T().Context()

	// Setup: Create a cart with an item
	newUserID := uuid.New()
	cartResult, err := s.app.Create(ctx, http.CreateCartRequestDto{
		Data: http.CreateCartData{
			UserID: newUserID,
		},
	})
	s.Require().NoError(err)
	newCartID := cartResult.ID

	itemResult, err := s.app.CreateItem(ctx, http.CreateCartItemRequestDto{
		UserID: newUserID,
		CartID: newCartID,
		Data: http.CreateCartItemData{
			ProductID:        s.seededProductID,
			ProductVariantID: s.seededVariantID,
			Quantity:         2,
		},
	})
	s.Require().NoError(err)
	newItemID := itemResult.ID

	s.Run("Security: Try to create item in another user's cart fails", func() {
		anotherUserID := uuid.New()
		_, err := s.app.CreateItem(ctx, http.CreateCartItemRequestDto{
			UserID: anotherUserID,
			CartID: newCartID,
			Data: http.CreateCartItemData{
				ProductID:        s.seededProductID,
				ProductVariantID: s.seededVariantID,
				Quantity:         1,
			},
		})
		s.Require().Error(err)
		s.ErrorIs(err, domain.ErrForbidden)
	})

	s.Run("Security: Try to update another user's cart item fails", func() {
		anotherUserID := uuid.New()
		_, err := s.app.UpdateItem(ctx, http.UpdateCartItemRequestDto{
			UserID: anotherUserID,
			CartID: newCartID,
			ItemID: newItemID,
			Data: http.UpdateCartItemData{
				Quantity: 100,
			},
		})
		s.Require().Error(err)
		s.ErrorIs(err, domain.ErrForbidden)
	})

	s.Run("Security: Try to delete from another user's cart fails", func() {
		anotherUserID := uuid.New()
		err := s.app.DeleteItem(ctx, http.DeleteCartItemRequestDto{
			UserID: anotherUserID,
			CartID: newCartID,
			ItemID: newItemID,
		})
		s.Require().Error(err)
		s.ErrorIs(err, domain.ErrForbidden)
	})
}

func (s *CartTestSuite) TestCartErrorCases() {
	ctx := s.T().Context()
	nonExistentID := uuid.MustParse("00000000-0000-0000-0000-000000000001")
	newUserID := uuid.New()

	// Setup: Create a valid cart
	cartResult, err := s.app.Create(ctx, http.CreateCartRequestDto{
		Data: http.CreateCartData{
			UserID: newUserID,
		},
	})
	s.Require().NoError(err)
	newCartID := cartResult.ID

	s.Run("Get non-existent cart fails", func() {
		_, err := s.app.Get(ctx, http.GetCartRequestDto{
			CartID: nonExistentID,
		})
		s.Require().Error(err)
	})

	s.Run("Add item to non-existent cart fails", func() {
		_, err := s.app.CreateItem(ctx, http.CreateCartItemRequestDto{
			UserID: newUserID,
			CartID: nonExistentID,
			Data: http.CreateCartItemData{
				ProductID:        s.seededProductID,
				ProductVariantID: s.seededVariantID,
				Quantity:         1,
			},
		})
		s.Require().Error(err)
	})

	s.Run("Update non-existent item fails", func() {
		_, err := s.app.UpdateItem(ctx, http.UpdateCartItemRequestDto{
			UserID: newUserID,
			CartID: newCartID,
			ItemID: nonExistentID,
			Data: http.UpdateCartItemData{
				Quantity: 5,
			},
		})
		s.Require().Error(err)
	})

	s.Run("Add item with invalid product fails", func() {
		_, err := s.app.CreateItem(ctx, http.CreateCartItemRequestDto{
			UserID: newUserID,
			CartID: newCartID,
			Data: http.CreateCartItemData{
				ProductID:        nonExistentID,
				ProductVariantID: nonExistentID,
				Quantity:         1,
			},
		})
		s.Require().Error(err)
	})
}

func (s *CartTestSuite) TestCartValidation() {
	ctx := s.T().Context()
	newUserID := uuid.New()

	// Setup: Create a cart with an item
	cartResult, err := s.app.Create(ctx, http.CreateCartRequestDto{
		Data: http.CreateCartData{
			UserID: newUserID,
		},
	})
	s.Require().NoError(err)
	newCartID := cartResult.ID

	itemResult, err := s.app.CreateItem(ctx, http.CreateCartItemRequestDto{
		UserID: newUserID,
		CartID: newCartID,
		Data: http.CreateCartItemData{
			ProductID:        s.seededProductID,
			ProductVariantID: s.seededVariantID,
			Quantity:         1,
		},
	})
	s.Require().NoError(err)
	newItemID := itemResult.ID

	s.Run("Validation: Quantity must be positive", func() {
		_, err := s.app.UpdateItem(ctx, http.UpdateCartItemRequestDto{
			UserID: newUserID,
			CartID: newCartID,
			ItemID: newItemID,
			Data: http.UpdateCartItemData{
				Quantity: -1,
			},
		})
		s.Require().Error(err)
	})

	s.Run("Validation: Quantity exceeds max (100)", func() {
		_, err := s.app.CreateItem(ctx, http.CreateCartItemRequestDto{
			UserID: newUserID,
			CartID: newCartID,
			Data: http.CreateCartItemData{
				ProductID:        s.seededProductID,
				ProductVariantID: s.seededVariantID,
				Quantity:         101,
			},
		})
		s.Require().Error(err)
	})
}

func (s *CartTestSuite) TestCartWithMultipleItems() {
	ctx := s.T().Context()
	newUserID := uuid.New()

	// Setup: Create a cart
	cartResult, err := s.app.Create(ctx, http.CreateCartRequestDto{
		Data: http.CreateCartData{
			UserID: newUserID,
		},
	})
	s.Require().NoError(err)
	newCartID := cartResult.ID

	s.Run("Add second different variant to test multiple items", func() {
		_, err := s.app.CreateItem(ctx, http.CreateCartItemRequestDto{
			UserID: newUserID,
			CartID: newCartID,
			Data: http.CreateCartItemData{
				ProductID:        s.seededProductID,
				ProductVariantID: s.seededVariantID,
				Quantity:         3,
			},
		})
		s.Require().NoError(err)

		_, err = s.app.CreateItem(ctx, http.CreateCartItemRequestDto{
			UserID: newUserID,
			CartID: newCartID,
			Data: http.CreateCartItemData{
				ProductID:        s.seededProductID,
				ProductVariantID: s.seededSecondVariantID,
				Quantity:         2,
			},
		})
		s.Require().NoError(err)

		cart, err := s.app.Get(ctx, http.GetCartRequestDto{
			CartID: newCartID,
		})
		s.Require().NoError(err)
		s.Len(cart.Items, 2)
	})
}

func (s *CartTestSuite) TestSeededCartAccess() {
	ctx := s.T().Context()

	s.Run("Test with seeded cart and items", func() {
		result, err := s.app.Get(ctx, http.GetCartRequestDto{
			CartID: s.seededCartID,
		})
		s.Require().NoError(err)
		s.Require().NotNil(result)
		s.Equal(s.seededUserID, result.UserID)
		s.NotEmpty(result.Items)

		for _, item := range result.Items {
			s.NotNil(item.Product.ID)
			s.NotEmpty(item.Product.Name)
			s.NotNil(item.ProductVariant.ID)
			s.NotEmpty(item.ProductVariant.SKU)
			s.Positive(item.Quantity)
		}
	})
}
