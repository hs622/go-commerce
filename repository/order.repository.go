package repository

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"slices"
	"time"

	"github.com/google/uuid"
	"github.com/hs622/ecommerce-cart/constants"
	"github.com/hs622/ecommerce-cart/schemas"
	"github.com/hs622/ecommerce-cart/utils"
	"github.com/hs622/ecommerce-cart/utils/database"
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
	order.UpdatedAt = time.Now()

	result, err := r.collection.InsertOne(ctx, order)
	if err != nil {
		utils.Error(err.Error())
		return err
	}

	order.Id = result.InsertedID.(bson.ObjectID)
	return nil
}

func (r *OrderRepository) UpdateSingleOrder(
	ctx context.Context,
	orderId string,
	editedFields *schemas.PatchOrderRequest,
	updatedProduct *schemas.CreateOrderRequest,
) error {

	filter := bson.D{{
		Key: "order_id", Value: orderId,
	}}
	options := options.FindOneAndUpdate().SetReturnDocument(options.After)

	editedFields.Status.OrderHold = true
	editedFields.Status.PaymemtStatus = constants.ORDER_PENDING
	editedFields.UpdatedAt = time.Now()

	if err := r.collection.FindOneAndUpdate(
		ctx, filter, bson.D{{Key: "$set", Value: &editedFields}}, options).Decode(&updatedProduct); err != nil {
		return err
	}

	return nil
}

func (r *OrderRepository) SuspendOrder(ctx context.Context, orderId string) error {

	var order schemas.CreateOrderRequest
	filter := bson.M{"order_id": orderId}

	if err := r.collection.FindOne(ctx, filter).Decode(&order); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return fmt.Errorf("order %s not found", orderId)
		}
		return fmt.Errorf("fetching order: %w", err)
	}

	if slices.Contains([]string{"completed"}, string(order.Status.PaymemtStatus)) {
		// extra-steps to reversing payment flow.
		// reversePaymentFlowHookFunction(ctx, &order)
	}

	if order.Status.PaymemtStatus == constants.ORDER_WITHDRAW {
		return nil
	}

	update := bson.D{{
		Key: "$set",
		Value: bson.M{
			"status.order":          false,
			"status.payment_status": constants.ORDER_WITHDRAW,
		},
	}}

	result, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("cancelling order: %w", err)
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("order %s not found during update.", orderId)
	}
	if result.ModifiedCount == 0 {
		return fmt.Errorf("order %s was not modified.", orderId)
	}

	return nil
}

func (r *OrderRepository) SoftDeleteOrder(ctx context.Context, orderId string, order *schemas.CreateOrderRequest) error {

	options := options.FindOneAndUpdate().SetReturnDocument(options.After)
	if err := r.collection.FindOneAndUpdate(
		ctx, bson.M{"order_id": orderId}, bson.D{{
			Key:   "$set",
			Value: bson.E{Key: "deleted_at", Value: time.Now()},
		}}, options).Decode(&order); err != nil {
		return err
	}

	return nil
}

func (r *OrderRepository) FetchOrders(ctx context.Context) error { return nil }

func (r *OrderRepository) FetchOrder(ctx context.Context, orderId string, order *schemas.CreateOrderRequest, url url.URL) error {
	var findOneOptions options.FindOneOptionsBuilder

	database.FindOneOptionsParams(&findOneOptions, &url)

	if err := r.collection.FindOne(ctx, bson.M{"order_id": orderId}, &findOneOptions).Decode(&order); err != nil {
		return err
	}

	return nil
}
