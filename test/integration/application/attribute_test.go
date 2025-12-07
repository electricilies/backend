// vim: tabstop=4:
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

type AttributeTestSuite struct {
	suite.Suite
	containers *component.Containers
	app        *application.Attribute
}

func TestAttributeSuite(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(AttributeTestSuite))
}

func (s *AttributeTestSuite) newContainersConfig() *component.ContainersConfig {
	containersConfig := component.NewContainersConfig(&component.NewContainersConfigParam{
		DBEnabled:    true,
		RedisEnabled: true,
	})
	return containersConfig
}

func (s *AttributeTestSuite) newConfig(
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

func (s *AttributeTestSuite) SetupSuite() {
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
	attributeRepo := repositorypostgres.ProvideAttribute(queries, conn)

	attributeService := service.ProvideAttribute(validate)

	redisClient := client.NewRedis(ctx, cfg)
	attributeCache := cacheredis.ProvideAttribute(redisClient)
	s.app = application.ProvideAttribute(attributeRepo, attributeService, attributeCache)
}

func (s *AttributeTestSuite) TearDownSuite() {
	s.containers.Cleanup(s.T())
}

func (s *AttributeTestSuite) TestAttributeLifecycle() {
	ctx := s.T().Context()

	var firstAttributeID uuid.UUID
	var secondAttributeID uuid.UUID
	var firstValueID uuid.UUID
	var secondValueID uuid.UUID

	s.Run("Create first attribute", func() {
		result, err := s.app.Create(ctx, http.CreateAttributeRequestDto{
			Data: http.CreateAttributeData{
				Code: "color",
				Name: "Color",
			},
		})
		s.Require().NoError(err)
		s.Require().NotNil(result)
		s.Equal("color", result.Code)
		s.Equal("Color", result.Name)
		s.Empty(result.Values)
		firstAttributeID = result.ID
	})

	s.Run("Get created attribute", func() {
		result, err := s.app.Get(ctx, http.GetAttributeRequestDto{
			AttributeID: firstAttributeID,
		})
		s.Require().NoError(err)
		s.Require().NotNil(result)
		s.Equal(firstAttributeID, result.ID)
		s.Equal("color", result.Code)
		s.Equal("Color", result.Name)
	})

	s.Run("Create duplicate attribute with same code fails", func() {
		_, err := s.app.Create(ctx, http.CreateAttributeRequestDto{
			Data: http.CreateAttributeData{
				Code: "color",
				Name: "Another Color",
			},
		})
		s.Require().Error(err)
	})

	s.Run("Create second attribute", func() {
		result, err := s.app.Create(ctx, http.CreateAttributeRequestDto{
			Data: http.CreateAttributeData{
				Code: "size",
				Name: "Size",
			},
		})
		s.Require().NoError(err)
		s.Require().NotNil(result)
		s.Equal("size", result.Code)
		s.Equal("Size", result.Name)
		secondAttributeID = result.ID
	})

	s.Run("List attributes returns 2", func() {
		result, err := s.app.List(ctx, http.ListAttributesRequestDto{
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

	s.Run("Update first attribute", func() {
		result, err := s.app.Update(ctx, http.UpdateAttributeRequestDto{
			AttributeID: firstAttributeID,
			Data: http.UpdateAttributeData{
				Name: "Updated Color",
			},
		})
		s.Require().NoError(err)
		s.Require().NotNil(result)
		s.Equal("Updated Color", result.Name)
		s.Equal("color", result.Code)
	})

	s.Run("Delete first attribute", func() {
		err := s.app.Delete(ctx, http.DeleteAttributeRequestDto{
			AttributeID: firstAttributeID,
		})
		s.Require().NoError(err)
	})

	s.Run("Get deleted attribute fails", func() {
		_, err := s.app.Get(ctx, http.GetAttributeRequestDto{
			AttributeID: firstAttributeID,
		})
		s.Require().Error(err)
	})

	s.Run("List attributes after delete returns 1", func() {
		result, err := s.app.List(ctx, http.ListAttributesRequestDto{
			PaginationRequestDto: http.PaginationRequestDto{
				Page:  1,
				Limit: 10,
			},
		})
		s.Require().NoError(err)
		s.Require().NotNil(result)
		s.Equal(1, result.Meta.TotalItems)
		s.Len(result.Data, 1)
		s.Equal(secondAttributeID, result.Data[0].ID)
	})

	s.Run("Create first attribute value", func() {
		result, err := s.app.CreateValue(ctx, http.CreateAttributeValueRequestDto{
			AttributeID: secondAttributeID,
			Data: http.CreateAttributeValueData{
				Value: "Small",
			},
		})
		s.Require().NoError(err)
		s.Require().NotNil(result)
		s.Equal("Small", result.Value)
		firstValueID = result.ID
	})

	s.Run("Create second attribute value", func() {
		result, err := s.app.CreateValue(ctx, http.CreateAttributeValueRequestDto{
			AttributeID: secondAttributeID,
			Data: http.CreateAttributeValueData{
				Value: "Medium",
			},
		})
		s.Require().NoError(err)
		s.Require().NotNil(result)
		s.Equal("Medium", result.Value)
		secondValueID = result.ID
	})

	s.Run("List attribute values returns 2", func() {
		result, err := s.app.ListValues(ctx, http.ListAttributeValuesRequestDto{
			PaginationRequestDto: http.PaginationRequestDto{
				Page:  1,
				Limit: 10,
			},
			AttributeID: secondAttributeID,
		})
		s.Require().NoError(err)
		s.Require().NotNil(result)
		s.Equal(2, result.Meta.TotalItems)
		s.Len(result.Data, 2)
	})

	s.Run("Update attribute value", func() {
		result, err := s.app.UpdateValue(ctx, http.UpdateAttributeValueRequestDto{
			AttributeID:      secondAttributeID,
			AttributeValueID: firstValueID,
			Data: http.UpdateAttributeValueData{
				Value: "Extra Small",
			},
		})
		s.Require().NoError(err)
		s.Require().NotNil(result)
		s.Equal("Extra Small", result.Value)
	})

	s.Run("Delete attribute value", func() {
		err := s.app.DeleteValue(ctx, http.DeleteAttributeValueRequestDto{
			AttributeID:      secondAttributeID,
			AttributeValueID: secondValueID,
		})
		s.Require().NoError(err)
	})

	s.Run("List attribute values after delete returns 1", func() {
		result, err := s.app.ListValues(ctx, http.ListAttributeValuesRequestDto{
			PaginationRequestDto: http.PaginationRequestDto{
				Page:  1,
				Limit: 10,
			},
			AttributeID: secondAttributeID,
		})
		s.Require().NoError(err)
		s.Require().NotNil(result)
		s.Equal(1, result.Meta.TotalItems)
		s.Len(result.Data, 1)
		s.Equal(firstValueID, result.Data[0].ID)
	})

	nonExistentID := uuid.MustParse("00000000-0000-0000-0000-000000000001")

	s.Run("Get non-existent attribute fails", func() {
		_, err := s.app.Get(ctx, http.GetAttributeRequestDto{
			AttributeID: nonExistentID,
		})
		s.Require().Error(err)
	})

	s.Run("Update non-existent attribute fails", func() {
		_, err := s.app.Update(ctx, http.UpdateAttributeRequestDto{
			AttributeID: nonExistentID,
			Data: http.UpdateAttributeData{
				Name: "New Name",
			},
		})
		s.Require().Error(err)
	})

	s.Run("Delete non-existent attribute fails", func() {
		err := s.app.Delete(ctx, http.DeleteAttributeRequestDto{
			AttributeID: nonExistentID,
		})
		s.Require().Error(err)
	})

	s.Run("Create value on non-existent attribute fails", func() {
		_, err := s.app.CreateValue(ctx, http.CreateAttributeValueRequestDto{
			AttributeID: nonExistentID,
			Data: http.CreateAttributeValueData{
				Value: "Test",
			},
		})
		s.Require().Error(err)
	})

	s.Run("Update non-existent attribute value fails", func() {
		attr, err := s.app.Create(ctx, http.CreateAttributeRequestDto{
			Data: http.CreateAttributeData{
				Code: "test-attr",
				Name: "Test Attribute",
			},
		})
		s.Require().NoError(err)

		_, err = s.app.UpdateValue(ctx, http.UpdateAttributeValueRequestDto{
			AttributeID:      attr.ID,
			AttributeValueID: nonExistentID,
			Data: http.UpdateAttributeValueData{
				Value: "New Value",
			},
		})
		s.Require().Error(err)
	})
}
