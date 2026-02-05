package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hs622/ecommerce-cart/handlers"
	"github.com/hs622/ecommerce-cart/repository"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func ProductRoutes(rg *gin.RouterGroup, db *mongo.Database) {

	productRepo := repository.NewProductRepository(db)
	productHandlers := handlers.NewProductHandler(productRepo)

	// specific middleware declaration
	products := rg.Group("/v1/products")

	products.GET("", productHandlers.GetProdcutByQuery)
	products.GET("/:productId", productHandlers.GetProductById)
	products.POST("/create-product", productHandlers.CreateProduct)

	products.PATCH("/:productId", productHandlers.UpdateProduct)
	products.DELETE("/:productId", productHandlers.DeleteProduct)
	products.POST("/restore/:productId", productHandlers.RestoreProduct)
}
