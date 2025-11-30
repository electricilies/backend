// vimm: tabstop=4:
//go:build integration

package application_test

import (
	"context"
	"strings"

	"backend/config"
	"backend/internal/application"
	"backend/internal/client"
	"backend/internal/infrastructure/cacheredis"
	"backend/internal/infrastructure/repositorypostgres"
	"backend/internal/service"
	"backend/test/integration/component"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/suite"
)

type ProductTestSuite struct {
	suite.Suite
	containers *component.Containers
	app        *application.Attribute
}

func (s *ProductTestSuite) newContainersConfig() *component.ContainersConfig {
	containersConfig := component.NewContainersConfig()
	containersConfig.DB.Enabled = true
	containersConfig.Redis.Enabled = true
	return containersConfig
}

func (s *ProductTestSuite) newConfig(
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

func (s *ProductTestSuite) SetupSuite() {
	ctx := s.T().Context()
	containersConfig := s.newContainersConfig()
	config := s.newConfig(ctx)

	var err error
	s.containers, err = component.NewContainers(ctx, containersConfig)
	s.Require().NoError(err, "failed to start containers")

	validate := validator.New(
		validator.WithRequiredStructEnabled(),
	)

	conn := client.NewDBConnection(ctx, config)
	queries := client.NewDBQueries(conn)
	attributeRepo := repositorypostgres.ProvideAttribute(queries, conn)

	attributeService := service.ProvideAttribute(validate)

	redisClient := client.NewRedis(ctx, config)
	attributeCache := cacheredis.ProvideAttribute(redisClient)
	s.app = application.ProvideAttribute(attributeRepo, attributeService, attributeCache)
}
