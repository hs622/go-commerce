package repository

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type PaymentRepository struct {
	collection *mongo.Collection
}

func NewPaymentRepository(db *mongo.Database) *PaymentRepository {
	return &PaymentRepository{
		collection: db.Collection("payments"),
	}
}

func (r *PaymentRepository) InitIndexes(
	ctx context.Context) error {

	indexModel := mongo.IndexModel{
		Keys:    bson.D{{Key: "payment_id", Value: 1}},
		Options: options.Index().SetUnique(true),
	}

	_, err := r.collection.Indexes().CreateOne(ctx, indexModel)
	if err != nil {
		return fmt.Errorf("Failed to create unique index: %w", err)
	}

	return nil
}

func (r *PaymentRepository) CreatePayment(ctx context.Context)   {}
func (r *PaymentRepository) UpdatePayment(ctx context.Context)   {}
func (r *PaymentRepository) DeletePayment(ctx context.Context)   {}
func (r *PaymentRepository) RetrievePayment(ctx context.Context) {}
func (r *PaymentRepository) Payments(ctx context.Context)        {}
func (r *PaymentRepository) Payment(ctx context.Context)         {}
