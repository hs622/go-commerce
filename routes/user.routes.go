package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hs622/ecommerce-cart/handlers"
	"github.com/hs622/ecommerce-cart/repository"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func UserRouters(rg *gin.RouterGroup, db *mongo.Database) {

	repo := repository.NewUserRepository(db)
	userHandler := handlers.NewUserHandler(repo)

	user := rg.Group("/v1/users")

	user.POST("/create", userHandler.CreateUser)

	user.GET("/:userId", userHandler.GetUser)
	user.GET("/users", userHandler.GetUsers)
}
