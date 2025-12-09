package domain_test

import (
	"strings"
	"testing"

	"backend/internal/domain"

	"github.com/stretchr/testify/assert"
)

func TestNewCategoryBoundaryValues(t *testing.T) {
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
		t.Run(tc.name, func(t *testing.T) {
			category, err := domain.NewCategory(tc.input)

			if tc.expectErr {
				assert.Error(t, err, tc.name)
				assert.Nil(t, category, tc.name)
			} else {
				assert.NoError(t, err, tc.name)
				assert.NotNil(t, category, tc.name)
				assert.Equal(t, tc.input, category.Name, tc.name)
				assert.NotNil(t, category.ID, tc.name)
				assert.NotNil(t, category.CreatedAt, tc.name)
				assert.NotNil(t, category.UpdatedAt, tc.name)
			}
		})
	}
}

func TestCategoryUpdateBoundaryValues(t *testing.T) {
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
		t.Run(tc.name, func(t *testing.T) {
			category, err := domain.NewCategory(tc.initialName)
			assert.NoError(t, err, "failed to create initial category")
			assert.NotNil(t, category, "category should not be nil")

			originalUpdatedAt := category.UpdatedAt

			err = category.Update(tc.updateName)

			if tc.expectErr {
				assert.Error(t, err, tc.name)
			} else {
				assert.NoError(t, err, tc.name)
			}

			assert.Equal(t, tc.expectedResult, category.Name, tc.name)

			if tc.expectOk && tc.updateName != "" && tc.updateName != tc.initialName {
				assert.True(t, category.UpdatedAt.After(originalUpdatedAt), "UpdatedAt should be updated")
			} else if !tc.expectErr {
				assert.Equal(t, originalUpdatedAt, category.UpdatedAt, "UpdatedAt should not change")
			}
		})
	}
}
