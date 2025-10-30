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
	userHandler       handler.User
	healthHandler     handler.HealthCheck
	metricMiddleware  middleware.Metric
	loggingMiddleware middleware.Logging
}

func New(userHandler handler.User, healthCheckHandler handler.HealthCheck, metricMiddleware middleware.Metric, loggingMiddleware middleware.Logging) Router {
	return &router{
		userHandler:       userHandler,
		healthHandler:     healthCheckHandler,
		metricMiddleware:  metricMiddleware,
		loggingMiddleware: loggingMiddleware,
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
			users.GET("/:id", r.userHandler.Get)
			users.PUT("/:id", r.userHandler.Update)
			users.DELETE("/:id", r.userHandler.Delete)
			// cart := users.Group("/:user_id/cart")
			// {
			// 	cart.GET("", r.cartHandler.Get)
			// 	cart.POST("", r.cartHandler.AddItem)
			// 	cart.PUT("/:item_id", r.cartHandler.UpdateItem)
			// 	cart.DELETE("/:item_id", r.cartHandler.RemoveItem)
			// }
		}

		// categories := api.Group("/categories")
		// {
		// 	categories.GET("", r.categoryHandler.List)
		// 	categories.POST("", r.categoryHandler.Create)
		// 	categories.GET("/:id", r.categoryHandler.Get)
		// 	categories.PUT("/:id", r.categoryHandler.Update)
		// 	categories.DELETE("/:id", r.categoryHandler.Delete)
		// }
		//
		// products := api.Group("/products")
		// {
		// 	products.GET("", r.productHandler.List)
		// 	products.POST("", r.productHandler.Create)
		// 	products.GET("/:id", r.productHandler.Get)
		// 	products.PUT("/:id", r.productHandler.Update)
		// 	products.DELETE("/:id", r.productHandler.Delete)
		//
		// 	products.GET("/:id/variants", r.productVariantHandler.ListByProduct)
		// 	products.GET("/:id/images", r.productImageHandler.ListByProduct)
		// 	products.GET("/:id/reviews", r.reviewHandler.ListByProduct)
		// products.GET("/:id/attributes", r.productAttributeHandler.ListByProduct)
		// products.POST("/:id/attributes", r.productAttributeHandler.AddValues)
		// products.DELETE("/:id/attributes/:attribute_value_id", r.productAttributeHandler.RemoveValue)
		// }
		//
		// attributes := api.Group("/attributes")
		// {
		// 	attributes.GET("", r.attributeHandler.List)
		// 	attributes.POST("", r.attributeHandler.Create)
		// 	attributes.GET("/:id", r.attributeHandler.Get)
		// 	attributes.PUT("/:id", r.attributeHandler.Update)
		// 	attributes.DELETE("/:id", r.attributeHandler.Delete)
		//
		// 	attributes.GET("/:id/values", r.attributeValueHandler.ListByAttribute)
		// 	attributes.POST("/:id/values", r.attributeValueHandler.Create)
		// }
		//
		// payments := api.Group("/payment")
		// {
		// 	payments.GET("", r.paymentHandler.Calculate)
		// 	payments.POST("/coupon", r.paymentHandler.ApplyCoupon)
		// 	payments.GET("/methods", r.paymentHandler.ListMethods)
		// }
		//
		// orders := api.Group("/orders")
		// {
		// 	orders.GET("", r.orderHandler.List)
		// 	orders.POST("", r.orderHandler.Create)
		// 	orders.GET("/:id", r.orderHandler.Get)
		// 	orders.PUT("/:id", r.orderHandler.UpdateStatus)
		// 	orders.DELETE("/:id", r.orderHandler.Delete)
		// 	orders.GET("/:id/items", r.orderItemHandler.ListByOrder)
		// }
		//
		// returns := api.Group("/returns")
		// {
		// 	returns.GET("", r.returnHandler.List)
		// 	returns.POST("", r.returnHandler.Create)
		// 	returns.GET("/:id", r.returnHandler.Get)
		// 	returns.PUT("/:id", r.returnHandler.UpdateStatus)
		// }
		//
		// refunds := api.Group("/refunds")
		// {
		// 	refunds.GET("", r.refundHandler.List)
		// 	refunds.GET("/:id", r.refundHandler.Get)
		// }
		//
		// reviews := api.Group("/reviews")
		// {
		// 	reviews.GET("", r.reviewHandler.List)
		// 	reviews.POST("", r.reviewHandler.Create)
		// 	reviews.GET("/:id", r.reviewHandler.Get)
		// 	reviews.PUT("/:id", r.reviewHandler.Update)
		// 	reviews.DELETE("/:id", r.reviewHandler.Delete)
		// }
	}
}
