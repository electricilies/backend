package service

import (
	"testing"

	"backend/internal/domain"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
)

type AttributeServiceTestSuite struct {
	suite.Suite
	validator *validator.Validate
	service   *Attribute
}

func TestAttributeServiceTestSuite(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(AttributeServiceTestSuite))
}

func (s *AttributeServiceTestSuite) SetupTest() {
	s.validator = validator.New()
	err := domain.RegisterAttributeValidators(s.validator)
	s.Require().NoError(err)
	s.service = ProvideAttribute(s.validator)
}

func (s *AttributeServiceTestSuite) TestProvideAttribute() {
	tests := []struct {
		name string
	}{
		{
			name: "creates attribute service successfully",
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			validate := validator.New()
			err := domain.RegisterAttributeValidators(validate)
			s.Require().NoError(err)

			service := ProvideAttribute(validate)

			s.NotNil(service)
			s.NotNil(service.validate)
		})
	}
}

func (s *AttributeServiceTestSuite) TestAttribute_Validate() {
	tests := []struct {
		name        string
		setupAttr   func() domain.Attribute
		expectError bool
		errorType   error
	}{
		{
			name: "validates valid attribute successfully",
			setupAttr: func() domain.Attribute {
				attr, _ := domain.NewAttribute("color", "Color")
				val1, _ := domain.NewAttributeValue("Red")
				val2, _ := domain.NewAttributeValue("Blue")
				attr.AddValues(*val1, *val2)
				return *attr
			},
			expectError: false,
		},
		{
			name: "returns error for empty code",
			setupAttr: func() domain.Attribute {
				return domain.Attribute{
					ID:     uuid.Must(uuid.NewV7()),
					Code:   "",
					Name:   "Color",
					Values: []domain.AttributeValue{},
				}
			},
			expectError: true,
			errorType:   domain.ErrInvalid,
		},
		{
			name: "returns error for empty name",
			setupAttr: func() domain.Attribute {
				return domain.Attribute{
					ID:     uuid.Must(uuid.NewV7()),
					Code:   "color",
					Name:   "",
					Values: []domain.AttributeValue{},
				}
			},
			expectError: true,
			errorType:   domain.ErrInvalid,
		},
		{
			name: "returns error for code too short",
			setupAttr: func() domain.Attribute {
				return domain.Attribute{
					ID:     uuid.Must(uuid.NewV7()),
					Code:   "c",
					Name:   "Color",
					Values: []domain.AttributeValue{},
				}
			},
			expectError: true,
			errorType:   domain.ErrInvalid,
		},
		{
			name: "returns error for code too long",
			setupAttr: func() domain.Attribute {
				return domain.Attribute{
					ID:     uuid.Must(uuid.NewV7()),
					Code:   "this_is_a_very_long_code_that_exceeds_fifty_characters_limit_by_far",
					Name:   "Color",
					Values: []domain.AttributeValue{},
				}
			},
			expectError: true,
			errorType:   domain.ErrInvalid,
		},
		{
			name: "returns error for name too short",
			setupAttr: func() domain.Attribute {
				return domain.Attribute{
					ID:     uuid.Must(uuid.NewV7()),
					Code:   "color",
					Name:   "C",
					Values: []domain.AttributeValue{},
				}
			},
			expectError: true,
			errorType:   domain.ErrInvalid,
		},
		{
			name: "returns error for name too long",
			setupAttr: func() domain.Attribute {
				return domain.Attribute{
					ID:     uuid.Must(uuid.NewV7()),
					Code:   "color",
					Name:   "This is a very long name that exceeds one hundred characters limit and should fail validation test case",
					Values: []domain.AttributeValue{},
				}
			},
			expectError: true,
			errorType:   domain.ErrInvalid,
		},
		{
			name: "returns error for invalid attribute value",
			setupAttr: func() domain.Attribute {
				attr, _ := domain.NewAttribute("color", "Color")
				invalidValue := domain.AttributeValue{
					ID:    uuid.Must(uuid.NewV7()),
					Value: "",
				}
				attr.AddValues(invalidValue)
				return *attr
			},
			expectError: true,
			errorType:   domain.ErrInvalid,
		},
		{
			name: "returns error for attribute value too long",
			setupAttr: func() domain.Attribute {
				attr, _ := domain.NewAttribute("color", "Color")
				longValue := domain.AttributeValue{
					ID:    uuid.Must(uuid.NewV7()),
					Value: "This is a very long value that exceeds one hundred characters limit and should fail validation test",
				}
				attr.AddValues(longValue)
				return *attr
			},
			expectError: true,
			errorType:   domain.ErrInvalid,
		},
		{
			name: "validates attribute with nil ID as invalid",
			setupAttr: func() domain.Attribute {
				return domain.Attribute{
					ID:     uuid.Nil,
					Code:   "color",
					Name:   "Color",
					Values: []domain.AttributeValue{},
				}
			},
			expectError: true,
			errorType:   domain.ErrInvalid,
		},
		{
			name: "validates attribute with multiple valid values",
			setupAttr: func() domain.Attribute {
				attr, _ := domain.NewAttribute("color", "Color")
				val1, _ := domain.NewAttributeValue("Red")
				val2, _ := domain.NewAttributeValue("Blue")
				val3, _ := domain.NewAttributeValue("Green")
				val4, _ := domain.NewAttributeValue("Yellow")
				attr.AddValues(*val1, *val2, *val3, *val4)
				return *attr
			},
			expectError: false,
		},
		{
			name: "returns error for duplicate attribute values",
			setupAttr: func() domain.Attribute {
				attr, _ := domain.NewAttribute("color", "Color")
				val1, _ := domain.NewAttributeValue("Red")
				val2, _ := domain.NewAttributeValue("Red")
				attr.AddValues(*val1, *val2)
				return *attr
			},
			expectError: true,
			errorType:   domain.ErrInvalid,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			attr := tt.setupAttr()

			err := s.service.Validate(attr)

			if tt.expectError {
				s.Error(err)
				if tt.errorType != nil {
					s.ErrorIs(err, tt.errorType)
				}
			} else {
				s.NoError(err)
			}
		})
	}
}

func (s *AttributeServiceTestSuite) TestAttribute_FilterAttributeValuesFromAttributes() {
	tests := []struct {
		name           string
		setupData      func() ([]domain.Attribute, []uuid.UUID)
		expectedCount  int
		validateResult func([]domain.AttributeValue, []domain.Attribute)
	}{
		{
			name: "filters attribute values successfully",
			setupData: func() ([]domain.Attribute, []uuid.UUID) {
				attr1, _ := domain.NewAttribute("color", "Color")
				val1, _ := domain.NewAttributeValue("Red")
				val2, _ := domain.NewAttributeValue("Blue")
				attr1.AddValues(*val1, *val2)

				attr2, _ := domain.NewAttribute("size", "Size")
				val3, _ := domain.NewAttributeValue("Small")
				val4, _ := domain.NewAttributeValue("Large")
				attr2.AddValues(*val3, *val4)

				attributes := []domain.Attribute{*attr1, *attr2}
				attributeValueIDs := []uuid.UUID{val1.ID, val3.ID}
				return attributes, attributeValueIDs
			},
			expectedCount: 2,
			validateResult: func(result []domain.AttributeValue, attrs []domain.Attribute) {
				s.Contains(result, attrs[0].Values[0])
				s.Contains(result, attrs[1].Values[0])
			},
		},
		{
			name: "returns empty slice when no matching values",
			setupData: func() ([]domain.Attribute, []uuid.UUID) {
				attr1, _ := domain.NewAttribute("color", "Color")
				val1, _ := domain.NewAttributeValue("Red")
				attr1.AddValues(*val1)

				attributes := []domain.Attribute{*attr1}
				randomID := uuid.Must(uuid.NewV7())
				attributeValueIDs := []uuid.UUID{randomID}
				return attributes, attributeValueIDs
			},
			expectedCount: 0,
		},
		{
			name: "returns empty slice for empty attributes",
			setupData: func() ([]domain.Attribute, []uuid.UUID) {
				attributes := []domain.Attribute{}
				attributeValueIDs := []uuid.UUID{uuid.Must(uuid.NewV7())}
				return attributes, attributeValueIDs
			},
			expectedCount: 0,
		},
		{
			name: "returns empty slice for empty value IDs",
			setupData: func() ([]domain.Attribute, []uuid.UUID) {
				attr1, _ := domain.NewAttribute("color", "Color")
				val1, _ := domain.NewAttributeValue("Red")
				attr1.AddValues(*val1)

				attributes := []domain.Attribute{*attr1}
				attributeValueIDs := []uuid.UUID{}
				return attributes, attributeValueIDs
			},
			expectedCount: 0,
		},
		{
			name: "filters values from multiple attributes",
			setupData: func() ([]domain.Attribute, []uuid.UUID) {
				attr1, _ := domain.NewAttribute("color", "Color")
				val1, _ := domain.NewAttributeValue("Red")
				val2, _ := domain.NewAttributeValue("Blue")
				val3, _ := domain.NewAttributeValue("Green")
				attr1.AddValues(*val1, *val2, *val3)

				attr2, _ := domain.NewAttribute("size", "Size")
				val4, _ := domain.NewAttributeValue("Small")
				val5, _ := domain.NewAttributeValue("Medium")
				val6, _ := domain.NewAttributeValue("Large")
				attr2.AddValues(*val4, *val5, *val6)

				attr3, _ := domain.NewAttribute("brand", "Brand")
				val7, _ := domain.NewAttributeValue("Nike")
				val8, _ := domain.NewAttributeValue("Adidas")
				attr3.AddValues(*val7, *val8)

				attributes := []domain.Attribute{*attr1, *attr2, *attr3}
				attributeValueIDs := []uuid.UUID{val2.ID, val5.ID, val7.ID, val8.ID}
				return attributes, attributeValueIDs
			},
			expectedCount: 4,
		},
		{
			name: "handles duplicate value IDs",
			setupData: func() ([]domain.Attribute, []uuid.UUID) {
				attr1, _ := domain.NewAttribute("color", "Color")
				val1, _ := domain.NewAttributeValue("Red")
				attr1.AddValues(*val1)

				attributes := []domain.Attribute{*attr1}
				attributeValueIDs := []uuid.UUID{val1.ID, val1.ID}
				return attributes, attributeValueIDs
			},
			expectedCount: 1,
		},
		{
			name: "filters all values when all IDs match",
			setupData: func() ([]domain.Attribute, []uuid.UUID) {
				attr1, _ := domain.NewAttribute("color", "Color")
				val1, _ := domain.NewAttributeValue("Red")
				val2, _ := domain.NewAttributeValue("Blue")
				attr1.AddValues(*val1, *val2)

				attributes := []domain.Attribute{*attr1}
				attributeValueIDs := []uuid.UUID{val1.ID, val2.ID}
				return attributes, attributeValueIDs
			},
			expectedCount: 2,
		},
		{
			name: "preserves order of found values",
			setupData: func() ([]domain.Attribute, []uuid.UUID) {
				attr1, _ := domain.NewAttribute("color", "Color")
				val1, _ := domain.NewAttributeValue("Red")
				val2, _ := domain.NewAttributeValue("Blue")
				val3, _ := domain.NewAttributeValue("Green")
				attr1.AddValues(*val1, *val2, *val3)

				attributes := []domain.Attribute{*attr1}
				attributeValueIDs := []uuid.UUID{val1.ID, val2.ID, val3.ID}
				return attributes, attributeValueIDs
			},
			expectedCount: 3,
			validateResult: func(result []domain.AttributeValue, attrs []domain.Attribute) {
				s.Require().Len(result, 3)
				s.Equal(attrs[0].Values[0].ID, result[0].ID)
				s.Equal(attrs[0].Values[1].ID, result[1].ID)
				s.Equal(attrs[0].Values[2].ID, result[2].ID)
			},
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			attributes, attributeValueIDs := tt.setupData()

			result := s.service.FilterAttributeValuesFromAttributes(attributes, attributeValueIDs)

			s.Len(result, tt.expectedCount)
			if tt.validateResult != nil {
				tt.validateResult(result, attributes)
			}
		})
	}
}
