package domain_test

import (
	"testing"

	"backend/internal/domain"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
)

type AttributeValidatorTestSuite struct {
	suite.Suite
	validator *validator.Validate
}

func (s *AttributeValidatorTestSuite) SetupTest() {
	s.validator = validator.New()
	err := domain.RegisterAttributeValidators(s.validator)
	s.Require().NoError(err)
}

func (s *AttributeValidatorTestSuite) TestValidateUniqueAttributeValues() {
	testcases := []struct {
		name     string
		values   []domain.AttributeValue
		expectOk bool
	}{
		{
			name: "valid unique values",
			values: []domain.AttributeValue{
				{ID: uuid.New(), Value: "Red"},
				{ID: uuid.New(), Value: "Blue"},
			},
			expectOk: false,
		},
		{
			name: "duplicate values",
			values: []domain.AttributeValue{
				{ID: uuid.New(), Value: "Red"},
				{ID: uuid.New(), Value: "Red"},
			},
			expectOk: true,
		},
		{
			name: "duplicate IDs",
			values: []domain.AttributeValue{
				{ID: uuid.MustParse("123e4567-e89b-12d3-a456-426614174000"), Value: "Red"},
				{ID: uuid.MustParse("123e4567-e89b-12d3-a456-426614174000"), Value: "Blue"},
			},
			expectOk: true,
		},
		{
			name:     "Omit checking",
			values:   nil,
			expectOk: false,
		},
	}
	for _, tc := range testcases {
		s.Run(tc.name, func() {
			attr := &domain.Attribute{
				ID:     uuid.New(),
				Code:   "color",
				Name:   "Color",
				Values: tc.values,
			}
			err := s.validator.Struct(attr)
			if tc.expectOk {
				s.Error(err)
			} else {
				s.NoError(err)
			}
		})
	}
}

func (s *AttributeValidatorTestSuite) TestValidateUniqueAttributeForInvalidAnnotateion() {
	attribute := &struct {
		Values []int `validate:"uniqueAttributeValues"`
	}{
		Values: []int{1, 2, 3},
	}
	err := s.validator.Struct(attribute)
	s.NoError(err)
}

func TestAttributeValidatorTestSuite(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(AttributeValidatorTestSuite))
}
