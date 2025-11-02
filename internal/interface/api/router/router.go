package router

import (
	"backend/internal/interface/api/handler"
	"backend/internal/interface/api/middleware"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Router interface {
	RegisterRoutes(engine *gin.Engine)
}

type router struct {
	userHandler      handler.User
	categoryHandler  handler.Category
	productHandler   handler.Product
	attributeHandler handler.Attribute
	paymentHandler   handler.Payment
	orderHandler     handler.Order
	returnHandler    handler.ReturnRequest
	refundHandler    handler.Refund
	reviewHandler    handler.Review
	cartHandler      handler.Cart

	healthHandler     handler.HealthCheck
	metricMiddleware  middleware.Metric
	loggingMiddleware middleware.Logging
}

func New(
	userHandler handler.User,
	healthCheckHandler handler.HealthCheck,
	metricMiddleware middleware.Metric,
	loggingMiddleware middleware.Logging,
	categoryHandler handler.Category,
	productHandler handler.Product,
	attributeHandler handler.Attribute,
	paymentHandler handler.Payment,
	orderHandler handler.Order,
	returnHandler handler.ReturnRequest,
	refundHandler handler.Refund,
	reviewHandler handler.Review,
	cartHandler handler.Cart,
) Router {
	return &router{
		userHandler:       userHandler,
		healthHandler:     healthCheckHandler,
		metricMiddleware:  metricMiddleware,
		loggingMiddleware: loggingMiddleware,
		categoryHandler:   categoryHandler,
		productHandler:    productHandler,
		attributeHandler:  attributeHandler,
		paymentHandler:    paymentHandler,
		orderHandler:      orderHandler,
		returnHandler:     returnHandler,
		refundHandler:     refundHandler,
		reviewHandler:     reviewHandler,
		cartHandler:       cartHandler,
	}
}

func (r *router) RegisterRoutes(engine *gin.Engine) {
	engine.Use(r.metricMiddleware.Handler())
	engine.Use(r.loggingMiddleware.Handler())
	engine.GET("/metrics", gin.WrapH(promhttp.Handler()))
	engine.GET("/health", r.healthHandler.Liveness)
	engine.GET("/ready", r.healthHandler.Readiness)
	api := engine.Group("/api")
	{
		users := api.Group("/users")
		{
			users.GET("", r.userHandler.List)
			users.POST("", r.userHandler.Create)
			users.GET("/:user_id", r.userHandler.Get)
			users.PUT("/:user_id", r.userHandler.Update)
			users.DELETE("/:user_id", r.userHandler.Delete)
			users.GET("/:user_id/return_requests", r.userHandler.GetReturnRequests)
		}
		cart := api.Group("/carts")
		{
			cart.GET("", r.cartHandler.GetCartByUser)
			cart.POST("", r.cartHandler.AddItem)
			cart.PUT("", r.cartHandler.UpdateItem)
			cart.DELETE("", r.cartHandler.RemoveItem)
		}
		categories := api.Group("/categories")
		{
			categories.GET("", r.categoryHandler.List)
			categories.POST("", r.categoryHandler.Create)
			categories.PUT("/:id", r.categoryHandler.Update)
			categories.DELETE("/:id", r.categoryHandler.Delete)
		}

		products := api.Group("/products")
		{
			products.GET("", r.productHandler.List)
			products.POST("", r.productHandler.Create)
			products.GET("/:id", r.productHandler.Get)
			products.PUT("/:id", r.productHandler.Update)
			products.DELETE("/:id", r.productHandler.Delete)
		}

		attributes := api.Group("/attributes")
		{
			attributes.GET("", r.attributeHandler.List)
			attributes.POST("", r.attributeHandler.Create)
			attributes.GET("/:id", r.attributeHandler.Get)
			attributes.PUT("/:id", r.attributeHandler.Update)
			attributes.DELETE("/:id", r.attributeHandler.Delete)
		}

		payments := api.Group("/payment")
		{
			payments.GET("/methods", r.paymentHandler.ListMethods)
		}

		orders := api.Group("/orders")
		{
			orders.GET("", r.orderHandler.List)
			orders.POST("", r.orderHandler.Create)
			orders.GET("/:id", r.orderHandler.Get)
			orders.PUT("/:id", r.orderHandler.Update)
			orders.DELETE("/:id", r.orderHandler.Delete)
		}

		returnRequests := api.Group("/return-requests")
		{
			returnRequests.GET("", r.returnHandler.List)
			returnRequests.POST("", r.returnHandler.Create)
			returnRequests.GET("/:id", r.returnHandler.Get)
			returnRequests.PUT("/:id", r.returnHandler.Update)
		}

		refunds := api.Group("/refunds")
		{
			refunds.GET("", r.refundHandler.List)
			refunds.GET("/:id", r.refundHandler.Get)
		}

		reviews := api.Group("/reviews")
		{
			reviews.GET("", r.reviewHandler.ListReviewsByProduct)
			reviews.POST("", r.reviewHandler.Create)
			reviews.GET("/:id", r.reviewHandler.Get)
			reviews.PUT("/:id", r.reviewHandler.Update)
			reviews.DELETE("/:id", r.reviewHandler.Delete)
		}
	}
}
