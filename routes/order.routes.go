package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hs622/ecommerce-cart/handlers"
	"github.com/hs622/ecommerce-cart/repository"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func OrderRoutes(rg *gin.RouterGroup, db *mongo.Database) {

	orderRepo := repository.NewOrderRepository(db)
	orderHandler := handlers.NewOrderHandler(orderRepo)

	// specific middleware declaration
	order := rg.Group("/v1/orders")

	order.GET("", orderHandler.GetOrders)
	order.GET("/:orderId", orderHandler.GetOrder)
	order.POST("/create-order", orderHandler.CreateOrder)
	order.PATCH("/:orderId", orderHandler.UpdateOrder)
	order.DELETE("/:orderId", orderHandler.UpdateOrder)

}
