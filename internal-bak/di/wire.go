//go:build wireinject
// +build wireinject

package di

import (
	"context"

	"backend/config"
	app "backend/internal/application"
	"backend/internal/client"

	"backend/internal/helper"
	domainattribute "backend/internal/domain/attribute"
	domaincart "backend/internal/domain/cart"
	domaincategory "backend/internal/domain/category"
	domainproduct "backend/internal/domain/product"
	domainreview "backend/internal/domain/review"
	domainuser "backend/internal/domain/user"
	domainpayment "backend/internal/domain/payment"
	infrasattribute "backend/internal/infrastructure/attribute"
	infrascart "backend/internal/infrastructure/cart"
	infrascategory "backend/internal/infrastructure/category"
	infrasproduct "backend/internal/infrastructure/product"
	infrasreview "backend/internal/infrastructure/review"
	infrasuser "backend/internal/infrastructure/user"
	infraspayment "backend/internal/infrastructure/payment"

	handler "backend/internal/interface/api/handler"
	middleware "backend/internal/interface/api/middleware"
	"backend/internal/interface/api/router"
	"backend/internal/server"
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

var RepositorySet = wire.NewSet(
	infrasuser.ProvideRepository,
	wire.Bind(
		new(domainuser.Repository),
		new(*infrasuser.RepositoryImpl),
	),
	infrascategory.ProvideRepository,
	wire.Bind(
		new(domaincategory.Repository),
		new(*infrascategory.RepositoryImpl),
	),
	infrasproduct.ProvideRepository,
	wire.Bind(
		new(domainproduct.Repository),
		new(*infrasproduct.RepositoryImpl),
	),
	infrasreview.ProvideRepository,
	wire.Bind(
		new(domainreview.Repository),
		new(*infrasreview.RepositoryImpl),
	),
	infrascart.ProvideRepository,
	wire.Bind(
		new(domaincart.Repository),
		new(*infrascart.RepositoryImpl),
	),
	infrasattribute.ProvideRepository,
	wire.Bind(
		new(domainattribute.Repository),
		new(*infrasattribute.RepositoryImpl),
	),
	infraspayment.ProvideRepository,
	wire.Bind(
		new(domainpayment.Repository),
		new(*infraspayment.RepositoryImpl),
	),
)

var ServiceSet = wire.NewSet(
	domainuser.ProvideService,
	wire.Bind(
		new(domainuser.Service),
		new(*domainuser.ServiceImpl),
	),
)

var AppSet = wire.NewSet(
	app.ProvideAttribute,
	wire.Bind(
		new(app.Attribute),
		new(*app.AttributeImpl),
	),
	app.ProvideCategory,
	wire.Bind(
		new(app.Category),
		new(*app.CategoryImpl),
		),
	app.ProvideProduct,
	wire.Bind(
		new(app.Product),
		new(*app.ProductImpl),
	),

	app.ProvideUser,
	wire.Bind(
		new(app.User),
		new(*app.UserImpl),
	),
	app.ProvideReview,
	wire.Bind(
		new(app.Review),
		new(*app.ReviewImpl),
	),
	app.ProvideCart,
	wire.Bind(
		new(app.Cart),
		new(*app.CartImpl),
	),
)

var MiddlewareSet = wire.NewSet(
	middleware.ProvideAuth,
	wire.Bind(
		new(middleware.Auth),
		new(*middleware.AuthImpl),
	),
	middleware.ProvideLogging,
	wire.Bind(
		new(middleware.Logging),
		new(*middleware.LoggingImpl),
		),
	middleware.ProvideMetric,
	wire.Bind(
		new(middleware.Metric),
		new(*middleware.MetricImpl),
	),
	middleware.ProvideRole,
	wire.Bind(
		new(middleware.Role),
		new(*middleware.RoleImpl),
	),
)

var HandlerSet = wire.NewSet(
	handler.ProvideAttribute,
	wire.Bind(
		new(handler.Attribute),
		new(*handler.AttributeImpl),
	),
	handler.ProvideAuth,
	wire.Bind(
		new(handler.Auth),
		new(*handler.AuthImpl),
	),
	handler.ProvideCategory,
	wire.Bind(
		new(handler.Category),
		new(*handler.CategoryImpl),
	),
	handler.ProvideHealthCheck,
	wire.Bind(
		new(handler.HealthCheck),
		new(*handler.HealthCheckImpl),
	),
	handler.ProvideOrder,
	wire.Bind(
		new(handler.Order),
		new(*handler.OrderImpl),
	),

	handler.ProvidePayment,
	wire.Bind(
		new(handler.Payment),
		new(*handler.PaymentImpl),
	),
	handler.ProvideProduct,
	wire.Bind(
		new(handler.Product),
		new(*handler.ProductImpl),
	),
	handler.ProvideUser,
	wire.Bind(
		new(handler.User),
		new(*handler.UserImpl),
	),
	handler.ProvideReview,
	wire.Bind(
		new(handler.Review),
		new(*handler.ReviewImpl),
	),
	handler.ProvideCart,
	wire.Bind(
		new(handler.Cart),
		new(*handler.CartImpl),
	),
	handler.ProvideReturnRequest,
	wire.Bind(
		new(handler.ReturnRequest),
		new(*handler.ReturnRequestImpl),
	),
	handler.ProvideRefund,
	wire.Bind(
		new(handler.Refund),
		new(*handler.RefundImpl),
	),
)
var RouterSet = wire.NewSet(
	router.Provide,
	wire.Bind(
		new(router.Router),
		new(*router.RouterImpl),
	),
)

var ClientSet = wire.NewSet(
	client.NewRedis,
	client.NewS3,
	client.NewKeycloak,
	client.NewS3Presign,
)

func InitializeServer(ctx context.Context) *server.Server {
	wire.Build(
		AppSet,
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
		helper.NewTokenManager,
		server.New,
	)
	return nil
}
