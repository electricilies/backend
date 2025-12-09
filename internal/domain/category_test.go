// vim: tabstop=4:

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
		// Min - 1: 1 character (below minimum)
		{
			name:      "name length 1 (min - 1)",
			input:     "a",
			expectOk:  false,
			expectErr: true,
		},
		// Min: 2 characters (minimum valid)
		{
			name:      "name length 2 (min)",
			input:     "ab",
			expectOk:  true,
			expectErr: false,
		},
		// Min + 1: 3 characters (above minimum)
		{
			name:      "name length 3 (min + 1)",
			input:     "abc",
			expectOk:  true,
			expectErr: false,
		},
		// Max - 1: 99 characters (below maximum)
		{
			name:      "name length 99 (max - 1)",
			input:     strings.Repeat("a", 99),
			expectOk:  true,
			expectErr: false,
		},
		// Max: 100 characters (maximum valid)
		{
			name:      "name length 100 (max)",
			input:     strings.Repeat("a", 100),
			expectOk:  true,
			expectErr: false,
		},
		// Max + 1: 101 characters (above maximum)
		{
			name:      "name length 101 (max + 1)",
			input:     strings.Repeat("a", 101),
			expectOk:  false,
			expectErr: true,
		},
		// Empty string test
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
		// Min - 1: 1 character (below minimum)
		{
			name:           "update to name length 1 (min - 1)",
			initialName:    "ValidName",
			updateName:     "a",
			expectOk:       false,
			expectErr:      true,
			expectedResult: "a",
		},
		// Min: 2 characters (minimum valid)
		{
			name:           "update to name length 2 (min)",
			initialName:    "ValidName",
			updateName:     "ab",
			expectOk:       true,
			expectErr:      false,
			expectedResult: "ab",
		},
		// Min + 1: 3 characters (above minimum)
		{
			name:           "update to name length 3 (min + 1)",
			initialName:    "ValidName",
			updateName:     "abc",
			expectOk:       true,
			expectErr:      false,
			expectedResult: "abc",
		},
		// Max - 1: 99 characters (below maximum)
		{
			name:           "update to name length 99 (max - 1)",
			initialName:    "ValidName",
			updateName:     strings.Repeat("a", 99),
			expectOk:       true,
			expectErr:      false,
			expectedResult: strings.Repeat("a", 99),
		},
		// Max: 100 characters (maximum valid)
		{
			name:           "update to name length 100 (max)",
			initialName:    "ValidName",
			updateName:     strings.Repeat("a", 100),
			expectOk:       true,
			expectErr:      false,
			expectedResult: strings.Repeat("a", 100),
		},
		// Max + 1: 101 characters (above maximum)
		{
			name:           "update to name length 101 (max + 1)",
			initialName:    "ValidName",
			updateName:     strings.Repeat("a", 101),
			expectOk:       false,
			expectErr:      true,
			expectedResult: strings.Repeat("a", 101),
		},
		// Empty string test (should not update)
		{
			name:           "update with empty name (no change)",
			initialName:    "ValidName",
			updateName:     "",
			expectOk:       true,
			expectErr:      false,
			expectedResult: "ValidName",
		},
		// Same name test (should not update)
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

			// If update was successful and name changed, UpdatedAt should be updated
			if tc.expectOk && tc.updateName != "" && tc.updateName != tc.initialName {
				assert.True(t, category.UpdatedAt.After(originalUpdatedAt), "UpdatedAt should be updated")
			} else if !tc.expectErr {
				// If no error but no change, UpdatedAt should remain the same
				assert.Equal(t, originalUpdatedAt, category.UpdatedAt, "UpdatedAt should not change")
			}
		})
	}
}
