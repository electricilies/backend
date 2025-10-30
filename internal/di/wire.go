//go:build wireinject
// +build wireinject

package di

import (
	app "backend/internal/application"
	"backend/internal/di/client"
	"backend/internal/di/db"
	"backend/internal/di/ginengine"
	userservice "backend/internal/domain/user"
	userrepo "backend/internal/infrastructure/user"
	handler "backend/internal/interface/api/handler"
	middleware "backend/internal/interface/api/middleware"
	"backend/internal/interface/api/router"
	"backend/internal/server"

	"github.com/google/wire"
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
)

var ServiceSet = wire.NewSet(
	userservice.NewService,
)

var AppSet = wire.NewSet(
	app.NewUser,
)

var MiddlewareSet = wire.NewSet(
	middleware.NewMetric,
	middleware.NewLogging,
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
)

func InitializeServer() *server.Server {
	wire.Build(DbSet, EngineSet, RepositorySet, ServiceSet, AppSet, MiddlewareSet, HandlerSet, RouterSet, ClientSet, server.New)
	return nil
}
