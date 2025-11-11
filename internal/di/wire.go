//go:build wireinject
// +build wireinject

package di

import (
	"backend/config"
	app "backend/internal/application"
	"backend/internal/di/client"
	"backend/internal/di/db"
	"backend/internal/di/ginengine"
	userservice "backend/internal/domain/user"
	"backend/internal/helper"
	cartrepo "backend/internal/infrastructure/cart"
	categoryrepo "backend/internal/infrastructure/category"
	productrepo "backend/internal/infrastructure/product"
	reviewrepo "backend/internal/infrastructure/review"
	userrepo "backend/internal/infrastructure/user"
	handler "backend/internal/interface/api/handler"
	middleware "backend/internal/interface/api/middleware"
	"backend/internal/interface/api/router"
	"backend/internal/server"
	"backend/pkg/logger"

	"github.com/google/wire"
)

var ConfigSet = wire.NewSet(
	config.New,
	logger.NewConfig,
)

var LoggerSet = wire.NewSet(
	logger.New,
)

var DbSet = wire.NewSet(
	db.NewConnection,
	db.New,
	db.NewTransactor,
)

var EngineSet = wire.NewSet(
	ginengine.New,
)

var RepositorySet = wire.NewSet(
	userrepo.NewRepository,
	productrepo.NewRepository,
	reviewrepo.NewRepository,
	cartrepo.NewRepository,
	categoryrepo.NewRepository,
)

var ServiceSet = wire.NewSet(
	userservice.NewService,
)

var AppSet = wire.NewSet(
	app.NewUser,
	app.NewProduct,
	app.NewCart,
	app.NewReview,
	app.NewCategory,
)

var MiddlewareSet = wire.NewSet(
	middleware.NewMetric,
	middleware.NewLogging,
	middleware.NewJWTVerify,
)

var HandlerSet = wire.NewSet(
	handler.NewUser,
	handler.NewHealthCheck,
	handler.NewCategory,
	handler.NewProduct,
	handler.NewAttribute,
	handler.NewPayment,
	handler.NewOrder,
	handler.NewReturn,
	handler.NewRefund,
	handler.NewReview,
	handler.NewCart,
)

var RouterSet = wire.NewSet(
	router.New,
)

var ClientSet = wire.NewSet(
	client.NewRedis,
	client.NewS3,
	client.NewKeycloak,
	client.NewS3Presign,
)

func InitializeServer() *server.Server {
	wire.Build(ConfigSet, LoggerSet, DbSet, EngineSet, RepositorySet, ServiceSet, AppSet, MiddlewareSet, HandlerSet, RouterSet, ClientSet, helper.NewTokenManager, server.New)
	return nil
}
