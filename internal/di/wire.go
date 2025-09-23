//go:build wireinject
// +build wireinject

package di

import (
	app "backend/internal/application"
	userService "backend/internal/domain/user"
	userRepo "backend/internal/infrastructure/user"
	handler "backend/internal/interface/api/handler"
	"backend/internal/interface/api/router"
	"backend/internal/server"

	"backend/internal/di/db"
	"backend/internal/di/ginengine"

	"github.com/google/wire"
)

var DBSet = wire.NewSet(
	db.NewDB,
)

var EngineSet = wire.NewSet(
	ginengine.NewEngine,
)

var RepositorySet = wire.NewSet(
	userRepo.NewRepository,
)

var ServiceSet = wire.NewSet(
	userService.NewService,
)

var AppSet = wire.NewSet(
	app.NewUser,
)

var HandlerSet = wire.NewSet(
	handler.NewUserHandler,
	handler.NewHealthCheck,
)

var RouterSet = wire.NewSet(
	router.NewRouter,
)

func InitializeServer() *server.Server {
	wire.Build(DBSet, EngineSet, RepositorySet, ServiceSet, AppSet, HandlerSet, RouterSet, server.NewServer)
	return nil
}
