// vim: tabstop=4:
package domain_test

import (
	"strings"
	"testing"

	"backend/internal/domain"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
)

type AttributeTestSuite struct {
	suite.Suite
	validate *validator.Validate
}

func (s *AttributeTestSuite) SetupSuite() {
	s.validate = validator.New(validator.WithRequiredStructEnabled())
}

func (s *AttributeTestSuite) TestNewAttributeBoundaryValues() {
	s.T().Parallel()
	testcases := []struct {
		name      string
		code      string
		attrName  string
		expectOk  bool
		expectErr bool
	}{
		{
			name:      "code length 1 (min - 1)",
			code:      "a",
			attrName:  "ValidName",
			expectOk:  false,
			expectErr: true,
		},
		{
			name:      "code length 2 (min)",
			code:      "ab",
			attrName:  "ValidName",
			expectOk:  true,
			expectErr: false,
		},
		{
			name:      "code length 3 (min + 1)",
			code:      "abc",
			attrName:  "ValidName",
			expectOk:  true,
			expectErr: false,
		},
		{
			name:      "code length 50 (max)",
			code:      strings.Repeat("a", 50),
			attrName:  "ValidName",
			expectOk:  true,
			expectErr: false,
		},
		{
			name:      "code length 51 (max + 1)",
			code:      strings.Repeat("a", 51),
			attrName:  "ValidName",
			expectOk:  false,
			expectErr: true,
		},
		{
			name:      "name length 1 (min - 1)",
			code:      "validcode",
			attrName:  "a",
			expectOk:  false,
			expectErr: true,
		},
		{
			name:      "name length 2 (min)",
			code:      "validcode",
			attrName:  "ab",
			expectOk:  true,
			expectErr: false,
		},
		{
			name:      "name length 3 (min + 1)",
			code:      "validcode",
			attrName:  "abc",
			expectOk:  true,
			expectErr: false,
		},
		{
			name:      "name length 100 (max)",
			code:      "validcode",
			attrName:  strings.Repeat("a", 100),
			expectOk:  true,
			expectErr: false,
		},
		{
			name:      "name length 101 (max + 1)",
			code:      "validcode",
			attrName:  strings.Repeat("a", 101),
			expectOk:  false,
			expectErr: true,
		},
		{
			name:      "valid attribute",
			code:      "color",
			attrName:  "Color",
			expectOk:  true,
			expectErr: false,
		},
		{
			name:      "empty code",
			code:      "",
			attrName:  "ValidName",
			expectOk:  false,
			expectErr: true,
		},
		{
			name:      "empty name",
			code:      "validcode",
			attrName:  "",
			expectOk:  false,
			expectErr: true,
		},
	}

	for _, tc := range testcases {
		s.Run(tc.name, func() {
			attribute, err := domain.NewAttribute(tc.code, tc.attrName)

			s.NoError(err, tc.name)
			s.NotNil(attribute, tc.name)
			s.Equal(tc.code, attribute.Code, tc.name)
			s.Equal(tc.attrName, attribute.Name, tc.name)
			s.NotNil(attribute.ID, tc.name)
			s.NotNil(attribute.Values, tc.name)

			validationErr := s.validate.Struct(attribute)
			if tc.expectErr {
				s.Error(validationErr, tc.name)
			} else {
				s.NoError(validationErr, tc.name)
			}
		})
	}
}

func (s *AttributeTestSuite) TestNewAttributeValueBoundaryValues() {
	s.T().Parallel()
	testcases := []struct {
		name      string
		value     string
		expectOk  bool
		expectErr bool
	}{
		{
			name:      "value length 0 (empty)",
			value:     "",
			expectOk:  false,
			expectErr: true,
		},
		{
			name:      "value length 1 (min)",
			value:     "a",
			expectOk:  true,
			expectErr: false,
		},
		{
			name:      "value length 2 (min + 1)",
			value:     "ab",
			expectOk:  true,
			expectErr: false,
		},
		{
			name:      "value length 50",
			value:     strings.Repeat("a", 50),
			expectOk:  true,
			expectErr: false,
		},
		{
			name:      "value length 100 (max)",
			value:     strings.Repeat("a", 100),
			expectOk:  true,
			expectErr: false,
		},
		{
			name:      "value length 101 (max + 1)",
			value:     strings.Repeat("a", 101),
			expectOk:  false,
			expectErr: true,
		},
		{
			name:      "valid value",
			value:     "Red",
			expectOk:  true,
			expectErr: false,
		},
	}

	for _, tc := range testcases {
		s.Run(tc.name, func() {
			attrValue, err := domain.NewAttributeValue(tc.value)

			s.NoError(err, tc.name)
			s.NotNil(attrValue, tc.name)
			s.Equal(tc.value, attrValue.Value, tc.name)
			s.NotNil(attrValue.ID, tc.name)

			validationErr := s.validate.Struct(attrValue)
			if tc.expectErr {
				s.Error(validationErr, tc.name)
			} else {
				s.NoError(validationErr, tc.name)
			}
		})
	}
}

func (s *AttributeTestSuite) TestAttributeUpdate() {
	s.T().Parallel()
	testcases := []struct {
		name           string
		initialName    string
		updateName     string
		expectedResult string
	}{
		{
			name:           "update with empty name (no change)",
			initialName:    "Original Name",
			updateName:     "",
			expectedResult: "Original Name",
		},
		{
			name:           "update with same name (no change)",
			initialName:    "Original Name",
			updateName:     "Original Name",
			expectedResult: "Original Name",
		},
		{
			name:           "update with new valid name",
			initialName:    "Original Name",
			updateName:     "Updated Name",
			expectedResult: "Updated Name",
		},
		{
			name:           "update with different case",
			initialName:    "Name",
			updateName:     "name",
			expectedResult: "name",
		},
	}

	for _, tc := range testcases {
		s.Run(tc.name, func() {
			attribute, err := domain.NewAttribute("code", tc.initialName)
			s.Require().NoError(err)
			s.Require().NotNil(attribute)

			attribute.Update(tc.updateName)

			s.Equal(tc.expectedResult, attribute.Name, tc.name)
		})
	}
}

func (s *AttributeTestSuite) TestAttributeUpdateValue() {
	s.T().Parallel()
	testcases := []struct {
		name           string
		setupValues    []string
		targetIndex    int
		newValue       string
		expectErr      bool
		expectedResult string
	}{
		{
			name:           "update existing value with new value",
			setupValues:    []string{"Red", "Blue"},
			targetIndex:    0,
			newValue:       "Green",
			expectErr:      false,
			expectedResult: "Green",
		},
		{
			name:           "update value with empty string (no change)",
			setupValues:    []string{"Red", "Blue"},
			targetIndex:    0,
			newValue:       "",
			expectErr:      false,
			expectedResult: "Red",
		},
		{
			name:        "update non-existent value ID",
			setupValues: []string{"Red"},
			targetIndex: -1,
			newValue:    "Green",
			expectErr:   true,
		},
	}

	for _, tc := range testcases {
		s.Run(tc.name, func() {
			attribute, err := domain.NewAttribute("color", "Color")
			s.Require().NoError(err)

			for _, val := range tc.setupValues {
				attrVal, err := domain.NewAttributeValue(val)
				s.Require().NoError(err)
				attribute.AddValues(*attrVal)
			}

			var targetID uuid.UUID
			if tc.targetIndex == -1 {
				targetID = uuid.New()
			} else {
				targetID = attribute.Values[tc.targetIndex].ID
			}

			err = attribute.UpdateValue(targetID, tc.newValue)

			if tc.expectErr {
				s.Error(err, tc.name)
				s.Equal(domain.ErrNotFound, err, tc.name)
			} else {
				s.NoError(err, tc.name)
				if tc.targetIndex >= 0 {
					s.Equal(tc.expectedResult, attribute.Values[tc.targetIndex].Value, tc.name)
				}
			}
		})
	}
}

func (s *AttributeTestSuite) TestAttributeRemoveValue() {
	s.T().Parallel()
	testcases := []struct {
		name           string
		setupValues    []string
		targetIndex    int
		expectErr      bool
		expectedLength int
	}{
		{
			name:           "remove existing value",
			setupValues:    []string{"Red", "Blue", "Green"},
			targetIndex:    1,
			expectErr:      false,
			expectedLength: 2,
		},
		{
			name:           "remove non-existent value",
			setupValues:    []string{"Red"},
			targetIndex:    -1,
			expectErr:      false,
			expectedLength: 1,
		},
		{
			name:           "remove from single value",
			setupValues:    []string{"Red"},
			targetIndex:    0,
			expectErr:      false,
			expectedLength: 0,
		},
	}

	for _, tc := range testcases {
		s.Run(tc.name, func() {
			attribute, err := domain.NewAttribute("color", "Color")
			s.Require().NoError(err)

			for _, val := range tc.setupValues {
				attrVal, err := domain.NewAttributeValue(val)
				s.Require().NoError(err)
				attribute.AddValues(*attrVal)
			}

			var targetID uuid.UUID
			if tc.targetIndex == -1 {
				targetID = uuid.New()
			} else {
				targetID = attribute.Values[tc.targetIndex].ID
			}

			err = attribute.RemoveValue(targetID)

			if tc.expectErr {
				s.Error(err, tc.name)
			} else {
				s.NoError(err, tc.name)
				s.Len(attribute.Values, tc.expectedLength, tc.name)
			}
		})
	}
}

func TestAttribute(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(AttributeTestSuite))
}
