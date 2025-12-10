package domain_test

import (
	"testing"

	"backend/internal/domain"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
)

type CartTestSuite struct {
	suite.Suite
	validate *validator.Validate
}

func (s *CartTestSuite) SetupSuite() {
	s.validate = validator.New(validator.WithRequiredStructEnabled())
}

func (s *CartTestSuite) TestNewCart() {
	s.T().Parallel()
	testcases := []struct {
		name      string
		userID    uuid.UUID
		expectOk  bool
		expectErr bool
	}{
		{
			name:      "valid user ID",
			userID:    uuid.New(),
			expectOk:  true,
			expectErr: false,
		},
		{
			name:      "nil user ID",
			userID:    uuid.Nil,
			expectOk:  false,
			expectErr: true,
		},
	}

	for _, tc := range testcases {
		s.Run(tc.name, func() {
			cart, err := domain.NewCart(tc.userID)

			s.NoError(err, tc.name)
			s.NotNil(cart, tc.name)
			s.Equal(tc.userID, cart.UserID, tc.name)
			s.NotNil(cart.ID, tc.name)
			s.NotNil(cart.Items, tc.name)
			s.Empty(cart.Items, tc.name)

			validationErr := s.validate.Struct(cart)
			if tc.expectErr {
				s.Error(validationErr, tc.name)
			} else {
				s.NoError(validationErr, tc.name)
			}
		})
	}
}

func (s *CartTestSuite) TestNewCartItemBoundaryValues() {
	s.T().Parallel()
	testcases := []struct {
		name      string
		quantity  int
		expectOk  bool
		expectErr bool
	}{
		{
			name:      "quantity 0 (invalid)",
			quantity:  0,
			expectOk:  false,
			expectErr: true,
		},
		{
			name:      "quantity 1 (min)",
			quantity:  1,
			expectOk:  true,
			expectErr: false,
		},
		{
			name:      "quantity 2 (min + 1)",
			quantity:  2,
			expectOk:  true,
			expectErr: false,
		},
		{
			name:      "quantity 50",
			quantity:  50,
			expectOk:  true,
			expectErr: false,
		},
		{
			name:      "quantity 99 (max - 1)",
			quantity:  99,
			expectOk:  true,
			expectErr: false,
		},
		{
			name:      "quantity 100 (max)",
			quantity:  100,
			expectOk:  true,
			expectErr: false,
		},
		{
			name:      "quantity 101 (max + 1)",
			quantity:  101,
			expectOk:  false,
			expectErr: true,
		},
		{
			name:      "quantity -1 (negative)",
			quantity:  -1,
			expectOk:  false,
			expectErr: true,
		},
	}

	productID := uuid.New()
	variantID := uuid.New()

	for _, tc := range testcases {
		s.Run(tc.name, func() {
			cartItem, err := domain.NewCartItem(productID, variantID, tc.quantity)

			s.NoError(err, tc.name)
			s.NotNil(cartItem, tc.name)
			s.Equal(tc.quantity, cartItem.Quantity, tc.name)
			s.Equal(productID, cartItem.ProductID, tc.name)
			s.Equal(variantID, cartItem.ProductVariantID, tc.name)
			s.NotNil(cartItem.ID, tc.name)

			validationErr := s.validate.Struct(cartItem)
			if tc.expectErr {
				s.Error(validationErr, tc.name)
			} else {
				s.NoError(validationErr, tc.name)
			}
		})
	}
}

func (s *CartTestSuite) TestCartUpsertItem() {
	s.T().Parallel()
	testcases := []struct {
		name          string
		existingItems []struct {
			productID uuid.UUID
			quantity  int
		}
		newItem struct {
			productID uuid.UUID
			quantity  int
		}
		expectedLength   int
		expectedQuantity int
		isUpdate         bool
	}{
		{
			name: "add new item to empty cart",
			existingItems: []struct {
				productID uuid.UUID
				quantity  int
			}{},
			newItem: struct {
				productID uuid.UUID
				quantity  int
			}{productID: uuid.New(), quantity: 5},
			expectedLength:   1,
			expectedQuantity: 5,
			isUpdate:         false,
		},
		{
			name: "update existing item quantity",
			existingItems: []struct {
				productID uuid.UUID
				quantity  int
			}{
				{productID: uuid.MustParse("00000000-0000-0000-0000-000000000001"), quantity: 3},
			},
			newItem: struct {
				productID uuid.UUID
				quantity  int
			}{productID: uuid.MustParse("00000000-0000-0000-0000-000000000001"), quantity: 2},
			expectedLength:   1,
			expectedQuantity: 5,
			isUpdate:         true,
		},
		{
			name: "add different item to cart with existing items",
			existingItems: []struct {
				productID uuid.UUID
				quantity  int
			}{
				{productID: uuid.MustParse("00000000-0000-0000-0000-000000000001"), quantity: 3},
			},
			newItem: struct {
				productID uuid.UUID
				quantity  int
			}{productID: uuid.MustParse("00000000-0000-0000-0000-000000000002"), quantity: 2},
			expectedLength:   2,
			expectedQuantity: 2,
			isUpdate:         false,
		},
	}

	for _, tc := range testcases {
		s.Run(tc.name, func() {
			cart, err := domain.NewCart(uuid.New())
			s.Require().NoError(err)

			for _, existing := range tc.existingItems {
				item, err := domain.NewCartItem(existing.productID, uuid.New(), existing.quantity)
				s.Require().NoError(err)
				cart.UpsertItem(*item)
			}

			newItem, err := domain.NewCartItem(tc.newItem.productID, uuid.New(), tc.newItem.quantity)
			s.Require().NoError(err)

			result := cart.UpsertItem(*newItem)

			s.Len(cart.Items, tc.expectedLength, tc.name)
			s.Equal(tc.expectedQuantity, result.Quantity, tc.name)
		})
	}
}

func (s *CartTestSuite) TestCartUpdateItem() {
	s.T().Parallel()
	testcases := []struct {
		name             string
		setupQuantity    int
		updateQuantity   int
		expectedLength   int
		expectedQuantity int
		itemExists       bool
	}{
		{
			name:             "update quantity to positive value",
			setupQuantity:    5,
			updateQuantity:   10,
			expectedLength:   1,
			expectedQuantity: 10,
			itemExists:       true,
		},
		{
			name:             "update quantity to zero (removes item)",
			setupQuantity:    5,
			updateQuantity:   0,
			expectedLength:   0,
			expectedQuantity: 0,
			itemExists:       false,
		},
		{
			name:             "update quantity to 1",
			setupQuantity:    10,
			updateQuantity:   1,
			expectedLength:   1,
			expectedQuantity: 1,
			itemExists:       true,
		},
	}

	for _, tc := range testcases {
		s.Run(tc.name, func() {
			cart, err := domain.NewCart(uuid.New())
			s.Require().NoError(err)

			item, err := domain.NewCartItem(uuid.New(), uuid.New(), tc.setupQuantity)
			s.Require().NoError(err)

			cart.UpsertItem(*item)
			itemID := item.ID

			cart.UpdateItem(itemID, tc.updateQuantity)

			s.Len(cart.Items, tc.expectedLength, tc.name)

			if tc.itemExists {
				found := false
				for _, cartItem := range cart.Items {
					if cartItem.ID == itemID {
						s.Equal(tc.expectedQuantity, cartItem.Quantity, tc.name)
						found = true
						break
					}
				}
				s.True(found, "item should exist in cart")
			}
		})
	}
}

func (s *CartTestSuite) TestCartRemoveItem() {
	s.T().Parallel()
	testcases := []struct {
		name           string
		setupItems     int
		removeIndex    int
		expectedLength int
		targetNotFound bool
	}{
		{
			name:           "remove existing item",
			setupItems:     3,
			removeIndex:    1,
			expectedLength: 2,
			targetNotFound: false,
		},
		{
			name:           "remove non-existent item",
			setupItems:     2,
			removeIndex:    -1,
			expectedLength: 2,
			targetNotFound: true,
		},
		{
			name:           "remove from single item cart",
			setupItems:     1,
			removeIndex:    0,
			expectedLength: 0,
			targetNotFound: false,
		},
	}

	for _, tc := range testcases {
		s.Run(tc.name, func() {
			cart, err := domain.NewCart(uuid.New())
			s.Require().NoError(err)

			var itemIDs []uuid.UUID
			for i := 0; i < tc.setupItems; i++ {
				item, err := domain.NewCartItem(uuid.New(), uuid.New(), 1)
				s.Require().NoError(err)
				cart.UpsertItem(*item)
				itemIDs = append(itemIDs, item.ID)
			}

			var targetID uuid.UUID
			if tc.targetNotFound {
				targetID = uuid.New()
			} else {
				targetID = itemIDs[tc.removeIndex]
			}

			cart.RemoveItem(targetID)

			s.Len(cart.Items, tc.expectedLength, tc.name)
		})
	}
}

func (s *CartTestSuite) TestCartClearItems() {
	s.T().Parallel()
	s.Run("clear items from cart with items", func() {
		cart, err := domain.NewCart(uuid.New())
		s.Require().NoError(err)

		for range 5 {
			item, err := domain.NewCartItem(uuid.New(), uuid.New(), 1)
			s.Require().NoError(err)
			cart.UpsertItem(*item)
		}

		s.Len(cart.Items, 5)

		cart.ClearItems()

		s.Empty(cart.Items)
		s.NotNil(cart.Items)
	})

	s.Run("clear items from empty cart", func() {
		cart, err := domain.NewCart(uuid.New())
		s.Require().NoError(err)

		s.Empty(cart.Items)

		cart.ClearItems()

		s.Empty(cart.Items)
		s.NotNil(cart.Items)
	})
}

func TestCart(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(CartTestSuite))
}
