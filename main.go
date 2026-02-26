package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/hs622/ecommerce-cart/configuration"
	"github.com/hs622/ecommerce-cart/constants"
	"github.com/hs622/ecommerce-cart/middleware"
	"github.com/hs622/ecommerce-cart/routes"
	"github.com/hs622/ecommerce-cart/utils"
	"github.com/hs622/ecommerce-cart/utils/validation"
)

type ApplicationVariables struct {
	port     string
	mongoURI string
	dbName   string
}

func GetEnvVariables() *ApplicationVariables {

	return &ApplicationVariables{
		port:     string(constants.SYS_PORT),
		mongoURI: string(constants.DATABASE_URI),
		dbName:   string(constants.DATABASE_NAME),
	}
}

// Bootstrap the application
func main() {

	// Envorinment Variables
	utils.Info("Environment Variable configuration.")
	env := GetEnvVariables()

	// database connection
	mongodb, err := configuration.DatabaseInit(env.mongoURI, env.dbName)
	if err != nil {
		utils.Error("Couldn't connect to the database during application startup")
	}
	defer mongodb.Disconnect()

	// Custom request validation registration
	validation.RegisterCustomRequestValidation()

	// initialising repositories
	gin.SetMode(gin.ReleaseMode)

	utils.Info("Server is booting up.")
	http := gin.New()
	utils.Info("Server is listing at port: http://127.0.0.1:" + env.port)
	http.SetTrustedProxies([]string{"127.0.0.1"})

	http.Use(gin.Logger())
	http.Use(gin.Recovery())
	http.Use(middleware.CORSMiddleware())
	http.Use(middleware.HandleRequestMiddleware())
	// http.Use(middleware.Authentication())
	api := http.Group("/api")

	routes.ServerRoutes(api)
	routes.ProductRoutes(api, mongodb.Database)
	routes.OrderRoutes(api, mongodb.Database)

	if err := http.Run(":" + env.port); err != nil {
		log.Fatal(err)
	}

}
