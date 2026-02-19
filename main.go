package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/hs622/ecommerce-cart/configuration"
	"github.com/hs622/ecommerce-cart/middleware"
	"github.com/hs622/ecommerce-cart/routes"
	"github.com/hs622/ecommerce-cart/utils"
	"github.com/hs622/ecommerce-cart/utils/validation"
)

type variables struct {
	port     string
	mongoURI string
	dbName   string
}

func GetEnvVariables() *variables {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		mongoURI = "mongodb://localhost:27017"
	}

	dbName := os.Getenv("DATABASE_NAME")
	if dbName == "" {
		dbName = "__testing_ase23oe_go"
	}

	return &variables{
		port:     port,
		mongoURI: mongoURI,
		dbName:   dbName,
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
