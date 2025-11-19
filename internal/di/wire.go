//go:build wireinject
// +build wireinject

package di

import (
	"context"

	"backend/config"
	"backend/internal/client"
	"backend/internal/delivery/http"
	"backend/internal/domain"
	"backend/internal/serviceimpl"
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
	serviceimpl.ProvideAttribute,
	wire.Bind(
		new(domain.AttributeService),
		new(*serviceimpl.Attribute),
	),
	serviceimpl.ProvideCategory,
	wire.Bind(
		new(domain.CategoryService),
		new(*serviceimpl.Category),
	),
	serviceimpl.ProvideProduct,
	wire.Bind(
		new(domain.ProductService),
		new(*serviceimpl.Product),
	),
	serviceimpl.ProvideReview,
	wire.Bind(
		new(domain.ReviewService),
		new(*serviceimpl.Review),
	),
	serviceimpl.ProvideCart,
	wire.Bind(
		new(domain.CartService),
		new(*serviceimpl.Cart),
	),
	serviceimpl.ProvidePayment,
	wire.Bind(
		new(domain.PaymentService),
		new(*serviceimpl.Payment),
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
		new(*http.GinAttributeHandler),
	),
	http.ProvideAuthHandler,
	wire.Bind(
		new(http.AuthHandler),
		new(*http.GinAuthHandler),
	),
	http.ProvideCategoryHandler,
	wire.Bind(
		new(http.CategoryHandler),
		new(*http.GinCategoryHandler),
	),
	http.ProvideHealthHandler,
	wire.Bind(
		new(http.HealthHandler),
		new(*http.GinHealthHandler),
	),
	http.ProvideOrderHandler,
	wire.Bind(
		new(http.OrderHandler),
		new(*http.GinOrderHandler),
	),

	http.ProvidePaymentHandler,
	wire.Bind(
		new(http.PaymentHandler),
		new(*http.GinPaymentHandler),
	),
	http.ProvideProductHandler,
	wire.Bind(
		new(http.ProductHandler),
		new(*http.GinProductHandler),
	),
	http.ProvideReviewHandler,
	wire.Bind(
		new(http.ReviewHandler),
		new(*http.GinReviewHandler),
	),
	http.ProvideCartHandler,
	wire.Bind(
		new(http.CartHandler),
		new(*http.GinCartHandler),
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
)

func InitializeServer(ctx context.Context) *http.Server {
	wire.Build(
		ClientSet,
		ConfigSet,
		DbSet,
		EngineSet,
		HandlerSet,
		LoggerSet,
		MiddlewareSet,
		RouterSet,
		http.NewServer,
	)
	return nil
}
