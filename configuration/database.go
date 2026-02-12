package configuration

import (
	"context"
	"fmt"
	"time"

	"github.com/hs622/ecommerce-cart/repository"
	"github.com/hs622/ecommerce-cart/utils"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

type MongoDB struct {
	Client   *mongo.Client
	Database *mongo.Database
}

// Initialising Database session.
//
// @Param URI string
// @Param DatabaseName string
//
// @return &MongoDB
func DatabaseInit(uri, dbName string) (*MongoDB, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Create Client Options
	clientOptions := options.Client().ApplyURI(uri)

	// Connect to Database
	client, err := mongo.Connect(clientOptions)
	if err != nil {
		return nil, fmt.Errorf("Failed to connect to database: %w", err)
	}

	// Ping the database to verify connection
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, fmt.Errorf("Failed to ping database: %w", err)
	}

	utils.Info("Successfully establish a connection to the database.")
	DbConfiguration(ctx, client.Database(dbName))

	return &MongoDB{
		Client:   client,
		Database: client.Database(dbName),
	}, nil
}

// Discounting Database session.
//
// @Params m MongoDB
// @Return error if occurred
func (m *MongoDB) Disconnect() error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if err := m.Client.Disconnect(ctx); err != nil {
		return fmt.Errorf("%s", err)
	}
	utils.Info("MongoDB connection closed")
	return nil
}

// Database initial configuration
//
// @Param ctx Context
// @Param db mongo.Database
//
// @Return error if occurred
func DbConfiguration(ctx context.Context, db *mongo.Database) error {

	productRepo := repository.NewProductRepository(db)
	if err := productRepo.InitIndexes(ctx); err != nil {
		utils.Error("Failed to register product indexes")
		return fmt.Errorf("%s", err)
	}

	orderRepo := repository.NewOrderRepository(db)
	if err := orderRepo.RegisterIndexesForOrder(ctx); err != nil {
		utils.Error("Failed to register order indexes")
		return fmt.Errorf("%s", err)
	}

	utils.Info("Database indexes initialised.")
	return nil
}
