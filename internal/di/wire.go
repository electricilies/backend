//go:build wireinject
// +build wireinject

package di

import (
	"context"

	"backend/config"
	"backend/internal/application"
	"backend/internal/client"
	"backend/internal/delivery/http"
	"backend/internal/domain"
	"backend/internal/infrastructure/repository"
	"backend/internal/service"
	"backend/pkg/logger"

	"github.com/google/wire"
)

var ConfigSet = wire.NewSet(
	config.NewServer,
	logger.NewConfig,
)

var LoggerSet = wire.NewSet(
	logger.New,
)

var DbSet = wire.NewSet(
	client.NewDBConnection,
	client.NewDBQueries,
	client.NewDBTransactor,
)

var EngineSet = wire.NewSet(
	client.NewGin,
)

var ServiceSet = wire.NewSet(
	service.ProvideAttribute,
	wire.Bind(
		new(domain.AttributeService),
		new(*service.Attribute),
	),
	service.ProvideCart,
	wire.Bind(
		new(domain.CartService),
		new(*service.Cart),
	),

	service.ProvideCategory,
	wire.Bind(
		new(domain.CategoryService),
		new(*service.Category),
	),
	service.ProvideOrder,
	wire.Bind(
		new(domain.OrderService),
		new(*service.Order),
	),
	service.ProvideProduct,
	wire.Bind(
		new(domain.ProductService),
		new(*service.Product),
	),
	service.ProvideReview,
	wire.Bind(
		new(domain.ReviewService),
		new(*service.Review),
	),
)

var MiddlewareSet = wire.NewSet(
	http.ProvideAuthMiddleware,
	wire.Bind(
		new(http.AuthMiddleware),
		new(*http.GinAuthMiddleware),
	),
	http.ProvideLoggingMiddleware,
	wire.Bind(
		new(http.LoggingMiddleware),
		new(*http.LoggingMiddlewareImpl),
	),
	http.ProvideMetricMiddleware,
	wire.Bind(
		new(http.MetricMiddleware),
		new(*http.MetricMiddlewareImpl),
	),
	http.ProvideRoleMiddleware,
	wire.Bind(
		new(http.RoleMiddleware),
		new(*http.RoleMiddlewareImpl),
	),
)

var HandlerSet = wire.NewSet(
	http.ProvideAttributeHandler,
	wire.Bind(
		new(http.AttributeHandler),
		new(*http.AttributeHandlerImpl),
	),
	http.ProvideAuthHandler,
	wire.Bind(
		new(http.AuthHandler),
		new(*http.AuthHandlerImpl),
	),
	http.ProvideCategoryHandler,
	wire.Bind(
		new(http.CategoryHandler),
		new(*http.CategoryHandlerImpl),
	),
	http.ProvideHealthHandler,
	wire.Bind(
		new(http.HealthHandler),
		new(*http.HealthHandlerImpl),
	),
	http.ProvideOrderHandler,
	wire.Bind(
		new(http.OrderHandler),
		new(*http.OrderHandlerImpl),
	),
	http.ProvideProductHandler,
	wire.Bind(
		new(http.ProductHandler),
		new(*http.ProductHandlerImpl),
	),
	http.ProvideReviewHandler,
	wire.Bind(
		new(http.ReviewHandler),
		new(*http.ReviewHandlerImpl),
	),
	http.ProvideCartHandler,
	wire.Bind(
		new(http.CartHandler),
		new(*http.CartHandlerImpl),
	),
)

var ApplicationSet = wire.NewSet(
	application.ProvideAttribute,
	wire.Bind(
		new(application.Attribute),
		new(*application.AttributeImpl),
	),
	application.ProvideCart,
	wire.Bind(
		new(application.Cart),
		new(*application.CartImpl),
	),
	application.ProvideCategory,
	wire.Bind(
		new(application.Category),
		new(*application.CategoryImpl),
	),
	application.ProvideOrder,
	wire.Bind(
		new(application.Order),
		new(*application.OrderImpl),
	),
	application.ProvideProduct,
	wire.Bind(
		new(application.Product),
		new(*application.ProductImpl),
	),
	application.ProvideReview,
	wire.Bind(
		new(application.Review),
		new(*application.ReviewImpl),
	),
)

var RepositorySet = wire.NewSet(
	repository.ProvidePostgresAttribute,
	wire.Bind(
		new(domain.AttributeRepository),
		new(*repository.PostgresAttribute),
	),
	repository.ProvidePostgresCart,
	wire.Bind(
		new(domain.CartRepository),
		new(*repository.PostgresCart),
	),
	repository.ProvidePostgresCategory,
	wire.Bind(
		new(domain.CategoryRepository),
		new(*repository.PostgresCategory),
	),
	repository.ProvidePostgresOrder,
	wire.Bind(
		new(domain.OrderRepository),
		new(*repository.PostgresOrder),
	),
	repository.ProvidePostgresProduct,
	wire.Bind(
		new(domain.ProductRepository),
		new(*repository.PostgresProduct),
	),
	repository.ProvidePostgresReview,
	wire.Bind(
		new(domain.ReviewRepository),
		new(*repository.PostgresReview),
	),
)

var RouterSet = wire.NewSet(
	http.ProvideRouter,
	wire.Bind(
		new(http.Router),
		new(*http.GinRouter),
	),
)

var ClientSet = wire.NewSet(
	client.NewRedis,
	client.NewS3,
	client.NewKeycloak,
	client.NewS3Presign,
	client.NewValidate,
)

func InitializeServer(ctx context.Context) *http.Server {
	wire.Build(
		ApplicationSet,
		ClientSet,
		ConfigSet,
		DbSet,
		EngineSet,
		HandlerSet,
		LoggerSet,
		MiddlewareSet,
		RouterSet,
		ServiceSet,
		RepositorySet,
		http.NewServer,
	)
	return nil
}
