// vim: tabstop=4 shiftwidth=4:
package domain_test

import (
	"strings"
	"testing"
	"time"

	"backend/internal/domain"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
)

type ProductTestSuite struct {
	suite.Suite
	validate *validator.Validate
}

func (s *ProductTestSuite) SetupSuite() {
	s.validate = validator.New(validator.WithRequiredStructEnabled())
	err := domain.RegisterProductValidates(s.validate)
	s.Require().NoError(err, "Failed to register product validators")
}

func (s *ProductTestSuite) TestNewProductBoundaryValues() {
	categoryID := uuid.New()
	testcases := []struct {
		name        string
		prodName    string
		description string
		categoryID  uuid.UUID
		expectErr   bool
	}{
		{
			name:        "name length 2 (min - 1)",
			prodName:    "ab",
			description: "Valid description here",
			categoryID:  categoryID,
			expectErr:   true,
		},
		{
			name:        "name length 3 (min)",
			prodName:    "abc",
			description: "Valid description here",
			categoryID:  categoryID,
			expectErr:   false,
		},
		{
			name:        "name length 4 (min + 1)",
			prodName:    "abcd",
			description: "Valid description here",
			categoryID:  categoryID,
			expectErr:   false,
		},
		{
			name:        "name length 200 (max)",
			prodName:    strings.Repeat("a", 200),
			description: "Valid description here",
			categoryID:  categoryID,
			expectErr:   false,
		},
		{
			name:        "name length 201 (max + 1)",
			prodName:    strings.Repeat("a", 201),
			description: "Valid description here",
			categoryID:  categoryID,
			expectErr:   true,
		},
		{
			name:        "description length 9 (min - 1)",
			prodName:    "Valid Product Name",
			description: "123456789",
			categoryID:  categoryID,
			expectErr:   true,
		},
		{
			name:        "description length 10 (min)",
			prodName:    "Valid Product Name",
			description: "1234567890",
			categoryID:  categoryID,
			expectErr:   false,
		},
		{
			name:        "description length 11 (min + 1)",
			prodName:    "Valid Product Name",
			description: "12345678901",
			categoryID:  categoryID,
			expectErr:   false,
		},
		{
			name:        "valid product",
			prodName:    "T-Shirt",
			description: "A comfortable cotton t-shirt",
			categoryID:  categoryID,
			expectErr:   false,
		},
		{
			name:        "empty name",
			prodName:    "",
			description: "Valid description here",
			categoryID:  categoryID,
			expectErr:   true,
		},
		{
			name:        "empty description",
			prodName:    "Valid Product Name",
			description: "",
			categoryID:  categoryID,
			expectErr:   true,
		},
		{
			name:        "nil category ID",
			prodName:    "Valid Product Name",
			description: "Valid description here",
			categoryID:  uuid.Nil,
			expectErr:   true,
		},
	}

	for _, tc := range testcases {
		s.Run(tc.name, func() {
			product, err := domain.NewProduct(tc.prodName, tc.description, tc.categoryID)

			s.NoError(err, tc.name)
			s.NotNil(product, tc.name)
			s.Equal(tc.prodName, product.Name, tc.name)
			s.Equal(tc.description, product.Description, tc.name)
			s.Equal(tc.categoryID, product.CategoryID, tc.name)
			s.NotEqual(uuid.Nil, product.ID, tc.name)
			s.NotZero(product.CreatedAt, tc.name)
			s.NotZero(product.UpdatedAt, tc.name)
			s.True(product.CreatedAt.Equal(product.UpdatedAt), tc.name)

			if tc.expectErr {
				if tc.prodName == "" || len(tc.prodName) < 3 || len(tc.prodName) > 200 {
					err := s.validate.Var(product.Name, "required,gte=3,lte=200")
					s.Error(err, tc.name)
				}
				if tc.description == "" || len(tc.description) < 10 {
					err := s.validate.Var(product.Description, "required,gte=10")
					s.Error(err, tc.name)
				}
				if tc.categoryID == uuid.Nil {
					err := s.validate.Var(product.CategoryID, "required")
					s.Error(err, tc.name)
				}
			} else {
				err := s.validate.Var(product.Name, "required,gte=3,lte=200")
				s.NoError(err, tc.name)
				err = s.validate.Var(product.Description, "required,gte=10")
				s.NoError(err, tc.name)
				err = s.validate.Var(product.CategoryID, "required")
				s.NoError(err, tc.name)
			}
		})
	}
}

func (s *ProductTestSuite) TestNewProductOptionBoundaryValues() {
	testcases := []struct {
		name      string
		optName   string
		expectErr bool
	}{
		{
			name:      "empty option name",
			optName:   "",
			expectErr: true,
		},
		{
			name:      "valid option name",
			optName:   "Size",
			expectErr: false,
		},
		{
			name:      "single character option name",
			optName:   "S",
			expectErr: false,
		},
	}

	for _, tc := range testcases {
		s.Run(tc.name, func() {
			option, err := domain.NewProductOption(tc.optName)

			s.NoError(err, tc.name)
			s.NotNil(option, tc.name)
			s.Equal(tc.optName, option.Name, tc.name)
			s.NotEqual(uuid.Nil, option.ID, tc.name)

			validationErr := s.validate.Struct(option)
			if tc.expectErr {
				s.Error(validationErr, tc.name)
			} else {
				s.NoError(validationErr, tc.name)
			}
		})
	}
}

func (s *ProductTestSuite) TestNewProductImageBoundaryValues() {
	buildURL := func(id uuid.UUID) string {
		return "https://example.com/images/" + id.String()
	}

	testcases := []struct {
		name      string
		order     int
		expectErr bool
	}{
		{
			name:      "order -1 (negative)",
			order:     -1,
			expectErr: true,
		},
		{
			name:      "order 1 (min + 1)",
			order:     1,
			expectErr: false,
		},
		{
			name:      "order 100 (large value)",
			order:     100,
			expectErr: false,
		},
	}

	for _, tc := range testcases {
		s.Run(tc.name, func() {
			image, err := domain.NewProductImage(tc.order, buildURL)

			s.NoError(err, tc.name)
			s.NotNil(image, tc.name)
			s.Equal(tc.order, image.Order, tc.name)
			s.NotEqual(uuid.Nil, image.ID, tc.name)
			s.Contains(image.URL, "https://example.com/images/", tc.name)
			s.NotZero(image.CreatedAt, tc.name)

			validationErr := s.validate.Struct(image)
			if tc.expectErr {
				s.Error(validationErr, tc.name)
			} else {
				s.NoError(validationErr, tc.name)
			}
		})
	}
}

func (s *ProductTestSuite) TestNewVariantBoundaryValues() {
	testcases := []struct {
		name      string
		sku       string
		price     int64
		quantity  int
		expectErr bool
	}{
		{
			name:      "price 0 (invalid)",
			sku:       "SKU-001",
			price:     0,
			quantity:  10,
			expectErr: true,
		},
		{
			name:      "price 1 (min)",
			sku:       "SKU-001",
			price:     1,
			quantity:  10,
			expectErr: false,
		},
		{
			name:      "price negative",
			sku:       "SKU-001",
			price:     -100,
			quantity:  10,
			expectErr: true,
		},
		{
			name:      "quantity -1 (negative)",
			sku:       "SKU-001",
			price:     1000,
			quantity:  -1,
			expectErr: true,
		},
		{
			name:      "quantity 0 (min)",
			sku:       "SKU-001",
			price:     1000,
			quantity:  0,
			expectErr: false,
		},
		{
			name:      "quantity 1 (min + 1)",
			sku:       "SKU-001",
			price:     1000,
			quantity:  1,
			expectErr: false,
		},
		{
			name:      "empty SKU",
			sku:       "",
			price:     1000,
			quantity:  10,
			expectErr: true,
		},
		{
			name:      "valid variant",
			sku:       "SKU-TSHIRT-S-RED",
			price:     29900,
			quantity:  50,
			expectErr: false,
		},
	}

	for _, tc := range testcases {
		s.Run(tc.name, func() {
			variant, err := domain.NewVariant(tc.sku, tc.price, tc.quantity)

			s.NoError(err, tc.name)
			s.NotNil(variant, tc.name)
			s.Equal(tc.sku, variant.SKU, tc.name)
			s.Equal(tc.price, variant.Price, tc.name)
			s.Equal(tc.quantity, variant.Quantity, tc.name)
			s.Equal(0, variant.PurchaseCount, tc.name)
			s.NotEqual(uuid.Nil, variant.ID, tc.name)
			s.NotZero(variant.CreatedAt, tc.name)
			s.NotZero(variant.UpdatedAt, tc.name)

			validationErr := s.validate.Struct(variant)
			if tc.expectErr {
				s.Error(validationErr, tc.name)
			} else {
				s.NoError(validationErr, tc.name)
			}
		})
	}
}

func (s *ProductTestSuite) TestCreateOptionValues() {
	testcases := []struct {
		name          string
		values        []string
		expectedCount int
		expectErr     bool
	}{
		{
			name:          "empty values",
			values:        []string{},
			expectedCount: 0,
			expectErr:     false,
		},
		{
			name:          "single value",
			values:        []string{"Red"},
			expectedCount: 1,
			expectErr:     false,
		},
		{
			name:          "multiple values",
			values:        []string{"Red", "Blue", "Green"},
			expectedCount: 3,
			expectErr:     false,
		},
		{
			name:          "values with empty string",
			values:        []string{"Red", "", "Blue"},
			expectedCount: 3,
			expectErr:     false,
		},
	}

	for _, tc := range testcases {
		s.Run(tc.name, func() {
			optionValues, err := domain.CreateOptionValues(tc.values)

			if tc.expectErr {
				s.Error(err, tc.name)
			} else {
				s.NoError(err, tc.name)
				s.Len(optionValues, tc.expectedCount, tc.name)
				for i, ov := range optionValues {
					s.NotEqual(uuid.Nil, ov.ID, tc.name)
					s.Equal(tc.values[i], ov.Value, tc.name)
				}
			}
		})
	}
}

func (s *ProductTestSuite) TestProductUpdate() {
	testcases := []struct {
		name               string
		initialName        string
		initialDescription string
		initialCategoryID  uuid.UUID
		updateName         string
		updateDescription  string
		updateCategoryID   uuid.UUID
		expectedName       string
		expectedDesc       string
		expectedCatID      uuid.UUID
		shouldUpdateTime   bool
	}{
		{
			name:               "update name only",
			initialName:        "Original Name",
			initialDescription: "Original Description",
			initialCategoryID:  uuid.New(),
			updateName:         "Updated Name",
			updateDescription:  "",
			updateCategoryID:   uuid.Nil,
			expectedName:       "Updated Name",
			expectedDesc:       "Original Description",
			expectedCatID:      uuid.Nil,
			shouldUpdateTime:   true,
		},
		{
			name:               "update description only",
			initialName:        "Original Name",
			initialDescription: "Original Description",
			initialCategoryID:  uuid.New(),
			updateName:         "",
			updateDescription:  "Updated Description",
			updateCategoryID:   uuid.Nil,
			expectedName:       "Original Name",
			expectedDesc:       "Updated Description",
			expectedCatID:      uuid.Nil,
			shouldUpdateTime:   true,
		},
		{
			name:               "update category ID only",
			initialName:        "Original Name",
			initialDescription: "Original Description",
			initialCategoryID:  uuid.New(),
			updateName:         "",
			updateDescription:  "",
			updateCategoryID:   uuid.New(),
			expectedName:       "Original Name",
			expectedDesc:       "Original Description",
			expectedCatID:      uuid.Nil,
			shouldUpdateTime:   true,
		},
		{
			name:               "update all fields",
			initialName:        "Original Name",
			initialDescription: "Original Description",
			initialCategoryID:  uuid.New(),
			updateName:         "Updated Name",
			updateDescription:  "Updated Description",
			updateCategoryID:   uuid.New(),
			expectedName:       "Updated Name",
			expectedDesc:       "Updated Description",
			expectedCatID:      uuid.Nil,
			shouldUpdateTime:   true,
		},
		{
			name:               "update with empty values (no change)",
			initialName:        "Original Name",
			initialDescription: "Original Description",
			initialCategoryID:  uuid.New(),
			updateName:         "",
			updateDescription:  "",
			updateCategoryID:   uuid.Nil,
			expectedName:       "Original Name",
			expectedDesc:       "Original Description",
			expectedCatID:      uuid.Nil,
			shouldUpdateTime:   false,
		},
		{
			name:               "update with same values (no change)",
			initialName:        "Same Name",
			initialDescription: "Same Description",
			initialCategoryID:  uuid.New(),
			updateName:         "Same Name",
			updateDescription:  "Same Description",
			updateCategoryID:   uuid.Nil,
			expectedName:       "Same Name",
			expectedDesc:       "Same Description",
			expectedCatID:      uuid.Nil,
			shouldUpdateTime:   false,
		},
	}

	for _, tc := range testcases {
		s.Run(tc.name, func() {
			product, err := domain.NewProduct(tc.initialName, tc.initialDescription, tc.initialCategoryID)
			s.Require().NoError(err)
			s.Require().NotNil(product)

			initialUpdateTime := product.UpdatedAt
			time.Sleep(10 * time.Millisecond)

			product.Update(tc.updateName, tc.updateDescription, tc.updateCategoryID)

			s.Equal(tc.expectedName, product.Name, tc.name)
			s.Equal(tc.expectedDesc, product.Description, tc.name)
			if tc.expectedCatID != uuid.Nil {
				s.Equal(tc.expectedCatID, product.CategoryID, tc.name)
			} else if tc.updateCategoryID == uuid.Nil {
				s.Equal(tc.initialCategoryID, product.CategoryID, tc.name)
			}

			if tc.shouldUpdateTime {
				s.True(product.UpdatedAt.After(initialUpdateTime), tc.name)
			} else {
				s.Equal(initialUpdateTime, product.UpdatedAt, tc.name)
			}
		})
	}
}

func (s *ProductTestSuite) TestProductUpdateVariant() {
	testcases := []struct {
		name             string
		setupVariants    int
		targetIndex      int
		updatePrice      int64
		updateQuantity   int
		expectErr        bool
		expectedPrice    int64
		expectedQuantity int
		shouldUpdateTime bool
	}{
		{
			name:             "update price only",
			setupVariants:    1,
			targetIndex:      0,
			updatePrice:      50000,
			updateQuantity:   0,
			expectErr:        false,
			expectedPrice:    50000,
			expectedQuantity: 10,
			shouldUpdateTime: true,
		},
		{
			name:             "update quantity only",
			setupVariants:    1,
			targetIndex:      0,
			updatePrice:      0,
			updateQuantity:   20,
			expectErr:        false,
			expectedPrice:    10000,
			expectedQuantity: 20,
			shouldUpdateTime: true,
		},
		{
			name:             "update both price and quantity",
			setupVariants:    1,
			targetIndex:      0,
			updatePrice:      60000,
			updateQuantity:   30,
			expectErr:        false,
			expectedPrice:    60000,
			expectedQuantity: 30,
			shouldUpdateTime: true,
		},
		{
			name:             "update non-existent variant",
			setupVariants:    1,
			targetIndex:      -1,
			updatePrice:      50000,
			updateQuantity:   20,
			expectErr:        true,
			expectedPrice:    0,
			expectedQuantity: 0,
			shouldUpdateTime: false,
		},
		{
			name:             "update with zero values (no change)",
			setupVariants:    1,
			targetIndex:      0,
			updatePrice:      0,
			updateQuantity:   0,
			expectErr:        false,
			expectedPrice:    10000,
			expectedQuantity: 10,
			shouldUpdateTime: false,
		},
		{
			name:             "update second variant in list",
			setupVariants:    3,
			targetIndex:      1,
			updatePrice:      75000,
			updateQuantity:   15,
			expectErr:        false,
			expectedPrice:    75000,
			expectedQuantity: 15,
			shouldUpdateTime: true,
		},
	}

	for _, tc := range testcases {
		s.Run(tc.name, func() {
			product, err := domain.NewProduct("Test Product", "Test Description", uuid.New())
			s.Require().NoError(err)

			for i := 0; i < tc.setupVariants; i++ {
				variant, err := domain.NewVariant("SKU-"+string(rune('A'+i)), 10000, 10)
				s.Require().NoError(err)
				product.AddVariants(*variant)
			}

			var targetID uuid.UUID
			var initialUpdateTime time.Time
			if tc.targetIndex == -1 {
				targetID = uuid.New()
			} else {
				targetID = product.Variants[tc.targetIndex].ID
				initialUpdateTime = product.Variants[tc.targetIndex].UpdatedAt
			}

			time.Sleep(10 * time.Millisecond)

			err = product.UpdateVariant(targetID, tc.updatePrice, tc.updateQuantity)

			if tc.expectErr {
				s.Error(err, tc.name)
				s.ErrorIs(err, domain.ErrNotFound, tc.name)
			} else {
				s.NoError(err, tc.name)
				if tc.targetIndex >= 0 {
					s.Equal(tc.expectedPrice, product.Variants[tc.targetIndex].Price, tc.name)
					s.Equal(tc.expectedQuantity, product.Variants[tc.targetIndex].Quantity, tc.name)

					if tc.shouldUpdateTime {
						s.True(product.Variants[tc.targetIndex].UpdatedAt.After(initialUpdateTime), tc.name)
					} else {
						s.Equal(initialUpdateTime, product.Variants[tc.targetIndex].UpdatedAt, tc.name)
					}
				}
			}
		})
	}
}

func (s *ProductTestSuite) TestProductUpdateOption() {
	testcases := []struct {
		name         string
		setupOptions []string
		targetIndex  int
		newName      string
		expectErr    bool
		expectedName string
	}{
		{
			name:         "update existing option",
			setupOptions: []string{"Size"},
			targetIndex:  0,
			newName:      "Updated Size",
			expectErr:    false,
			expectedName: "Updated Size",
		},
		{
			name:         "update with empty name (no change)",
			setupOptions: []string{"Size"},
			targetIndex:  0,
			newName:      "",
			expectErr:    false,
			expectedName: "Size",
		},
		{
			name:         "update non-existent option",
			setupOptions: []string{"Size"},
			targetIndex:  -1,
			newName:      "New Name",
			expectErr:    true,
			expectedName: "",
		},
		{
			name:         "update second option in list",
			setupOptions: []string{"Size", "Color", "Material"},
			targetIndex:  1,
			newName:      "Updated Color",
			expectErr:    false,
			expectedName: "Updated Color",
		},
	}

	for _, tc := range testcases {
		s.Run(tc.name, func() {
			product, err := domain.NewProduct("Test Product", "Test Description", uuid.New())
			s.Require().NoError(err)

			for _, optName := range tc.setupOptions {
				option, err := domain.NewProductOption(optName)
				s.Require().NoError(err)
				product.AddOptions(*option)
			}

			var targetID uuid.UUID
			if tc.targetIndex == -1 {
				targetID = uuid.New()
			} else {
				targetID = product.Options[tc.targetIndex].ID
			}

			err = product.UpdateOption(targetID, tc.newName)

			if tc.expectErr {
				s.Error(err, tc.name)
				s.ErrorIs(err, domain.ErrNotFound, tc.name)
			} else {
				s.NoError(err, tc.name)
				if tc.targetIndex >= 0 {
					s.Equal(tc.expectedName, product.Options[tc.targetIndex].Name, tc.name)
				}
			}
		})
	}
}

func (s *ProductTestSuite) TestProductUpdateOptionValue() {
	testcases := []struct {
		name            string
		setupValues     []string
		targetValIndex  int
		newValue        string
		optionNotExists bool
		expectErr       bool
		expectedValue   string
	}{
		{
			name:            "update existing option value",
			setupValues:     []string{"Red", "Blue"},
			targetValIndex:  0,
			newValue:        "Green",
			optionNotExists: false,
			expectErr:       false,
			expectedValue:   "Green",
		},
		{
			name:            "update with empty value (no change)",
			setupValues:     []string{"Red"},
			targetValIndex:  0,
			newValue:        "",
			optionNotExists: false,
			expectErr:       false,
			expectedValue:   "Red",
		},
		{
			name:            "update non-existent option value",
			setupValues:     []string{"Red"},
			targetValIndex:  -1,
			newValue:        "Green",
			optionNotExists: false,
			expectErr:       true,
			expectedValue:   "",
		},
		{
			name:            "update with non-existent option ID",
			setupValues:     []string{"Red"},
			targetValIndex:  0,
			newValue:        "Green",
			optionNotExists: true,
			expectErr:       true,
			expectedValue:   "",
		},
	}

	for _, tc := range testcases {
		s.Run(tc.name, func() {
			product, err := domain.NewProduct("Test Product", "Test Description", uuid.New())
			s.Require().NoError(err)

			option, err := domain.NewProductOption("Color")
			s.Require().NoError(err)

			optionValues, err := domain.CreateOptionValues(tc.setupValues)
			s.Require().NoError(err)
			option.AddOptionValues(optionValues...)

			product.AddOptions(*option)

			var optionID, targetValueID uuid.UUID
			if tc.optionNotExists {
				optionID = uuid.New()
				targetValueID = uuid.New()
			} else {
				optionID = option.ID
				if tc.targetValIndex == -1 {
					targetValueID = uuid.New()
				} else {
					targetValueID = option.Values[tc.targetValIndex].ID
				}
			}

			err = product.UpdateOptionValue(optionID, targetValueID, tc.newValue)

			if tc.expectErr {
				s.Error(err, tc.name)
				s.ErrorIs(err, domain.ErrNotFound, tc.name)
			} else {
				s.NoError(err, tc.name)
				if tc.targetValIndex >= 0 {
					s.Equal(tc.expectedValue, product.Options[0].Values[tc.targetValIndex].Value, tc.name)
				}
			}
		})
	}
}

func (s *ProductTestSuite) TestProductGetOptionByID() {
	product, err := domain.NewProduct("Test Product", "Test Description", uuid.New())
	s.Require().NoError(err)

	option1, _ := domain.NewProductOption("Size")
	option2, _ := domain.NewProductOption("Color")
	product.AddOptions(*option1, *option2)

	testcases := []struct {
		name      string
		optionID  uuid.UUID
		expectNil bool
	}{
		{
			name:      "get existing option",
			optionID:  option1.ID,
			expectNil: false,
		},
		{
			name:      "get non-existent option",
			optionID:  uuid.New(),
			expectNil: true,
		},
	}

	for _, tc := range testcases {
		s.Run(tc.name, func() {
			result := product.GetOptionByID(tc.optionID)

			if tc.expectNil {
				s.Nil(result, tc.name)
			} else {
				s.NotNil(result, tc.name)
				s.Equal(tc.optionID, result.ID, tc.name)
			}
		})
	}
}

func (s *ProductTestSuite) TestProductGetOptionsByIDs() {
	product, err := domain.NewProduct("Test Product", "Test Description", uuid.New())
	s.Require().NoError(err)

	option1, _ := domain.NewProductOption("Size")
	option2, _ := domain.NewProductOption("Color")
	option3, _ := domain.NewProductOption("Material")
	product.AddOptions(*option1, *option2, *option3)

	testcases := []struct {
		name          string
		optionIDs     []uuid.UUID
		expectedCount int
	}{
		{
			name:          "get single option",
			optionIDs:     []uuid.UUID{option1.ID},
			expectedCount: 1,
		},
		{
			name:          "get multiple options",
			optionIDs:     []uuid.UUID{option1.ID, option2.ID},
			expectedCount: 2,
		},
		{
			name:          "get all options",
			optionIDs:     []uuid.UUID{option1.ID, option2.ID, option3.ID},
			expectedCount: 3,
		},
		{
			name:          "get non-existent options",
			optionIDs:     []uuid.UUID{uuid.New(), uuid.New()},
			expectedCount: 0,
		},
		{
			name:          "get mixed (existing and non-existing)",
			optionIDs:     []uuid.UUID{option1.ID, uuid.New()},
			expectedCount: 1,
		},
		{
			name:          "empty option IDs",
			optionIDs:     []uuid.UUID{},
			expectedCount: 0,
		},
	}

	for _, tc := range testcases {
		s.Run(tc.name, func() {
			result := product.GetOptionsByIDs(tc.optionIDs)
			s.Len(result, tc.expectedCount, tc.name)
		})
	}
}

func (s *ProductTestSuite) TestProductGetVariantByID() {
	product, err := domain.NewProduct("Test Product", "Test Description", uuid.New())
	s.Require().NoError(err)

	variant1, _ := domain.NewVariant("SKU-001", 10000, 10)
	variant2, _ := domain.NewVariant("SKU-002", 20000, 20)
	product.AddVariants(*variant1, *variant2)

	testcases := []struct {
		name      string
		variantID uuid.UUID
		expectNil bool
	}{
		{
			name:      "get existing variant",
			variantID: variant1.ID,
			expectNil: false,
		},
		{
			name:      "get non-existent variant",
			variantID: uuid.New(),
			expectNil: true,
		},
	}

	for _, tc := range testcases {
		s.Run(tc.name, func() {
			result := product.GetVariantByID(tc.variantID)

			if tc.expectNil {
				s.Nil(result, tc.name)
			} else {
				s.NotNil(result, tc.name)
				s.Equal(tc.variantID, result.ID, tc.name)
			}
		})
	}
}

func (s *ProductTestSuite) TestProductUpdateMinPrice() {
	testcases := []struct {
		name          string
		variantPrices []int64
		expectedPrice int64
	}{
		{
			name:          "single variant",
			variantPrices: []int64{10000},
			expectedPrice: 10000,
		},
		{
			name:          "multiple variants - ascending",
			variantPrices: []int64{10000, 20000, 30000},
			expectedPrice: 10000,
		},
		{
			name:          "multiple variants - descending",
			variantPrices: []int64{30000, 20000, 10000},
			expectedPrice: 10000,
		},
		{
			name:          "multiple variants - random order",
			variantPrices: []int64{25000, 10000, 35000, 15000},
			expectedPrice: 10000,
		},
		{
			name:          "all same prices",
			variantPrices: []int64{20000, 20000, 20000},
			expectedPrice: 20000,
		},
		{
			name:          "no variants",
			variantPrices: []int64{},
			expectedPrice: 0,
		},
	}

	for _, tc := range testcases {
		s.Run(tc.name, func() {
			product, err := domain.NewProduct("Test Product", "Test Description", uuid.New())
			s.Require().NoError(err)

			for i, price := range tc.variantPrices {
				variant, err := domain.NewVariant("SKU-"+string(rune('A'+i)), price, 10)
				s.Require().NoError(err)
				product.AddVariants(*variant)
			}

			product.UpdateMinPrice()

			s.Equal(tc.expectedPrice, product.Price, tc.name)
		})
	}
}

func (s *ProductTestSuite) TestProductAddVariantImages() {
	buildURL := func(id uuid.UUID) string {
		return "https://example.com/images/" + id.String()
	}

	testcases := []struct {
		name          string
		setupVariants int
		targetIndex   int
		imagesToAdd   int
		expectErr     bool
		expectedCount int
	}{
		{
			name:          "add images to existing variant",
			setupVariants: 1,
			targetIndex:   0,
			imagesToAdd:   2,
			expectErr:     false,
			expectedCount: 2,
		},
		{
			name:          "add images to non-existent variant",
			setupVariants: 1,
			targetIndex:   -1,
			imagesToAdd:   2,
			expectErr:     true,
			expectedCount: 0,
		},
		{
			name:          "add no images",
			setupVariants: 1,
			targetIndex:   0,
			imagesToAdd:   0,
			expectErr:     false,
			expectedCount: 0,
		},
		{
			name:          "add multiple images",
			setupVariants: 1,
			targetIndex:   0,
			imagesToAdd:   5,
			expectErr:     false,
			expectedCount: 5,
		},
	}

	for _, tc := range testcases {
		s.Run(tc.name, func() {
			product, err := domain.NewProduct("Test Product", "Test Description", uuid.New())
			s.Require().NoError(err)

			for i := 0; i < tc.setupVariants; i++ {
				variant, err := domain.NewVariant("SKU-"+string(rune('A'+i)), 10000, 10)
				s.Require().NoError(err)
				product.AddVariants(*variant)
			}

			var targetID uuid.UUID
			if tc.targetIndex == -1 {
				targetID = uuid.New()
			} else {
				targetID = product.Variants[tc.targetIndex].ID
			}

			images := make([]domain.ProductImage, 0, tc.imagesToAdd)
			for i := 0; i < tc.imagesToAdd; i++ {
				image, err := domain.NewProductImage(i, buildURL)
				s.Require().NoError(err)
				images = append(images, *image)
			}

			err = product.AddVariantImages(targetID, images...)

			if tc.expectErr {
				s.Error(err, tc.name)
				s.ErrorIs(err, domain.ErrNotFound, tc.name)
			} else {
				s.NoError(err, tc.name)
				if tc.targetIndex >= 0 {
					s.Len(product.Variants[tc.targetIndex].Images, tc.expectedCount, tc.name)
				}
			}
		})
	}
}

func (s *ProductTestSuite) TestProductRemove() {
	buildURL := func(id uuid.UUID) string {
		return "https://example.com/images/" + id.String()
	}

	product, err := domain.NewProduct("Test Product", "Test Description", uuid.New())
	s.Require().NoError(err)

	option1, _ := domain.NewProductOption("Size")
	optionValues, _ := domain.CreateOptionValues([]string{"S", "M", "L"})
	option1.AddOptionValues(optionValues...)
	product.AddOptions(*option1)

	variant1, _ := domain.NewVariant("SKU-001", 10000, 10)
	product.AddVariants(*variant1)

	image1, _ := domain.NewProductImage(0, buildURL)
	product.AddImages(*image1)

	beforeTime := time.Now()
	product.Remove()
	afterTime := time.Now()

	s.NotZero(product.DeletedAt, "product DeletedAt should be set")
	s.True(product.DeletedAt.After(beforeTime) || product.DeletedAt.Equal(beforeTime), "product DeletedAt should be after or equal to beforeTime")
	s.True(product.DeletedAt.Before(afterTime) || product.DeletedAt.Equal(afterTime), "product DeletedAt should be before or equal to afterTime")

	s.NotZero(product.UpdatedAt, "product UpdatedAt should be set")

	for _, opt := range product.Options {
		s.NotZero(opt.DeletedAt, "option DeletedAt should be set")
		for _, val := range opt.Values {
			s.NotZero(val.DeletedAt, "option value DeletedAt should be set")
		}
	}

	for _, variant := range product.Variants {
		s.NotZero(variant.DeletedAt, "variant DeletedAt should be set")
		s.NotZero(variant.UpdatedAt, "variant UpdatedAt should be set")
	}

	for _, image := range product.Images {
		s.NotZero(image.DeletedAt, "image DeletedAt should be set")
	}
}

func (s *ProductTestSuite) TestOptionGetValueByID() {
	option, err := domain.NewProductOption("Color")
	s.Require().NoError(err)

	values, err := domain.CreateOptionValues([]string{"Red", "Blue", "Green"})
	s.Require().NoError(err)
	option.AddOptionValues(values...)

	testcases := []struct {
		name      string
		valueID   uuid.UUID
		expectNil bool
	}{
		{
			name:      "get existing value",
			valueID:   values[0].ID,
			expectNil: false,
		},
		{
			name:      "get non-existent value",
			valueID:   uuid.New(),
			expectNil: true,
		},
	}

	for _, tc := range testcases {
		s.Run(tc.name, func() {
			result := option.GetValueByID(tc.valueID)

			if tc.expectNil {
				s.Nil(result, tc.name)
			} else {
				s.NotNil(result, tc.name)
				s.Equal(tc.valueID, result.ID, tc.name)
			}
		})
	}
}

func (s *ProductTestSuite) TestOptionGetValuesByIDs() {
	option, err := domain.NewProductOption("Color")
	s.Require().NoError(err)

	values, err := domain.CreateOptionValues([]string{"Red", "Blue", "Green", "Yellow"})
	s.Require().NoError(err)
	option.AddOptionValues(values...)

	testcases := []struct {
		name          string
		valueIDs      []uuid.UUID
		expectedCount int
	}{
		{
			name:          "get single value",
			valueIDs:      []uuid.UUID{values[0].ID},
			expectedCount: 1,
		},
		{
			name:          "get multiple values",
			valueIDs:      []uuid.UUID{values[0].ID, values[1].ID},
			expectedCount: 2,
		},
		{
			name:          "get all values",
			valueIDs:      []uuid.UUID{values[0].ID, values[1].ID, values[2].ID, values[3].ID},
			expectedCount: 4,
		},
		{
			name:          "get non-existent values",
			valueIDs:      []uuid.UUID{uuid.New(), uuid.New()},
			expectedCount: 0,
		},
		{
			name:          "get mixed (existing and non-existing)",
			valueIDs:      []uuid.UUID{values[0].ID, uuid.New()},
			expectedCount: 1,
		},
		{
			name:          "empty value IDs",
			valueIDs:      []uuid.UUID{},
			expectedCount: 0,
		},
	}

	for _, tc := range testcases {
		s.Run(tc.name, func() {
			result := option.GetValuesByIDs(tc.valueIDs)
			s.Len(result, tc.expectedCount, tc.name)
		})
	}
}

func (s *ProductTestSuite) TestProductVariantDecreaseQuantity() {
	testcases := []struct {
		name                  string
		initialQuantity       int
		decreaseBy            int
		expectedQuantity      int
		expectedPurchaseCount int
	}{
		{
			name:                  "decrease by valid amount",
			initialQuantity:       10,
			decreaseBy:            5,
			expectedQuantity:      5,
			expectedPurchaseCount: 5,
		},
		{
			name:                  "decrease to zero",
			initialQuantity:       10,
			decreaseBy:            10,
			expectedQuantity:      0,
			expectedPurchaseCount: 10,
		},
		{
			name:                  "decrease by more than available (floor to zero)",
			initialQuantity:       10,
			decreaseBy:            15,
			expectedQuantity:      0,
			expectedPurchaseCount: 15,
		},
		{
			name:                  "decrease by zero (no change)",
			initialQuantity:       10,
			decreaseBy:            0,
			expectedQuantity:      10,
			expectedPurchaseCount: 0,
		},
		{
			name:                  "decrease by negative (no change)",
			initialQuantity:       10,
			decreaseBy:            -5,
			expectedQuantity:      10,
			expectedPurchaseCount: 0,
		},
		{
			name:                  "decrease from zero quantity",
			initialQuantity:       0,
			decreaseBy:            5,
			expectedQuantity:      0,
			expectedPurchaseCount: 5,
		},
	}

	for _, tc := range testcases {
		s.Run(tc.name, func() {
			variant, err := domain.NewVariant("SKU-001", 10000, tc.initialQuantity)
			s.Require().NoError(err)

			initialUpdateTime := variant.UpdatedAt
			time.Sleep(10 * time.Millisecond)

			variant.DecreaseQuantity(tc.decreaseBy)

			s.Equal(tc.expectedQuantity, variant.Quantity, tc.name)
			s.Equal(tc.expectedPurchaseCount, variant.PurchaseCount, tc.name)

			if tc.decreaseBy > 0 {
				s.True(variant.UpdatedAt.After(initialUpdateTime), tc.name)
			}
		})
	}
}

func (s *ProductTestSuite) TestProductAddAttributeIDs() {
	product, err := domain.NewProduct("Test Product", "Test Description", uuid.New())
	s.Require().NoError(err)

	attrID1 := uuid.New()
	attrID2 := uuid.New()
	attrID3 := uuid.New()

	product.AddAttributeIDs(attrID1)
	s.Len(product.AttributeIDs, 1)
	s.Equal(attrID1, product.AttributeIDs[0])

	product.AddAttributeIDs(attrID2, attrID3)
	s.Len(product.AttributeIDs, 3)
	s.Equal(attrID2, product.AttributeIDs[1])
	s.Equal(attrID3, product.AttributeIDs[2])
}

func (s *ProductTestSuite) TestProductAddAttributeValueIDs() {
	product, err := domain.NewProduct("Test Product", "Test Description", uuid.New())
	s.Require().NoError(err)

	valID1 := uuid.New()
	valID2 := uuid.New()

	product.AddAttributeValueIDs(valID1)
	s.Len(product.AttributeValueIDs, 1)
	s.Equal(valID1, product.AttributeValueIDs[0])

	product.AddAttributeValueIDs(valID2)
	s.Len(product.AttributeValueIDs, 2)
	s.Equal(valID2, product.AttributeValueIDs[1])
}

func (s *ProductTestSuite) TestProductAddOptions() {
	product, err := domain.NewProduct("Test Product", "Test Description", uuid.New())
	s.Require().NoError(err)

	option1, _ := domain.NewProductOption("Size")
	option2, _ := domain.NewProductOption("Color")

	product.AddOptions(*option1)
	s.Len(product.Options, 1)
	s.Equal(option1.ID, product.Options[0].ID)

	product.AddOptions(*option2)
	s.Len(product.Options, 2)
	s.Equal(option2.ID, product.Options[1].ID)
}

func (s *ProductTestSuite) TestProductAddVariants() {
	product, err := domain.NewProduct("Test Product", "Test Description", uuid.New())
	s.Require().NoError(err)

	variant1, _ := domain.NewVariant("SKU-001", 10000, 10)
	variant2, _ := domain.NewVariant("SKU-002", 20000, 20)

	product.AddVariants(*variant1)
	s.Len(product.Variants, 1)
	s.Equal(variant1.ID, product.Variants[0].ID)

	product.AddVariants(*variant2)
	s.Len(product.Variants, 2)
	s.Equal(variant2.ID, product.Variants[1].ID)
}

func (s *ProductTestSuite) TestProductAddImages() {
	product, err := domain.NewProduct("Test Product", "Test Description", uuid.New())
	s.Require().NoError(err)

	buildURL := func(id uuid.UUID) string {
		return "https://example.com/images/" + id.String()
	}

	image1, _ := domain.NewProductImage(0, buildURL)
	image2, _ := domain.NewProductImage(1, buildURL)

	product.AddImages(*image1)
	s.Len(product.Images, 1)
	s.Equal(image1.ID, product.Images[0].ID)

	product.AddImages(*image2)
	s.Len(product.Images, 2)
	s.Equal(image2.ID, product.Images[1].ID)
}

func (s *ProductTestSuite) TestOptionAddOptionValues() {
	option, err := domain.NewProductOption("Color")
	s.Require().NoError(err)

	values1, _ := domain.CreateOptionValues([]string{"Red"})
	values2, _ := domain.CreateOptionValues([]string{"Blue", "Green"})

	option.AddOptionValues(values1...)
	s.Len(option.Values, 1)
	s.Equal("Red", option.Values[0].Value)

	option.AddOptionValues(values2...)
	s.Len(option.Values, 3)
	s.Equal("Blue", option.Values[1].Value)
	s.Equal("Green", option.Values[2].Value)
}

func (s *ProductTestSuite) TestOptionRemove() {
	option, err := domain.NewProductOption("Color")
	s.Require().NoError(err)

	values, _ := domain.CreateOptionValues([]string{"Red", "Blue"})
	option.AddOptionValues(values...)

	beforeTime := time.Now()
	option.Remove()
	afterTime := time.Now()

	s.NotZero(option.DeletedAt)
	s.True(option.DeletedAt.After(beforeTime) || option.DeletedAt.Equal(beforeTime))
	s.True(option.DeletedAt.Before(afterTime) || option.DeletedAt.Equal(afterTime))

	for _, val := range option.Values {
		s.NotZero(val.DeletedAt)
	}
}

func (s *ProductTestSuite) TestOptionValueRemove() {
	values, _ := domain.CreateOptionValues([]string{"Red"})
	optionValue := values[0]

	beforeTime := time.Now()
	optionValue.Remove()
	afterTime := time.Now()

	s.NotZero(optionValue.DeletedAt)
	s.True(optionValue.DeletedAt.After(beforeTime) || optionValue.DeletedAt.Equal(beforeTime))
	s.True(optionValue.DeletedAt.Before(afterTime) || optionValue.DeletedAt.Equal(afterTime))
}

func (s *ProductTestSuite) TestProductVariantRemove() {
	variant, err := domain.NewVariant("SKU-001", 10000, 10)
	s.Require().NoError(err)

	beforeTime := time.Now()
	variant.Remove()
	afterTime := time.Now()

	s.NotZero(variant.DeletedAt)
	s.NotZero(variant.UpdatedAt)
	s.True(variant.DeletedAt.After(beforeTime) || variant.DeletedAt.Equal(beforeTime))
	s.True(variant.DeletedAt.Before(afterTime) || variant.DeletedAt.Equal(afterTime))
}

func (s *ProductTestSuite) TestProductImageRemove() {
	buildURL := func(id uuid.UUID) string {
		return "https://example.com/images/" + id.String()
	}
	image, err := domain.NewProductImage(0, buildURL)
	s.Require().NoError(err)

	beforeTime := time.Now()
	image.Remove()
	afterTime := time.Now()

	s.NotZero(image.DeletedAt)
	s.True(image.DeletedAt.After(beforeTime) || image.DeletedAt.Equal(beforeTime))
	s.True(image.DeletedAt.Before(afterTime) || image.DeletedAt.Equal(afterTime))
}

func TestProduct(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(ProductTestSuite))
}
