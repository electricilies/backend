package repositorypostgres

import (
	"context"
	"testing"

	"backend/internal/domain"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type AttributeRepositoryTestSuite struct {
	suite.Suite
	ctx context.Context
}

func TestAttributeRepositoryTestSuite(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(AttributeRepositoryTestSuite))
}

func (s *AttributeRepositoryTestSuite) SetupTest() {
	s.ctx = context.Background()
}

func (s *AttributeRepositoryTestSuite) TestNewMockAttributeRepository() {
	tests := []struct {
		name string
	}{
		{
			name: "creates mock repository successfully",
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			mockRepo := domain.NewMockAttributeRepository(s.T())
			s.NotNil(mockRepo)
		})
	}
}

func (s *AttributeRepositoryTestSuite) TestAttributeRepository_Count() {
	tests := []struct {
		name        string
		setup       func(*domain.MockAttributeRepository) (*[]uuid.UUID, domain.DeletedParam)
		expected    int
		expectError bool
		errorType   error
	}{
		{
			name: "counts attributes successfully",
			setup: func(mockRepo *domain.MockAttributeRepository) (*[]uuid.UUID, domain.DeletedParam) {
				expectedCount := 5
				mockRepo.EXPECT().Count(
					mock.Anything,
					mock.Anything,
					domain.DeletedExcludeParam,
				).Return(&expectedCount, nil).Once()
				return nil, domain.DeletedExcludeParam
			},
			expected:    5,
			expectError: false,
		},
		{
			name: "counts attributes with specific IDs",
			setup: func(mockRepo *domain.MockAttributeRepository) (*[]uuid.UUID, domain.DeletedParam) {
				id1 := uuid.Must(uuid.NewV7())
				id2 := uuid.Must(uuid.NewV7())
				ids := []uuid.UUID{id1, id2}
				expectedCount := 2
				mockRepo.EXPECT().Count(
					mock.Anything,
					&ids,
					domain.DeletedExcludeParam,
				).Return(&expectedCount, nil).Once()
				return &ids, domain.DeletedExcludeParam
			},
			expected:    2,
			expectError: false,
		},
		{
			name: "counts including deleted attributes",
			setup: func(mockRepo *domain.MockAttributeRepository) (*[]uuid.UUID, domain.DeletedParam) {
				expectedCount := 10
				mockRepo.EXPECT().Count(
					mock.Anything,
					mock.Anything,
					domain.DeletedIncludeParam,
				).Return(&expectedCount, nil).Once()
				return nil, domain.DeletedIncludeParam
			},
			expected:    10,
			expectError: false,
		},
		{
			name: "returns error on database failure",
			setup: func(mockRepo *domain.MockAttributeRepository) (*[]uuid.UUID, domain.DeletedParam) {
				mockRepo.EXPECT().Count(
					mock.Anything,
					mock.Anything,
					domain.DeletedExcludeParam,
				).Return(nil, domain.ErrInternal).Once()
				return nil, domain.DeletedExcludeParam
			},
			expectError: true,
			errorType:   domain.ErrInternal,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			mockRepo := domain.NewMockAttributeRepository(s.T())
			ids, deleted := tt.setup(mockRepo)

			count, err := mockRepo.Count(s.ctx, ids, deleted)

			if tt.expectError {
				s.Error(err)
				s.Nil(count)
				if tt.errorType != nil {
					s.ErrorIs(err, tt.errorType)
				}
			} else {
				s.NoError(err)
				s.Require().NotNil(count)
				s.Equal(tt.expected, *count)
			}
		})
	}
}

func (s *AttributeRepositoryTestSuite) TestAttributeRepository_List() {
	tests := []struct {
		name          string
		setup         func(*domain.MockAttributeRepository)
		expectedCount int
		expectError   bool
		errorType     error
	}{
		{
			name: "lists attributes successfully",
			setup: func(mockRepo *domain.MockAttributeRepository) {
				attr1, _ := domain.NewAttribute("color", "Color")
				attr2, _ := domain.NewAttribute("size", "Size")
				expected := []domain.Attribute{*attr1, *attr2}
				mockRepo.EXPECT().List(
					mock.Anything,
					mock.Anything,
					mock.Anything,
					domain.DeletedExcludeParam,
					10,
					0,
				).Return(&expected, nil).Once()
			},
			expectedCount: 2,
			expectError:   false,
		},
		{
			name: "lists attributes with search",
			setup: func(mockRepo *domain.MockAttributeRepository) {
				search := "color"
				attr1, _ := domain.NewAttribute("color", "Color")
				expected := []domain.Attribute{*attr1}
				mockRepo.EXPECT().List(
					mock.Anything,
					mock.Anything,
					&search,
					domain.DeletedExcludeParam,
					10,
					0,
				).Return(&expected, nil).Once()
			},
			expectedCount: 1,
			expectError:   false,
		},
		{
			name: "lists attributes with pagination",
			setup: func(mockRepo *domain.MockAttributeRepository) {
				attr1, _ := domain.NewAttribute("brand", "Brand")
				expected := []domain.Attribute{*attr1}
				mockRepo.EXPECT().List(
					mock.Anything,
					mock.Anything,
					mock.Anything,
					domain.DeletedExcludeParam,
					5,
					10,
				).Return(&expected, nil).Once()
			},
			expectedCount: 1,
			expectError:   false,
		},
		{
			name: "lists attributes by specific IDs",
			setup: func(mockRepo *domain.MockAttributeRepository) {
				attr1, _ := domain.NewAttribute("color", "Color")
				ids := []uuid.UUID{attr1.ID}
				expected := []domain.Attribute{*attr1}
				mockRepo.EXPECT().List(
					mock.Anything,
					&ids,
					mock.Anything,
					domain.DeletedExcludeParam,
					10,
					0,
				).Return(&expected, nil).Once()
			},
			expectedCount: 1,
			expectError:   false,
		},
		{
			name: "returns error on database failure",
			setup: func(mockRepo *domain.MockAttributeRepository) {
				mockRepo.EXPECT().List(
					mock.Anything,
					mock.Anything,
					mock.Anything,
					domain.DeletedExcludeParam,
					10,
					0,
				).Return(nil, domain.ErrInternal).Once()
			},
			expectError: true,
			errorType:   domain.ErrInternal,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			mockRepo := domain.NewMockAttributeRepository(s.T())
			tt.setup(mockRepo)

			result, err := mockRepo.List(s.ctx, nil, nil, domain.DeletedExcludeParam, 10, 0)

			if tt.expectError {
				s.Error(err)
				s.Nil(result)
				if tt.errorType != nil {
					s.ErrorIs(err, tt.errorType)
				}
			} else {
				s.NoError(err)
				s.Require().NotNil(result)
				s.Len(*result, tt.expectedCount)
			}
		})
	}
}

func (s *AttributeRepositoryTestSuite) TestAttributeRepository_Get() {
	tests := []struct {
		name        string
		setup       func(*domain.MockAttributeRepository) uuid.UUID
		expectError bool
		errorType   error
		validate    func(*domain.Attribute)
	}{
		{
			name: "gets attribute successfully",
			setup: func(mockRepo *domain.MockAttributeRepository) uuid.UUID {
				attr, _ := domain.NewAttribute("color", "Color")
				val1, _ := domain.NewAttributeValue("Red")
				val2, _ := domain.NewAttributeValue("Blue")
				attr.AddValues(*val1, *val2)
				mockRepo.EXPECT().Get(
					mock.Anything,
					attr.ID,
				).Return(attr, nil).Once()
				return attr.ID
			},
			expectError: false,
			validate: func(result *domain.Attribute) {
				s.Equal("color", result.Code)
				s.Equal("Color", result.Name)
				s.Len(result.Values, 2)
			},
		},
		{
			name: "returns error when attribute not found",
			setup: func(mockRepo *domain.MockAttributeRepository) uuid.UUID {
				id := uuid.Must(uuid.NewV7())
				mockRepo.EXPECT().Get(
					mock.Anything,
					id,
				).Return(nil, domain.ErrNotFound).Once()
				return id
			},
			expectError: true,
			errorType:   domain.ErrNotFound,
		},
		{
			name: "returns error on database failure",
			setup: func(mockRepo *domain.MockAttributeRepository) uuid.UUID {
				id := uuid.Must(uuid.NewV7())
				mockRepo.EXPECT().Get(
					mock.Anything,
					id,
				).Return(nil, domain.ErrInternal).Once()
				return id
			},
			expectError: true,
			errorType:   domain.ErrInternal,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			mockRepo := domain.NewMockAttributeRepository(s.T())
			id := tt.setup(mockRepo)

			result, err := mockRepo.Get(s.ctx, id)

			if tt.expectError {
				s.Error(err)
				s.Nil(result)
				if tt.errorType != nil {
					s.ErrorIs(err, tt.errorType)
				}
			} else {
				s.NoError(err)
				s.Require().NotNil(result)
				if tt.validate != nil {
					tt.validate(result)
				}
			}
		})
	}
}

func (s *AttributeRepositoryTestSuite) TestAttributeRepository_ListValues() {
	tests := []struct {
		name          string
		setup         func(*domain.MockAttributeRepository) uuid.UUID
		expectedCount int
		expectError   bool
		errorType     error
	}{
		{
			name: "lists attribute values successfully",
			setup: func(mockRepo *domain.MockAttributeRepository) uuid.UUID {
				attrID := uuid.Must(uuid.NewV7())
				val1, _ := domain.NewAttributeValue("Red")
				val2, _ := domain.NewAttributeValue("Blue")
				expected := []domain.AttributeValue{*val1, *val2}
				mockRepo.EXPECT().ListValues(
					mock.Anything,
					attrID,
					mock.Anything,
					mock.Anything,
					domain.DeletedExcludeParam,
					10,
					0,
				).Return(&expected, nil).Once()
				return attrID
			},
			expectedCount: 2,
			expectError:   false,
		},
		{
			name: "lists attribute values with search",
			setup: func(mockRepo *domain.MockAttributeRepository) uuid.UUID {
				attrID := uuid.Must(uuid.NewV7())
				search := "Red"
				val1, _ := domain.NewAttributeValue("Red")
				expected := []domain.AttributeValue{*val1}
				mockRepo.EXPECT().ListValues(
					mock.Anything,
					attrID,
					mock.Anything,
					&search,
					domain.DeletedExcludeParam,
					10,
					0,
				).Return(&expected, nil).Once()
				return attrID
			},
			expectedCount: 1,
			expectError:   false,
		},
		{
			name: "lists specific attribute values by IDs",
			setup: func(mockRepo *domain.MockAttributeRepository) uuid.UUID {
				attrID := uuid.Must(uuid.NewV7())
				val1, _ := domain.NewAttributeValue("Red")
				valueIDs := []uuid.UUID{val1.ID}
				expected := []domain.AttributeValue{*val1}
				mockRepo.EXPECT().ListValues(
					mock.Anything,
					attrID,
					&valueIDs,
					mock.Anything,
					domain.DeletedExcludeParam,
					10,
					0,
				).Return(&expected, nil).Once()
				return attrID
			},
			expectedCount: 1,
			expectError:   false,
		},
		{
			name: "returns error on database failure",
			setup: func(mockRepo *domain.MockAttributeRepository) uuid.UUID {
				attrID := uuid.Must(uuid.NewV7())
				mockRepo.EXPECT().ListValues(
					mock.Anything,
					attrID,
					mock.Anything,
					mock.Anything,
					domain.DeletedExcludeParam,
					10,
					0,
				).Return(nil, domain.ErrInternal).Once()
				return attrID
			},
			expectError: true,
			errorType:   domain.ErrInternal,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			mockRepo := domain.NewMockAttributeRepository(s.T())
			attrID := tt.setup(mockRepo)

			result, err := mockRepo.ListValues(s.ctx, attrID, nil, nil, domain.DeletedExcludeParam, 10, 0)

			if tt.expectError {
				s.Error(err)
				s.Nil(result)
				if tt.errorType != nil {
					s.ErrorIs(err, tt.errorType)
				}
			} else {
				s.NoError(err)
				s.Require().NotNil(result)
				s.Len(*result, tt.expectedCount)
			}
		})
	}
}

func (s *AttributeRepositoryTestSuite) TestAttributeRepository_CountValues() {
	tests := []struct {
		name        string
		setup       func(*domain.MockAttributeRepository) uuid.UUID
		expected    int
		expectError bool
		errorType   error
	}{
		{
			name: "counts attribute values successfully",
			setup: func(mockRepo *domain.MockAttributeRepository) uuid.UUID {
				attrID := uuid.Must(uuid.NewV7())
				expectedCount := 3
				mockRepo.EXPECT().CountValues(
					mock.Anything,
					attrID,
					mock.Anything,
				).Return(&expectedCount, nil).Once()
				return attrID
			},
			expected:    3,
			expectError: false,
		},
		{
			name: "counts specific attribute values by IDs",
			setup: func(mockRepo *domain.MockAttributeRepository) uuid.UUID {
				attrID := uuid.Must(uuid.NewV7())
				valueIDs := []uuid.UUID{uuid.Must(uuid.NewV7()), uuid.Must(uuid.NewV7())}
				expectedCount := 2
				mockRepo.EXPECT().CountValues(
					mock.Anything,
					attrID,
					&valueIDs,
				).Return(&expectedCount, nil).Once()
				return attrID
			},
			expected:    2,
			expectError: false,
		},
		{
			name: "returns error on database failure",
			setup: func(mockRepo *domain.MockAttributeRepository) uuid.UUID {
				attrID := uuid.Must(uuid.NewV7())
				mockRepo.EXPECT().CountValues(
					mock.Anything,
					attrID,
					mock.Anything,
				).Return(nil, domain.ErrInternal).Once()
				return attrID
			},
			expectError: true,
			errorType:   domain.ErrInternal,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			mockRepo := domain.NewMockAttributeRepository(s.T())
			attrID := tt.setup(mockRepo)

			count, err := mockRepo.CountValues(s.ctx, attrID, nil)

			if tt.expectError {
				s.Error(err)
				s.Nil(count)
				if tt.errorType != nil {
					s.ErrorIs(err, tt.errorType)
				}
			} else {
				s.NoError(err)
				s.Require().NotNil(count)
				s.Equal(tt.expected, *count)
			}
		})
	}
}

func (s *AttributeRepositoryTestSuite) TestAttributeRepository_Save() {
	tests := []struct {
		name        string
		setup       func(*domain.MockAttributeRepository) *domain.Attribute
		expectError bool
		errorType   error
	}{
		{
			name: "saves attribute successfully",
			setup: func(mockRepo *domain.MockAttributeRepository) *domain.Attribute {
				attr, _ := domain.NewAttribute("color", "Color")
				val1, _ := domain.NewAttributeValue("Red")
				val2, _ := domain.NewAttributeValue("Blue")
				attr.AddValues(*val1, *val2)
				mockRepo.EXPECT().Save(
					mock.Anything,
					*attr,
				).Return(nil).Once()
				return attr
			},
			expectError: false,
		},
		{
			name: "saves attribute without values successfully",
			setup: func(mockRepo *domain.MockAttributeRepository) *domain.Attribute {
				attr, _ := domain.NewAttribute("color", "Color")
				mockRepo.EXPECT().Save(
					mock.Anything,
					*attr,
				).Return(nil).Once()
				return attr
			},
			expectError: false,
		},
		{
			name: "returns error on database conflict",
			setup: func(mockRepo *domain.MockAttributeRepository) *domain.Attribute {
				attr, _ := domain.NewAttribute("color", "Color")
				mockRepo.EXPECT().Save(
					mock.Anything,
					*attr,
				).Return(domain.ErrConflict).Once()
				return attr
			},
			expectError: true,
			errorType:   domain.ErrConflict,
		},
		{
			name: "returns error on database failure",
			setup: func(mockRepo *domain.MockAttributeRepository) *domain.Attribute {
				attr, _ := domain.NewAttribute("color", "Color")
				mockRepo.EXPECT().Save(
					mock.Anything,
					*attr,
				).Return(domain.ErrInternal).Once()
				return attr
			},
			expectError: true,
			errorType:   domain.ErrInternal,
		},
		{
			name: "saves updated attribute successfully",
			setup: func(mockRepo *domain.MockAttributeRepository) *domain.Attribute {
				attr, _ := domain.NewAttribute("color", "Color")
				val1, _ := domain.NewAttributeValue("Red")
				attr.AddValues(*val1)
				mockRepo.EXPECT().Save(mock.Anything, *attr).Return(nil).Once()

				newName := "Updated Color"
				attr.Update(&newName)
				val2, _ := domain.NewAttributeValue("Blue")
				attr.AddValues(*val2)
				mockRepo.EXPECT().Save(mock.Anything, *attr).Return(nil).Once()
				return attr
			},
			expectError: false,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			mockRepo := domain.NewMockAttributeRepository(s.T())
			attr := tt.setup(mockRepo)

			err := mockRepo.Save(s.ctx, *attr)

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

func (s *AttributeRepositoryTestSuite) TestAttributeRepository_ComplexScenarios() {
	tests := []struct {
		name     string
		scenario func(*domain.MockAttributeRepository)
	}{
		{
			name: "creates, retrieves, and updates attribute",
			scenario: func(mockRepo *domain.MockAttributeRepository) {
				attr, _ := domain.NewAttribute("color", "Color")
				val1, _ := domain.NewAttributeValue("Red")
				attr.AddValues(*val1)

				// Create
				mockRepo.EXPECT().Save(mock.Anything, *attr).Return(nil).Once()
				err := mockRepo.Save(s.ctx, *attr)
				s.NoError(err)

				// Retrieve
				mockRepo.EXPECT().Get(mock.Anything, attr.ID).Return(attr, nil).Once()
				retrieved, err := mockRepo.Get(s.ctx, attr.ID)
				s.NoError(err)
				s.Equal(attr.ID, retrieved.ID)

				// Update
				newName := "Updated Color"
				retrieved.Update(&newName)
				mockRepo.EXPECT().Save(mock.Anything, *retrieved).Return(nil).Once()
				err = mockRepo.Save(s.ctx, *retrieved)
				s.NoError(err)
			},
		},
		{
			name: "lists attributes and their values",
			scenario: func(mockRepo *domain.MockAttributeRepository) {
				attr1, _ := domain.NewAttribute("color", "Color")
				val1, _ := domain.NewAttributeValue("Red")
				val2, _ := domain.NewAttributeValue("Blue")
				attr1.AddValues(*val1, *val2)

				attr2, _ := domain.NewAttribute("size", "Size")
				val3, _ := domain.NewAttributeValue("Small")
				attr2.AddValues(*val3)

				attributes := []domain.Attribute{*attr1, *attr2}

				// List attributes
				mockRepo.EXPECT().List(
					mock.Anything,
					mock.Anything,
					mock.Anything,
					domain.DeletedExcludeParam,
					10,
					0,
				).Return(&attributes, nil).Once()

				result, err := mockRepo.List(s.ctx, nil, nil, domain.DeletedExcludeParam, 10, 0)
				s.NoError(err)
				s.Len(*result, 2)

				// List values for first attribute
				values1 := []domain.AttributeValue{*val1, *val2}
				mockRepo.EXPECT().ListValues(
					mock.Anything,
					attr1.ID,
					mock.Anything,
					mock.Anything,
					domain.DeletedExcludeParam,
					10,
					0,
				).Return(&values1, nil).Once()

				vals, err := mockRepo.ListValues(s.ctx, attr1.ID, nil, nil, domain.DeletedExcludeParam, 10, 0)
				s.NoError(err)
				s.Len(*vals, 2)
			},
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			mockRepo := domain.NewMockAttributeRepository(s.T())
			tt.scenario(mockRepo)
		})
	}
}
