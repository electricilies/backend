package domain_test

import (
	"strings"
	"testing"

	"backend/internal/domain"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/suite"
)

type CategoryTestSuite struct {
	suite.Suite
	validate *validator.Validate
}

func (s *CategoryTestSuite) SetupSuite() {
	s.validate = validator.New(validator.WithRequiredStructEnabled())
}

func (s *CategoryTestSuite) TestCategoryCreationBoundaryValues() {
	testcases := []struct {
		name      string
		input     string
		expectOk  bool
		expectErr bool
	}{
		{
			name:      "name length 1 (min - 1)",
			input:     "a",
			expectOk:  false,
			expectErr: true,
		},
		{
			name:      "name length 2 (min)",
			input:     "ab",
			expectOk:  true,
			expectErr: false,
		},
		{
			name:      "name length 3 (min + 1)",
			input:     "abc",
			expectOk:  true,
			expectErr: false,
		},
		{
			name:      "name length 50",
			input:     strings.Repeat("a", 50),
			expectOk:  true,
			expectErr: false,
		},
		{
			name:      "name length 99 (max - 1)",
			input:     strings.Repeat("a", 99),
			expectOk:  true,
			expectErr: false,
		},
		{
			name:      "name length 100 (max)",
			input:     strings.Repeat("a", 100),
			expectOk:  true,
			expectErr: false,
		},
		{
			name:      "name length 101 (max + 1)",
			input:     strings.Repeat("a", 101),
			expectOk:  false,
			expectErr: true,
		},
		{
			name:      "empty name",
			input:     "",
			expectOk:  false,
			expectErr: true,
		},
	}

	for _, tc := range testcases {
		s.Run(tc.name, func() {
			category, err := domain.NewCategory(tc.input)

			s.NoError(err, tc.name)
			s.NotNil(category, tc.name)
			s.Equal(tc.input, category.Name, tc.name)
			s.NotNil(category.ID, tc.name)
			s.NotNil(category.CreatedAt, tc.name)
			s.NotNil(category.UpdatedAt, tc.name)

			validationErr := s.validate.Struct(category)
			if tc.expectErr {
				s.Error(validationErr, tc.name)
			} else {
				s.NoError(validationErr, tc.name)
			}
		})
	}
}

func (s *CategoryTestSuite) TestCategoryUpdateBoundaryValues() {
	testcases := []struct {
		name           string
		initialName    string
		updateName     string
		expectOk       bool
		expectErr      bool
		expectedResult string
	}{
		{
			name:           "update to name length 1 (min - 1)",
			initialName:    "ValidName",
			updateName:     "a",
			expectOk:       false,
			expectErr:      true,
			expectedResult: "a",
		},
		{
			name:           "update to name length 2 (min)",
			initialName:    "ValidName",
			updateName:     "ab",
			expectOk:       true,
			expectErr:      false,
			expectedResult: "ab",
		},
		{
			name:           "update to name length 3 (min + 1)",
			initialName:    "ValidName",
			updateName:     "abc",
			expectOk:       true,
			expectErr:      false,
			expectedResult: "abc",
		},
		{
			name:           "update to name length 99 (max - 1)",
			initialName:    "ValidName",
			updateName:     strings.Repeat("a", 99),
			expectOk:       true,
			expectErr:      false,
			expectedResult: strings.Repeat("a", 99),
		},
		{
			name:           "update to name length 100 (max)",
			initialName:    "ValidName",
			updateName:     strings.Repeat("a", 100),
			expectOk:       true,
			expectErr:      false,
			expectedResult: strings.Repeat("a", 100),
		},
		{
			name:           "update to name length 101 (max + 1)",
			initialName:    "ValidName",
			updateName:     strings.Repeat("a", 101),
			expectOk:       false,
			expectErr:      true,
			expectedResult: strings.Repeat("a", 101),
		},
		{
			name:           "update with empty name (no change)",
			initialName:    "ValidName",
			updateName:     "",
			expectOk:       true,
			expectErr:      false,
			expectedResult: "ValidName",
		},
		{
			name:           "update with same name (no change)",
			initialName:    "ValidName",
			updateName:     "ValidName",
			expectOk:       true,
			expectErr:      false,
			expectedResult: "ValidName",
		},
	}

	for _, tc := range testcases {
		s.Run(tc.name, func() {
			category, err := domain.NewCategory(tc.initialName)
			s.Require().NoError(err, "failed to create initial category")
			s.Require().NotNil(category, "category should not be nil")

			originalUpdatedAt := category.UpdatedAt

			// Update never returns error (validation moved to service layer)
			err = category.Update(tc.updateName)
			s.NoError(err, tc.name)

			s.Equal(tc.expectedResult, category.Name, tc.name)

			// Validation happens at service layer (application layer)
			validationErr := s.validate.Struct(category)
			if tc.expectErr {
				s.Error(validationErr, "category should fail validation when expected")
			} else {
				s.NoError(validationErr, "category should pass validation")
			}

			if tc.expectOk && tc.updateName != "" && tc.updateName != tc.initialName {
				s.True(category.UpdatedAt.After(originalUpdatedAt), "UpdatedAt should be updated")
			} else if !tc.expectErr {
				s.Equal(originalUpdatedAt, category.UpdatedAt, "UpdatedAt should not change")
			}
		})
	}
}

func TestCategory(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(CategoryTestSuite))
}
