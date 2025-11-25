package http

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Router interface {
	RegisterRoutes(*gin.Engine)
}

type GinRouter struct {
	categoryHandler  CategoryHandler
	productHandler   ProductHandler
	attributeHandler AttributeHandler
	orderHandler     OrderHandler
	reviewHandler    ReviewHandler
	cartHandler      CartHandler

	healthHandler     HealthHandler
	metricMiddleware  MetricMiddleware
	loggingMiddleware LoggingMiddleware
	authMiddleware    AuthMiddleware
}

func ProvideRouter(
	healthCheckHandler HealthHandler,
	metricMiddleware MetricMiddleware,
	loggingMiddleware LoggingMiddleware,
	authMiddleware AuthMiddleware,
	categoryHandler CategoryHandler,
	productHandler ProductHandler,
	attributeHandler AttributeHandler,
	orderHandler OrderHandler,
	reviewHandler ReviewHandler,
	cartHandler CartHandler,
) *GinRouter {
	return &GinRouter{
		healthHandler:     healthCheckHandler,
		metricMiddleware:  metricMiddleware,
		loggingMiddleware: loggingMiddleware,
		authMiddleware:    authMiddleware,
		categoryHandler:   categoryHandler,
		productHandler:    productHandler,
		attributeHandler:  attributeHandler,
		orderHandler:      orderHandler,
		reviewHandler:     reviewHandler,
		cartHandler:       cartHandler,
	}
}

func (r *GinRouter) RegisterRoutes(e *gin.Engine) {
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
		cart := api.Group("/carts")
		{
			cart.Use(r.authMiddleware.Handler())
			cart.GET("/:cart_id", r.cartHandler.Get)
			cart.POST("/:cart_id/item", r.cartHandler.CreateItem)
			cart.PATCH("/:cart_id/item", r.cartHandler.UpdateItem)
			cart.DELETE("/:cart_id/item", r.cartHandler.RemoveItem)
		}
		categories := api.Group("/categories")
		{
			categories.GET("", r.categoryHandler.List)
			categories.GET("/:category_id", r.categoryHandler.Get)
			categories.POST("", r.categoryHandler.Create)
			categories.PATCH("/:category_id", r.categoryHandler.Update)
		}

		products := api.Group("/products")
		{
			products.POST("", r.productHandler.Create)
			products.GET("", r.productHandler.List)
			products.GET("/:product_id", r.productHandler.Get)
			products.DELETE("/:product_id", r.productHandler.Delete)
			products.POST("/:product_id/images", r.productHandler.AddImages)
			products.DELETE("/:product_id/images", r.productHandler.DeleteImages)
			products.PATCH("/:product_id", r.productHandler.Update)
			products.GET("/images/upload-url", r.productHandler.GetUploadImageURL)
			products.GET("/images/delete-url/:image_id", r.productHandler.GetDeleteImageURL)
			products.POST("/:product_id/variants", r.productHandler.AddVariants)
			products.PATCH("/:product_id/variants/:variant_id", r.productHandler.UpdateVariant)
			products.PATCH("/:product_id/options", r.productHandler.UpdateOptions)
		}

		attributes := api.Group("/attributes")
		{
			attributes.GET("", r.attributeHandler.List)
			attributes.GET("/:attribute_id/values", r.attributeHandler.ListValues)
			attributes.POST("", r.attributeHandler.Create)
			attributes.POST("/:attribute_id/values", r.attributeHandler.CreateValue)
			attributes.GET("/:attribute_id", r.attributeHandler.Get)
			attributes.PATCH("/:attribute_id", r.attributeHandler.Update)
			attributes.DELETE("/:attribute_id", r.attributeHandler.Delete)
			attributes.DELETE("/:attribute_id/values/:value_id", r.attributeHandler.DeleteValue)
			attributes.PATCH("/:attribute_id/values/:value_id", r.attributeHandler.UpdateValue)
		}

		orders := api.Group("/orders")
		{
			orders.Use(r.authMiddleware.Handler())
			orders.GET("", r.orderHandler.List)
			orders.POST("", r.orderHandler.Create)
			orders.GET("/:order_id", r.orderHandler.Get)
			orders.PATCH("/:order_id", r.orderHandler.Update)
			orders.DELETE("/:order_id", r.orderHandler.Delete)
		}

		// returnRequests := api.Group("/return-requests")
		// {
		// 	returnRequests.GET("", r.returnHandler.List)
		// 	returnRequests.POST("", r.returnHandler.Create)
		// 	returnRequests.GET("/:return_request_id", r.returnHandler.Get)
		// 	returnRequests.PATCH("/:return_request_id", r.returnHandler.Update)
		// }
		//
		// refunds := api.Group("/refunds")
		// {
		// 	refunds.GET("", r.refundHandler.List)
		// 	refunds.GET("/:refund_id", r.refundHandler.Get)
		// }

		reviews := api.Group("/reviews")
		{
			reviews.GET("", r.reviewHandler.List)
			reviews.POST("", r.reviewHandler.Create)
			reviews.GET("/:review_id", r.reviewHandler.Get)
			reviews.PATCH("/:review_id", r.reviewHandler.Update)
			reviews.DELETE("/:review_id", r.reviewHandler.Delete)
		}
	}
}
