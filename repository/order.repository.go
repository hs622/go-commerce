package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/hs622/ecommerce-cart/constants"
	"github.com/hs622/ecommerce-cart/schemas"
	"github.com/hs622/ecommerce-cart/utils"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type OrderRepository struct {
	collection *mongo.Collection
}

func NewOrderRepository(db *mongo.Database) *OrderRepository {
	return &OrderRepository{
		collection: db.Collection("orders"),
	}
}

func (r *OrderRepository) RegisterIndexesForOrder(ctx context.Context) error {

	indexModel := mongo.IndexModel{
		Keys:    bson.D{{Key: "order_id", Value: 1}},
		Options: options.Index().SetUnique(true),
	}

	_, err := r.collection.Indexes().CreateOne(ctx, indexModel)
	if err != nil {
		return fmt.Errorf("Failed to create unique index %w", err)
	}

	return nil
}

func (r *OrderRepository) CreateSingleOrder(ctx context.Context, order *schemas.CreateOrderRequest) error {

	var status schemas.OrderStatus
	status.OrderHold = true
	status.PaymemtStatus = constants.ORDER_PENDING

	order.Status = &status
	order.OrderID = uuid.NewString()
	order.OrderedAt = time.Now()

	result, err := r.collection.InsertOne(ctx, order)
	if err != nil {
		utils.Error(err.Error())
		return err
	}

	order.Id = result.InsertedID.(bson.ObjectID)
	return nil
}

func (r *OrderRepository) UpdateOrder(ctx *context.Context)  {}
func (r *OrderRepository) SuspendOrder(ctx *context.Context) {}
func (r *OrderRepository) FetchOrders(ctx *context.Context)  {}
func (r *OrderRepository) FetchOrder(ctx *context.Context)   {}
