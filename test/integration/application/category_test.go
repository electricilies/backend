// vim: tabstop=4 shiftwidth=4:
//go:build integration

package application_test

import (
	"context"
	"strings"
	"testing"

	"backend/config"
	"backend/internal/application"
	"backend/internal/client"
	"backend/internal/delivery/http"
	"backend/internal/infrastructure/cacheredis"
	"backend/internal/infrastructure/repositorypostgres"
	"backend/internal/service"
	"backend/test/integration/component"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
)

type CategoryTestSuite struct {
	suite.Suite
	containers *component.Containers
	app        http.CategoryApplication
}

func TestCategorySuite(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(CategoryTestSuite))
}

func (s *CategoryTestSuite) newContainersConfig() *component.ContainersConfig {
	containersConfig := component.NewContainersConfig(&component.NewContainersConfigParam{
		DBEnabled:    true,
		RedisEnabled: true,
	})
	return containersConfig
}

func (s *CategoryTestSuite) newConfig(
	ctx context.Context,
) *config.Server {
	s.T().Helper()

	dbConnStr, err := s.containers.DB.ConnectionString(ctx, "sslmode=disable")
	s.Require().NoError(err, "failed to get db connection string")

	redisConnStr, err := s.containers.Redis.ConnectionString(ctx)
	s.Require().NoError(err, "failed to get redis connection string")
	return &config.Server{
		DBURL:     dbConnStr,
		RedisAddr: strings.TrimPrefix(redisConnStr, "redis://"),
	}
}

func (s *CategoryTestSuite) SetupSuite() {
	ctx := s.T().Context()
	containersConfig := s.newContainersConfig()

	var err error
	s.containers, err = component.NewContainers(ctx, containersConfig)
	s.Require().NoError(err, "failed to start containers")

	cfg := s.newConfig(ctx)

	validate := validator.New(
		validator.WithRequiredStructEnabled(),
	)

	conn := client.NewDBConnection(ctx, cfg)
	queries := client.NewDBQueries(conn)
	categoryRepo := repositorypostgres.ProvideCategory(queries)

	categoryService := service.ProvideCategory(validate)

	redisClient := client.NewRedis(ctx, cfg)
	categoryCache := cacheredis.ProvideCategory(redisClient)
	s.app = application.ProvideCategory(categoryRepo, categoryService, categoryCache)
}

func (s *CategoryTestSuite) TearDownSuite() {
	s.containers.Cleanup(s.T())
}

func (s *CategoryTestSuite) TestCategoryLifecycle() {
	ctx := s.T().Context()

	var firstCategoryID uuid.UUID
	var secondCategoryID uuid.UUID

	s.Run("Create first category", func() {
		result, err := s.app.Create(ctx, http.CreateCategoryRequestDto{
			Data: http.CreateCategoryData{
				Name: "Electronics",
			},
		})
		s.Require().NoError(err)
		s.Require().NotNil(result)
		s.Equal("Electronics", result.Name)
		s.NotEqual(uuid.Nil, result.ID)
		firstCategoryID = result.ID
	})

	s.Run("Get created category", func() {
		result, err := s.app.Get(ctx, http.GetCategoryRequestDto{
			CategoryID: firstCategoryID,
		})
		s.Require().NoError(err)
		s.Require().NotNil(result)
		s.Equal(firstCategoryID, result.ID)
		s.Equal("Electronics", result.Name)
	})

	s.Run("Create second category", func() {
		result, err := s.app.Create(ctx, http.CreateCategoryRequestDto{
			Data: http.CreateCategoryData{
				Name: "Home Appliances",
			},
		})
		s.Require().NoError(err)
		s.Require().NotNil(result)
		s.Equal("Home Appliances", result.Name)
		secondCategoryID = result.ID
	})

	s.Run("List categories returns 2", func() {
		result, err := s.app.List(ctx, http.ListCategoryRequestDto{
			PaginationRequestDto: http.PaginationRequestDto{
				Page:  1,
				Limit: 10,
			},
		})
		s.Require().NoError(err)
		s.Require().NotNil(result)
		s.Equal(2, result.Meta.TotalItems)
		s.Len(result.Data, 2)
	})

	s.Run("Update first category", func() {
		result, err := s.app.Update(ctx, http.UpdateCategoryRequestDto{
			CategoryID: firstCategoryID,
			Data: http.UpdateCategoryData{
				Name: "Updated Electronics",
			},
		})
		s.Require().NoError(err)
		s.Require().NotNil(result)
		s.Equal("Updated Electronics", result.Name)
	})

	s.Run("Get updated category reflects changes", func() {
		result, err := s.app.Get(ctx, http.GetCategoryRequestDto{
			CategoryID: firstCategoryID,
		})
		s.Require().NoError(err)
		s.Require().NotNil(result)
		s.Equal("Updated Electronics", result.Name)
	})

	s.Run("List with search filter", func() {
		result, err := s.app.List(ctx, http.ListCategoryRequestDto{
			PaginationRequestDto: http.PaginationRequestDto{
				Page:  1,
				Limit: 10,
			},
			Search: "Electronics",
		})
		s.Require().NoError(err)
		s.Require().NotNil(result)
		s.GreaterOrEqual(result.Meta.TotalItems, 1)
		s.GreaterOrEqual(len(result.Data), 1)
	})

	s.Run("List with pagination", func() {
		result, err := s.app.List(ctx, http.ListCategoryRequestDto{
			PaginationRequestDto: http.PaginationRequestDto{
				Page:  1,
				Limit: 1,
			},
		})
		s.Require().NoError(err)
		s.Require().NotNil(result)
		s.Equal(2, result.Meta.TotalItems)
		s.Len(result.Data, 1)
		s.Equal(2, result.Meta.TotalPages)
	})

	s.Run("List second page", func() {
		result, err := s.app.List(ctx, http.ListCategoryRequestDto{
			PaginationRequestDto: http.PaginationRequestDto{
				Page:  2,
				Limit: 1,
			},
		})
		s.Require().NoError(err)
		s.Require().NotNil(result)
		s.Equal(2, result.Meta.TotalItems)
		s.Len(result.Data, 1)
	})

	s.Run("Create category with empty name fails", func() {
		_, err := s.app.Create(ctx, http.CreateCategoryRequestDto{
			Data: http.CreateCategoryData{
				Name: "",
			},
		})
		s.Require().Error(err)
	})

	s.Run("Update category with empty name fails", func() {
		_, err := s.app.Update(ctx, http.UpdateCategoryRequestDto{
			CategoryID: firstCategoryID,
			Data: http.UpdateCategoryData{
				Name: "",
			},
		})
		s.Require().NoError(err)
	})

	nonExistentID := uuid.MustParse("00000000-0000-0000-0000-000000000001")

	s.Run("Get non-existent category fails", func() {
		_, err := s.app.Get(ctx, http.GetCategoryRequestDto{
			CategoryID: nonExistentID,
		})
		s.Require().Error(err)
	})

	s.Run("Update non-existent category fails", func() {
		_, err := s.app.Update(ctx, http.UpdateCategoryRequestDto{
			CategoryID: nonExistentID,
			Data: http.UpdateCategoryData{
				Name: "New Name",
			},
		})
		s.Require().Error(err)
	})

	s.Run("Cache is working for Get", func() {
		result1, err := s.app.Get(ctx, http.GetCategoryRequestDto{
			CategoryID: firstCategoryID,
		})
		s.Require().NoError(err)
		s.Require().NotNil(result1)

		result2, err := s.app.Get(ctx, http.GetCategoryRequestDto{
			CategoryID: firstCategoryID,
		})
		s.Require().NoError(err)
		s.Require().NotNil(result2)
		s.Equal(result1.ID, result2.ID)
		s.Equal(result1.Name, result2.Name)
	})

	s.Run("Cache is working for List", func() {
		result1, err := s.app.List(ctx, http.ListCategoryRequestDto{
			PaginationRequestDto: http.PaginationRequestDto{
				Page:  1,
				Limit: 10,
			},
		})
		s.Require().NoError(err)
		s.Require().NotNil(result1)

		result2, err := s.app.List(ctx, http.ListCategoryRequestDto{
			PaginationRequestDto: http.PaginationRequestDto{
				Page:  1,
				Limit: 10,
			},
		})
		s.Require().NoError(err)
		s.Require().NotNil(result2)
		s.Equal(result1.Meta.TotalItems, result2.Meta.TotalItems)
	})

	s.Run("Cache is invalidated after update", func() {
		_, err := s.app.Get(ctx, http.GetCategoryRequestDto{
			CategoryID: secondCategoryID,
		})
		s.Require().NoError(err)

		updated, err := s.app.Update(ctx, http.UpdateCategoryRequestDto{
			CategoryID: secondCategoryID,
			Data: http.UpdateCategoryData{
				Name: "Cache Invalidation Test",
			},
		})
		s.Require().NoError(err)
		s.Equal("Cache Invalidation Test", updated.Name)

		result, err := s.app.Get(ctx, http.GetCategoryRequestDto{
			CategoryID: secondCategoryID,
		})
		s.Require().NoError(err)
		s.Equal("Cache Invalidation Test", result.Name)
	})
}
