package domain

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
)

type AttributeTestSuite struct {
	suite.Suite
}

func TestAttributeTestSuite(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(AttributeTestSuite))
}

func (s *AttributeTestSuite) TestNewAttribute() {
	tests := []struct {
		name     string
		code     string
		given    string
		validate func(*Attribute, error)
	}{
		{
			name:  "creates attribute successfully",
			code:  "color",
			given: "Color",
			validate: func(attr *Attribute, err error) {
				s.Require().NoError(err)
				s.NotNil(attr)
				s.NotEqual(uuid.Nil, attr.ID)
				s.Equal("color", attr.Code)
				s.Equal("Color", attr.Name)
				s.NotNil(attr.Values)
				s.Empty(attr.Values)
				s.Nil(attr.DeletedAt)
			},
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			attr, err := NewAttribute(tt.code, tt.given)
			tt.validate(attr, err)
		})
	}
}

func (s *AttributeTestSuite) TestNewAttributeValue() {
	tests := []struct {
		name     string
		value    string
		validate func(*AttributeValue, error)
	}{
		{
			name:  "creates attribute value successfully",
			value: "Red",
			validate: func(attrVal *AttributeValue, err error) {
				s.Require().NoError(err)
				s.NotNil(attrVal)
				s.NotEqual(uuid.Nil, attrVal.ID)
				s.Equal("Red", attrVal.Value)
				s.Nil(attrVal.DeletedAt)
			},
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			attrVal, err := NewAttributeValue(tt.value)
			tt.validate(attrVal, err)
		})
	}
}

func (s *AttributeTestSuite) TestAttribute_Update() {
	tests := []struct {
		name     string
		setup    func() (*Attribute, *string)
		validate func(*Attribute, string)
	}{
		{
			name: "updates name when provided",
			setup: func() (*Attribute, *string) {
				attr, _ := NewAttribute("color", "Color")
				newName := "Updated Color"
				return attr, &newName
			},
			validate: func(attr *Attribute, expected string) {
				s.Equal(expected, attr.Name)
			},
		},
		{
			name: "does not update name when nil",
			setup: func() (*Attribute, *string) {
				attr, _ := NewAttribute("color", "Color")
				return attr, nil
			},
			validate: func(attr *Attribute, expected string) {
				s.Equal("Color", attr.Name)
			},
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			attr, newName := tt.setup()
			originalName := attr.Name
			attr.Update(newName)
			expected := originalName
			if newName != nil {
				expected = *newName
			}
			tt.validate(attr, expected)
		})
	}
}

func (s *AttributeTestSuite) TestAttribute_GetValueByID() {
	tests := []struct {
		name     string
		setup    func() (*Attribute, uuid.UUID)
		validate func(*AttributeValue)
	}{
		{
			name: "returns value when found",
			setup: func() (*Attribute, uuid.UUID) {
				attr, _ := NewAttribute("color", "Color")
				val1, _ := NewAttributeValue("Red")
				val2, _ := NewAttributeValue("Blue")
				attr.AddValues(*val1, *val2)
				return attr, val1.ID
			},
			validate: func(result *AttributeValue) {
				s.Require().NotNil(result)
				s.Equal("Red", result.Value)
			},
		},
		{
			name: "returns nil when value not found",
			setup: func() (*Attribute, uuid.UUID) {
				attr, _ := NewAttribute("color", "Color")
				val1, _ := NewAttributeValue("Red")
				attr.AddValues(*val1)
				randomID := uuid.Must(uuid.NewV7())
				return attr, randomID
			},
			validate: func(result *AttributeValue) {
				s.Nil(result)
			},
		},
		{
			name: "returns nil when values is nil",
			setup: func() (*Attribute, uuid.UUID) {
				attr := &Attribute{
					ID:     uuid.Must(uuid.NewV7()),
					Code:   "color",
					Name:   "Color",
					Values: nil,
				}
				randomID := uuid.Must(uuid.NewV7())
				return attr, randomID
			},
			validate: func(result *AttributeValue) {
				s.Nil(result)
			},
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			attr, id := tt.setup()
			result := attr.GetValueByID(id)
			tt.validate(result)
		})
	}
}

func (s *AttributeTestSuite) TestAttribute_AddValues() {
	tests := []struct {
		name          string
		setup         func() (*Attribute, []AttributeValue)
		expectedCount int
	}{
		{
			name: "adds single value",
			setup: func() (*Attribute, []AttributeValue) {
				attr, _ := NewAttribute("color", "Color")
				val, _ := NewAttributeValue("Red")
				return attr, []AttributeValue{*val}
			},
			expectedCount: 1,
		},
		{
			name: "adds multiple values",
			setup: func() (*Attribute, []AttributeValue) {
				attr, _ := NewAttribute("color", "Color")
				val1, _ := NewAttributeValue("Red")
				val2, _ := NewAttributeValue("Blue")
				val3, _ := NewAttributeValue("Green")
				return attr, []AttributeValue{*val1, *val2, *val3}
			},
			expectedCount: 3,
		},
		{
			name: "appends to existing values",
			setup: func() (*Attribute, []AttributeValue) {
				attr, _ := NewAttribute("color", "Color")
				val1, _ := NewAttributeValue("Red")
				attr.AddValues(*val1)
				val2, _ := NewAttributeValue("Blue")
				return attr, []AttributeValue{*val2}
			},
			expectedCount: 2,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			attr, values := tt.setup()
			attr.AddValues(values...)
			s.Len(attr.Values, tt.expectedCount)
		})
	}
}

func (s *AttributeTestSuite) TestAttribute_UpdateValue() {
	tests := []struct {
		name        string
		setup       func() (*Attribute, uuid.UUID, *string)
		expectError bool
		errorType   error
		validate    func(*Attribute)
	}{
		{
			name: "updates value successfully",
			setup: func() (*Attribute, uuid.UUID, *string) {
				attr, _ := NewAttribute("color", "Color")
				val, _ := NewAttributeValue("Red")
				attr.AddValues(*val)
				newValue := "Crimson Red"
				return attr, val.ID, &newValue
			},
			expectError: false,
			validate: func(attr *Attribute) {
				s.Equal("Crimson Red", attr.Values[0].Value)
			},
		},
		{
			name: "does not update when value is nil",
			setup: func() (*Attribute, uuid.UUID, *string) {
				attr, _ := NewAttribute("color", "Color")
				val, _ := NewAttributeValue("Red")
				attr.AddValues(*val)
				return attr, val.ID, nil
			},
			expectError: false,
			validate: func(attr *Attribute) {
				s.Equal("Red", attr.Values[0].Value)
			},
		},
		{
			name: "returns error when value not found",
			setup: func() (*Attribute, uuid.UUID, *string) {
				attr, _ := NewAttribute("color", "Color")
				val, _ := NewAttributeValue("Red")
				attr.AddValues(*val)
				randomID := uuid.Must(uuid.NewV7())
				newValue := "Blue"
				return attr, randomID, &newValue
			},
			expectError: true,
			errorType:   ErrNotFound,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			attr, id, newValue := tt.setup()
			err := attr.UpdateValue(id, newValue)

			if tt.expectError {
				s.Error(err)
				if tt.errorType != nil {
					s.ErrorIs(err, tt.errorType)
				}
			} else {
				s.NoError(err)
				if tt.validate != nil {
					tt.validate(attr)
				}
			}
		})
	}
}

func (s *AttributeTestSuite) TestAttribute_Remove() {
	tests := []struct {
		name     string
		setup    func() *Attribute
		validate func(*Attribute, time.Time, time.Time)
	}{
		{
			name: "sets DeletedAt on attribute and all values",
			setup: func() *Attribute {
				attr, _ := NewAttribute("color", "Color")
				val1, _ := NewAttributeValue("Red")
				val2, _ := NewAttributeValue("Blue")
				attr.AddValues(*val1, *val2)
				return attr
			},
			validate: func(attr *Attribute, beforeRemove, afterRemove time.Time) {
				s.Require().NotNil(attr.DeletedAt)
				s.True(attr.DeletedAt.After(beforeRemove) || attr.DeletedAt.Equal(beforeRemove))
				s.True(attr.DeletedAt.Before(afterRemove) || attr.DeletedAt.Equal(afterRemove))

				for _, val := range attr.Values {
					s.Require().NotNil(val.DeletedAt)
					s.True(val.DeletedAt.After(beforeRemove) || val.DeletedAt.Equal(beforeRemove))
					s.True(val.DeletedAt.Before(afterRemove) || val.DeletedAt.Equal(afterRemove))
				}
			},
		},
		{
			name: "does not update DeletedAt if already set",
			setup: func() *Attribute {
				attr, _ := NewAttribute("color", "Color")
				val, _ := NewAttributeValue("Red")
				attr.AddValues(*val)
				attr.Remove()
				time.Sleep(10 * time.Millisecond)
				return attr
			},
			validate: func(attr *Attribute, beforeRemove, afterRemove time.Time) {
				firstDeletedAt := *attr.DeletedAt
				firstValueDeletedAt := *attr.Values[0].DeletedAt
				attr.Remove()
				s.Equal(firstDeletedAt, *attr.DeletedAt)
				s.Equal(firstValueDeletedAt, *attr.Values[0].DeletedAt)
			},
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			attr := tt.setup()
			beforeRemove := time.Now()
			attr.Remove()
			afterRemove := time.Now()
			tt.validate(attr, beforeRemove, afterRemove)
		})
	}
}

func (s *AttributeTestSuite) TestAttribute_RemoveValue() {
	tests := []struct {
		name        string
		setup       func() (*Attribute, uuid.UUID)
		expectError bool
		errorType   error
		validate    func(*Attribute)
	}{
		{
			name: "removes value successfully",
			setup: func() (*Attribute, uuid.UUID) {
				attr, _ := NewAttribute("color", "Color")
				val1, _ := NewAttributeValue("Red")
				val2, _ := NewAttributeValue("Blue")
				val3, _ := NewAttributeValue("Green")
				attr.AddValues(*val1, *val2, *val3)
				return attr, val2.ID
			},
			expectError: false,
			validate: func(attr *Attribute) {
				s.Len(attr.Values, 2)
			},
		},
		{
			name: "returns error when attribute is nil",
			setup: func() (*Attribute, uuid.UUID) {
				randomID := uuid.Must(uuid.NewV7())
				return nil, randomID
			},
			expectError: true,
			errorType:   ErrInvalid,
		},
		{
			name: "does nothing when value not found",
			setup: func() (*Attribute, uuid.UUID) {
				attr, _ := NewAttribute("color", "Color")
				val, _ := NewAttributeValue("Red")
				attr.AddValues(*val)
				randomID := uuid.Must(uuid.NewV7())
				return attr, randomID
			},
			expectError: false,
			validate: func(attr *Attribute) {
				s.Len(attr.Values, 1)
			},
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			attr, id := tt.setup()
			err := attr.RemoveValue(id)

			if tt.expectError {
				s.Error(err)
				if tt.errorType != nil {
					s.ErrorIs(err, tt.errorType)
				}
			} else {
				s.NoError(err)
				if tt.validate != nil {
					tt.validate(attr)
				}
			}
		})
	}
}

func (s *AttributeTestSuite) TestValidateAttributeValueUniqueness() {
	tests := []struct {
		name        string
		setup       func() *Attribute
		expectError bool
		errorMsg    string
	}{
		{
			name: "returns nil for unique values",
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
			name: "returns error for duplicate values",
			setup: func() *Attribute {
				attr, _ := NewAttribute("color", "Color")
				val1, _ := NewAttributeValue("Red")
				val2, _ := NewAttributeValue("Red")
				attr.AddValues(*val1, *val2)
				return attr
			},
			expectError: true,
			errorMsg:    "duplicate attribute value found",
		},
		{
			name: "returns error for duplicate IDs",
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
			errorMsg:    "duplicate attribute value ID found",
		},
		{
			name: "returns nil for empty values",
			setup: func() *Attribute {
				attr, _ := NewAttribute("color", "Color")
				return attr
			},
			expectError: false,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			attr := tt.setup()
			err := ValidateAttributeValueUniqueness(attr)

			if tt.expectError {
				s.Error(err)
				s.Contains(err.Error(), tt.errorMsg)
			} else {
				s.NoError(err)
			}
		})
	}
}

func (s *AttributeTestSuite) TestValidateUniqueAttributeIDs() {
	tests := []struct {
		name        string
		setup       func() []Attribute
		expectError bool
		errorMsg    string
	}{
		{
			name: "returns nil for unique IDs",
			setup: func() []Attribute {
				attr1, _ := NewAttribute("color", "Color")
				attr2, _ := NewAttribute("size", "Size")
				attr3, _ := NewAttribute("brand", "Brand")
				return []Attribute{*attr1, *attr2, *attr3}
			},
			expectError: false,
		},
		{
			name: "returns error for duplicate IDs",
			setup: func() []Attribute {
				attr1, _ := NewAttribute("color", "Color")
				attr2 := &Attribute{
					ID:   attr1.ID,
					Code: "size",
					Name: "Size",
				}
				return []Attribute{*attr1, *attr2}
			},
			expectError: true,
			errorMsg:    "duplicate attribute ID found",
		},
		{
			name: "returns nil for empty slice",
			setup: func() []Attribute {
				return []Attribute{}
			},
			expectError: false,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			attributes := tt.setup()
			err := ValidateUniqueAttributeIDs(attributes)

			if tt.expectError {
				s.Error(err)
				s.Contains(err.Error(), tt.errorMsg)
			} else {
				s.NoError(err)
			}
		})
	}
}

func (s *AttributeTestSuite) TestValidateUniqueAttributeCodes() {
	tests := []struct {
		name        string
		setup       func() []Attribute
		expectError bool
		errorMsg    string
	}{
		{
			name: "returns nil for unique codes",
			setup: func() []Attribute {
				attr1, _ := NewAttribute("color", "Color")
				attr2, _ := NewAttribute("size", "Size")
				attr3, _ := NewAttribute("brand", "Brand")
				return []Attribute{*attr1, *attr2, *attr3}
			},
			expectError: false,
		},
		{
			name: "returns error for duplicate codes",
			setup: func() []Attribute {
				attr1, _ := NewAttribute("color", "Color")
				attr2, _ := NewAttribute("color", "Another Color")
				return []Attribute{*attr1, *attr2}
			},
			expectError: true,
			errorMsg:    "duplicate attribute code found",
		},
		{
			name: "returns nil for empty slice",
			setup: func() []Attribute {
				return []Attribute{}
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
				s.Contains(err.Error(), tt.errorMsg)
			} else {
				s.NoError(err)
			}
		})
	}
}
