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
	"backend/internal/infrastructure/cacheredis"
	"backend/internal/infrastructure/objectstorages3"
	"backend/internal/infrastructure/repositorypostgres"
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
	// service.ProvideReview,
	// wire.Bind(
	// 	new(domain.ReviewService),
	// 	new(*service.Review),
	// ),
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
	http.ProvideAuthHandler,
	wire.Bind(
		new(http.AuthHandler),
		new(*http.AuthHandlerImpl),
	),
	http.ProvideFlushCacheRedisHandler,
	wire.Bind(
		new(http.FlushCacheHandler),
		new(*http.FlushCacheRedisHandler),
	),
	http.ProvideHealthHandler,
	wire.Bind(
		new(http.HealthHandler),
		new(*http.HealthHandlerImpl),
	),
	http.ProvideAttributeHandler,
	wire.Bind(
		new(http.AttributeHandler),
		new(*http.AttributeHandlerImpl),
	),
	http.ProvideCategoryHandler,
	wire.Bind(
		new(http.CategoryHandler),
		new(*http.CategoryHandlerImpl),
	),
	http.ProvideProductHandler,
	wire.Bind(
		new(http.ProductHandler),
		new(*http.ProductHandlerImpl),
	),
	http.ProvideCartHandler,
	wire.Bind(
		new(http.CartHandler),
		new(*http.CartHandlerImpl),
	),
	http.ProvideOrderHandler,
	wire.Bind(
		new(http.OrderHandler),
		new(*http.OrderHandlerImpl),
	),
	// http.ProvideReviewHandler,
	// wire.Bind(
	// 	new(http.ReviewHandler),
	// 	new(*http.ReviewHandlerImpl),
	// ),
)

var ApplicationSet = wire.NewSet(
	application.ProvideAttribute,
	wire.Bind(
		new(http.AttributeApplication),
		new(*application.Attribute),
	),
	application.ProvideCart,
	wire.Bind(
		new(http.CartApplication),
		new(*application.Cart),
	),
	application.ProvideCategory,
	wire.Bind(
		new(http.CategoryApplication),
		new(*application.Category),
	),
	application.ProvideOrder,
	wire.Bind(
		new(http.OrderApplication),
		new(*application.Order),
	),
	application.ProvideProduct,
	wire.Bind(
		new(http.ProductApplication),
		new(*application.Product),
	),
	// application.ProvideReview,
	// wire.Bind(
	// 	new(http.ReviewApplication),
	// 	new(*application.Review),
	// ),
)

var RepositorySet = wire.NewSet(
	repositorypostgres.ProvideAttribute,
	wire.Bind(
		new(domain.AttributeRepository),
		new(*repositorypostgres.Attribute),
	),
	repositorypostgres.ProvideCart,
	wire.Bind(
		new(domain.CartRepository),
		new(*repositorypostgres.Cart),
	),
	repositorypostgres.ProvideCategory,
	wire.Bind(
		new(domain.CategoryRepository),
		new(*repositorypostgres.Category),
	),
	repositorypostgres.ProvideOrder,
	wire.Bind(
		new(domain.OrderRepository),
		new(*repositorypostgres.Order),
	),
	repositorypostgres.ProvideProduct,
	wire.Bind(
		new(domain.ProductRepository),
		new(*repositorypostgres.Product),
	),
	// repositorypostgres.ProvideReview,
	// wire.Bind(
	// 	new(domain.ReviewRepository),
	// 	new(*repositorypostgres.Review),
	// ),
)

var RouterSet = wire.NewSet(
	http.ProvideRouter,
	wire.Bind(
		new(http.Router),
		new(*http.GinRouter),
	),
)

var ClientSet = wire.NewSet(
	client.NewKeycloak,
	client.NewRedis,
	client.NewS3,
	client.NewS3Presign,
	client.NewValidate,
)

var CacheSet = wire.NewSet(
	cacheredis.ProvideProduct,
	wire.Bind(
		new(application.ProductCache),
		new(*cacheredis.Product),
	),
	// cacheredis.ProvideReview,
	// wire.Bind(
	// 	new(application.ReviewCache),
	// 	new(*cacheredis.Review),
	// ),
	cacheredis.ProvideCategory,
	wire.Bind(
		new(application.CategoryCache),
		new(*cacheredis.Category),
	),
	cacheredis.ProvideAttribute,
	wire.Bind(
		new(application.AttributeCache),
		new(*cacheredis.Attribute),
	),
	cacheredis.ProvideCart,
	wire.Bind(
		new(application.CartCache),
		new(*cacheredis.Cart),
	),
)

var ObjectStorageSet = wire.NewSet(
	objectstorages3.ProvideProduct,
	wire.Bind(
		new(application.ProductObjectStorage),
		new(*objectstorages3.Product),
	),
)

func InitializeServer(ctx context.Context) *http.Server {
	wire.Build(
		ApplicationSet,
		CacheSet,
		ClientSet,
		ConfigSet,
		DbSet,
		EngineSet,
		HandlerSet,
		LoggerSet,
		MiddlewareSet,
		RepositorySet,
		RouterSet,
		ServiceSet,
		ObjectStorageSet,
		http.NewServer,
	)
	return nil
}
