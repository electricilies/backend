package domain

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
)

type AttributeValidatorTestSuite struct {
	suite.Suite
	validator *validator.Validate
}

func TestAttributeValidatorTestSuite(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(AttributeValidatorTestSuite))
}

func (s *AttributeValidatorTestSuite) SetupTest() {
	s.validator = validator.New()
	err := RegisterAttributeValidators(s.validator)
	s.Require().NoError(err)
}

func (s *AttributeValidatorTestSuite) TestRegisterAttributeValidators() {
	tests := []struct {
		name        string
		setup       func() *validator.Validate
		expectError bool
	}{
		{
			name: "registers validators successfully",
			setup: func() *validator.Validate {
				return validator.New()
			},
			expectError: false,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			v := tt.setup()
			err := RegisterAttributeValidators(v)

			if tt.expectError {
				s.Error(err)
			} else {
				s.NoError(err)
			}
		})
	}
}

func (s *AttributeValidatorTestSuite) TestValidateUniqueAttributeValues() {
	tests := []struct {
		name        string
		setup       func() *Attribute
		expectError bool
		errorMsg    string
	}{
		{
			name: "validates attribute with unique values",
			setup: func() *Attribute {
				attr, _ := NewAttribute("color", "Color")
				val1, _ := NewAttributeValue("Red")
				val2, _ := NewAttributeValue("Blue")
				val3, _ := NewAttributeValue("Green")
				attr.AddValues(*val1, *val2, *val3)
				return attr
			},
			expectError: false,
		},
		{
			name: "fails validation for duplicate values",
			setup: func() *Attribute {
				attr, _ := NewAttribute("color", "Color")
				val1, _ := NewAttributeValue("Red")
				val2, _ := NewAttributeValue("Red")
				attr.AddValues(*val1, *val2)
				return attr
			},
			expectError: true,
			errorMsg:    "unique_attribute_values",
		},
		{
			name: "fails validation for duplicate IDs",
			setup: func() *Attribute {
				attr, _ := NewAttribute("color", "Color")
				val1, _ := NewAttributeValue("Red")
				val2 := &AttributeValue{
					ID:    val1.ID,
					Value: "Blue",
				}
				attr.AddValues(*val1, *val2)
				return attr
			},
			expectError: true,
			errorMsg:    "unique_attribute_values",
		},
		{
			name: "validates attribute with empty values",
			setup: func() *Attribute {
				attr, _ := NewAttribute("color", "Color")
				return attr
			},
			expectError: false,
		},
		{
			name: "validates attribute with single value",
			setup: func() *Attribute {
				attr, _ := NewAttribute("color", "Color")
				val, _ := NewAttributeValue("Red")
				attr.AddValues(*val)
				return attr
			},
			expectError: false,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			attr := tt.setup()
			err := s.validator.Struct(attr)

			if tt.expectError {
				s.Require().Error(err)
				s.Contains(err.Error(), tt.errorMsg)
			} else {
				s.Require().NoError(err)
			}
		})
	}
}

func (s *AttributeValidatorTestSuite) TestValidateAttributeValueUniqueness_Extended() {
	tests := []struct {
		name        string
		setup       func() *Attribute
		expectError bool
		errorMsg    string
	}{
		{
			name: "validates multiple unique values",
			setup: func() *Attribute {
				attr, _ := NewAttribute("size", "Size")
				values := []string{"Small", "Medium", "Large", "X-Large", "XX-Large"}
				for _, v := range values {
					val, _ := NewAttributeValue(v)
					attr.AddValues(*val)
				}
				return attr
			},
			expectError: false,
		},
		{
			name: "detects duplicate in middle of list",
			setup: func() *Attribute {
				attr, _ := NewAttribute("color", "Color")
				val1, _ := NewAttributeValue("Red")
				val2, _ := NewAttributeValue("Blue")
				val3, _ := NewAttributeValue("Red")
				val4, _ := NewAttributeValue("Green")
				attr.AddValues(*val1, *val2, *val3, *val4)
				return attr
			},
			expectError: true,
			errorMsg:    "duplicate attribute value found",
		},
		{
			name: "validates attribute with nil values slice",
			setup: func() *Attribute {
				return &Attribute{
					ID:     uuid.Must(uuid.NewV7()),
					Code:   "color",
					Name:   "Color",
					Values: nil,
				}
			},
			expectError: false,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			attr := tt.setup()
			err := ValidateAttributeValueUniqueness(attr)

			if tt.expectError {
				s.Require().Error(err)
				s.Contains(err.Error(), tt.errorMsg)
			} else {
				s.Require().NoError(err)
			}
		})
	}
}

func (s *AttributeValidatorTestSuite) TestValidateUniqueAttributeIDs_Extended() {
	tests := []struct {
		name        string
		setup       func() []Attribute
		expectError bool
		errorMsg    string
	}{
		{
			name: "validates large number of unique IDs",
			setup: func() []Attribute {
				attributes := make([]Attribute, 100)
				for i := range 100 {
					attr, _ := NewAttribute("attr"+string(rune(i)), "Attribute")
					attributes[i] = *attr
				}
				return attributes
			},
			expectError: false,
		},
		{
			name: "detects duplicate ID in large list",
			setup: func() []Attribute {
				attributes := make([]Attribute, 10)
				for i := 0; i < 10; i++ {
					attr, _ := NewAttribute("attr"+string(rune(i)), "Attribute")
					attributes[i] = *attr
				}
				attributes[5].ID = attributes[3].ID
				return attributes
			},
			expectError: true,
			errorMsg:    "duplicate attribute ID found",
		},
		{
			name: "validates single attribute",
			setup: func() []Attribute {
				attr, _ := NewAttribute("color", "Color")
				return []Attribute{*attr}
			},
			expectError: false,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			attributes := tt.setup()
			err := ValidateUniqueAttributeIDs(attributes)

			if tt.expectError {
				s.Require().Error(err)
				s.Contains(err.Error(), tt.errorMsg)
			} else {
				s.Require().NoError(err)
			}
		})
	}
}

func (s *AttributeValidatorTestSuite) TestValidateUniqueAttributeCodes_Extended() {
	tests := []struct {
		name        string
		setup       func() []Attribute
		expectError bool
	}{
		{
			name: "validates codes with different cases",
			setup: func() []Attribute {
				attr1, _ := NewAttribute("color", "Color")
				attr2, _ := NewAttribute("Color", "Color Uppercase")
				attr3, _ := NewAttribute("COLOR", "Color All Caps")
				return []Attribute{*attr1, *attr2, *attr3}
			},
			expectError: false,
		},
		{
			name: "detects exact duplicate codes",
			setup: func() []Attribute {
				attr1, _ := NewAttribute("color", "Color")
				attr2, _ := NewAttribute("size", "Size")
				attr3, _ := NewAttribute("color", "Another Color")
				return []Attribute{*attr1, *attr2, *attr3}
			},
			expectError: true,
		},
		{
			name: "validates codes with special characters",
			setup: func() []Attribute {
				attr1, _ := NewAttribute("audio_power_output", "Audio Power Output")
				attr2, _ := NewAttribute("audio-speaker-quantity", "Audio Speaker Quantity")
				attr3, _ := NewAttribute("audio.technology", "Audio Technology")
				return []Attribute{*attr1, *attr2, *attr3}
			},
			expectError: false,
		},
		{
			name: "validates similar but different codes",
			setup: func() []Attribute {
				attr1, _ := NewAttribute("brand", "Brand")
				attr2, _ := NewAttribute("brand_country", "Brand Country")
				attr3, _ := NewAttribute("sub_brand", "Sub Brand")
				return []Attribute{*attr1, *attr2, *attr3}
			},
			expectError: false,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			attributes := tt.setup()
			err := ValidateUniqueAttributeCodes(attributes)

			if tt.expectError {
				s.Error(err)
			} else {
				s.Require().NoError(err)
			}
		})
	}
}

func (s *AttributeValidatorTestSuite) TestValidatorIntegration() {
	tests := []struct {
		name        string
		setup       func() *Attribute
		expectError bool
		errorType   string
	}{
		{
			name: "validates complete attribute with all constraints",
			setup: func() *Attribute {
				attr, _ := NewAttribute("brand", "Brand")
				val1, _ := NewAttributeValue("Nike")
				val2, _ := NewAttributeValue("Adidas")
				val3, _ := NewAttributeValue("Puma")
				attr.AddValues(*val1, *val2, *val3)
				return attr
			},
			expectError: false,
		},
		{
			name: "fails validation for invalid code length",
			setup: func() *Attribute {
				return &Attribute{
					ID:     uuid.Must(uuid.NewV7()),
					Code:   "a",
					Name:   "Valid Name",
					Values: []AttributeValue{},
				}
			},
			expectError: true,
		},
		{
			name: "fails validation for invalid name length",
			setup: func() *Attribute {
				return &Attribute{
					ID:     uuid.Must(uuid.NewV7()),
					Code:   "valid_code",
					Name:   "N",
					Values: []AttributeValue{},
				}
			},
			expectError: true,
		},
		{
			name: "fails validation for invalid attribute value",
			setup: func() *Attribute {
				attr, _ := NewAttribute("color", "Color")
				invalidValue := AttributeValue{
					ID:    uuid.Must(uuid.NewV7()),
					Value: "",
				}
				attr.AddValues(invalidValue)
				return attr
			},
			expectError: true,
		},
		{
			name: "fails validation for duplicate values and invalid value",
			setup: func() *Attribute {
				attr, _ := NewAttribute("color", "Color")
				val1, _ := NewAttributeValue("Red")
				val2, _ := NewAttributeValue("Red")
				val3 := AttributeValue{
					ID:    uuid.Must(uuid.NewV7()),
					Value: "",
				}
				attr.AddValues(*val1, *val2, val3)
				return attr
			},
			expectError: true,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			attr := tt.setup()
			err := s.validator.Struct(attr)

			if tt.expectError {
				s.Error(err)
			} else {
				s.NoError(err)
			}
		})
	}
}

func (s *AttributeValidatorTestSuite) TestRealWorldScenarios() {
	tests := []struct {
		name        string
		setup       func() []Attribute
		validate    func([]Attribute)
		expectError bool
	}{
		{
			name: "validates product attributes like in database",
			setup: func() []Attribute {
				orgAddress, _ := NewAttribute("Organization_address", "Địa chỉ tổ chức chịu trách nhiệm về hàng hóa")

				audioTech, _ := NewAttribute("audio_technology", "Công nghệ âm thanh")
				val1, _ := NewAttributeValue("Dolby Atmos")
				val2, _ := NewAttributeValue("DTS:X")
				val3, _ := NewAttributeValue("Stereo")
				audioTech.AddValues(*val1, *val2, *val3)

				brand, _ := NewAttribute("brand", "Thương hiệu")
				valBrand1, _ := NewAttributeValue("Masstel")
				valBrand2, _ := NewAttributeValue("Tecno")
				valBrand3, _ := NewAttributeValue("Realme")
				valBrand4, _ := NewAttributeValue("OPPO")
				brand.AddValues(*valBrand1, *valBrand2, *valBrand3, *valBrand4)

				return []Attribute{*orgAddress, *audioTech, *brand}
			},
			validate: func(attributes []Attribute) {
				for _, attr := range attributes {
					s.Require().NoError(s.validator.Struct(attr))
				}
				s.Require().NoError(ValidateUniqueAttributeIDs(attributes))
				s.Require().NoError(ValidateUniqueAttributeCodes(attributes))
			},
			expectError: false,
		},
		{
			name: "detects duplicate battery capacity values",
			setup: func() []Attribute {
				battery, _ := NewAttribute("battery_capacity", "Dung lượng pin")
				val1, _ := NewAttributeValue("2500")
				val2, _ := NewAttributeValue("4400")
				val3, _ := NewAttributeValue("2500")
				battery.AddValues(*val1, *val2, *val3)
				return []Attribute{*battery}
			},
			validate: func(attributes []Attribute) {
				err := ValidateAttributeValueUniqueness(&attributes[0])
				s.Require().Error(err)
				s.Contains(err.Error(), "duplicate attribute value found")
			},
			expectError: true,
		},
		{
			name: "validates screen size attribute",
			setup: func() []Attribute {
				screenSize, _ := NewAttribute("man_hinh_size", "Kích thước màn hình")
				val1, _ := NewAttributeValue("2.4 inch")
				val2, _ := NewAttributeValue("6.7 inch")
				screenSize.AddValues(*val1, *val2)
				return []Attribute{*screenSize}
			},
			validate: func(attributes []Attribute) {
				err := s.validator.Struct(&attributes[0])
				s.NoError(err)
			},
			expectError: false,
		},
		{
			name: "validates warranty period attribute",
			setup: func() []Attribute {
				warranty, _ := NewAttribute("thoi_han_bao_hanh", "Thời hạn bảo hành")
				val1, _ := NewAttributeValue("12 Tháng")
				val2, _ := NewAttributeValue("24 Tháng")
				warranty.AddValues(*val1, *val2)
				return []Attribute{*warranty}
			},
			validate: func(attributes []Attribute) {
				err := s.validator.Struct(&attributes[0])
				s.NoError(err)
			},
			expectError: false,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			attributes := tt.setup()
			tt.validate(attributes)
		})
	}
}
