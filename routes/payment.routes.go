package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hs622/ecommerce-cart/handlers"
	"github.com/hs622/ecommerce-cart/repository"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func PaymentRoutes(rg *gin.RouterGroup, db *mongo.Database) {

	paymentRepo := repository.NewPaymentRepository(db)
	paymentHandler := handlers.NewPaymentHandler(paymentRepo)

	payments := rg.Group("/v1/payments/")

	payments.GET("", paymentHandler.GetPaymentIntents)
	payments.GET("/:paymentId", paymentHandler.GetPaymentIntent)

	payments.POST("/generate", paymentHandler.CreatePaymentIntent)
	payments.PATCH("/:paymentId", paymentHandler.UpdatePaymentIntent)

}
