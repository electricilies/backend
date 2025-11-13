package router

import (
	"backend/internal/interface/api/handler"
	"backend/internal/interface/api/middleware"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Router interface {
	RegisterRoutes(e *gin.Engine)
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
	authMiddleware    middleware.Auth
}

func New(
	userHandler handler.User,
	healthCheckHandler handler.HealthCheck,
	metricMiddleware middleware.Metric,
	loggingMiddleware middleware.Logging,
	authMiddleware middleware.Auth,
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
		authMiddleware:    authMiddleware,
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

func (r *router) RegisterRoutes(e *gin.Engine) {
	e.Use(r.metricMiddleware.Handler())
	e.GET("/metrics", gin.WrapH(promhttp.Handler()))
	health := e.Group("/health")
	{
		health.GET("/live", r.healthHandler.Liveness)
		health.GET("/ready", r.healthHandler.Readiness)
	}
	api := e.Group("/api")
	{
		api.Use(r.loggingMiddleware.Handler())
		api.Use(r.authMiddleware.Handler())
		users := api.Group("/users")
		{
			users.GET("", r.userHandler.List)
			users.POST("", r.userHandler.Create)
			users.GET("/:user_id", r.userHandler.Get)
			users.PUT("/:user_id", r.userHandler.Update)
			users.DELETE("/:user_id", r.userHandler.Delete)
			users.GET("/:user_id/return-requests", r.userHandler.GetReturnRequests)
		}
		cart := api.Group("/carts")
		{
			cart.GET("", r.cartHandler.GetCartByUser)
			cart.POST("/item", r.cartHandler.AddItem)
			cart.PUT("/item", r.cartHandler.UpdateItem)
			cart.DELETE("/item", r.cartHandler.RemoveItem)
		}
		categories := api.Group("/categories")
		{
			categories.GET("", r.categoryHandler.List)
			categories.POST("", r.categoryHandler.Create)
			categories.PUT("/:category_id", r.categoryHandler.Update)
		}

		products := api.Group("/products")
		{
			products.GET("", r.productHandler.List)
			products.POST("", r.productHandler.Create)
			products.GET("/:product_id", r.productHandler.Get)
			products.PUT("/:product_id", r.productHandler.Update)
			products.DELETE("/:product_id", r.productHandler.Delete)
			products.POST("/options", r.productHandler.CreateProductOption)
			products.GET("/images/upload-url", r.productHandler.GetUploadImageURL)
			products.GET("/images/delete-url/:image_id", r.productHandler.GetDeleteImageURL)
			products.POST("/variants", r.productHandler.CreateProductVariant)
			products.PUT("/variants/:variant_id", r.productHandler.UpdateProductVariant)
			products.PUT("/options/:option_id", r.productHandler.UpdateProductOption)
			products.PUT("/:product_id/images/bulk", r.productHandler.CreateProductImages)
		}

		attributes := api.Group("/attributes")
		{
			attributes.GET("", r.attributeHandler.List)
			attributes.POST("", r.attributeHandler.Create)
			attributes.GET("/:attribute_id", r.attributeHandler.Get)
			attributes.PUT("/:attribute_id", r.attributeHandler.Update)
			attributes.DELETE("/:id", r.attributeHandler.Delete)
			attributes.PUT("/:attribute_id/values/bulk", r.attributeHandler.UpdateAttributeValues)
		}

		// payments := api.Group("/payment")
		// {
		// }

		orders := api.Group("/orders")
		{
			orders.GET("", r.orderHandler.List)
			orders.POST("", r.orderHandler.Create)
			orders.GET("/:order_id", r.orderHandler.Get)
			orders.PUT("/:order_id", r.orderHandler.Update)
			orders.DELETE("/:order_id", r.orderHandler.Delete)
		}

		returnRequests := api.Group("/return-requests")
		{
			returnRequests.GET("", r.returnHandler.List)
			returnRequests.POST("", r.returnHandler.Create)
			returnRequests.GET("/:return_request_id", r.returnHandler.Get)
			returnRequests.PUT("/:return_request_id", r.returnHandler.Update)
		}

		refunds := api.Group("/refunds")
		{
			refunds.GET("", r.refundHandler.List)
			refunds.GET("/:refund_id", r.refundHandler.Get)
		}

		reviews := api.Group("/reviews")
		{
			reviews.GET("", r.reviewHandler.ListReviewsByProducts)
			reviews.POST("", r.reviewHandler.Create)
			reviews.GET("/:review_id", r.reviewHandler.Get)
			reviews.PUT("/:review_id", r.reviewHandler.Update)
			reviews.DELETE("/:review_id", r.reviewHandler.Delete)
		}
	}
}
