package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/hs622/ecommerce-cart/configuration"
	"github.com/hs622/ecommerce-cart/middleware"
	"github.com/hs622/ecommerce-cart/routes"
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
	env := GetEnvVariables()

	// database connection
	mongodb, err := configuration.DatabaseInit(env.mongoURI, env.dbName)
	if err != nil {
		fmt.Printf("Couldn't connect to the database during application startup: %s", err)
	}
	defer mongodb.Disconnect()

	// initialising repositories
	gin.SetMode(gin.ReleaseMode)

	http := gin.New()
	http.SetTrustedProxies([]string{"127.0.0.1"})

	http.Use(gin.Logger())
	http.Use(gin.Recovery())
	http.Use(middleware.CORSMiddleware())

	api := http.Group("/api")

	// common
	routes.ServerRoutes(api)

	// v1
	// routes.UserRoutes(http)
	// http.Use(middleware.Authentication())
	routes.ProductRoutes(api, mongodb.Database)
	// routes.CartRoutes(http)

	log.Fatal(http.Run(":" + env.port))

}
