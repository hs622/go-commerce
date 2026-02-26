package repository

import (
	"context"
	"fmt"

	"github.com/stripe/stripe-go/v84"
	intent "github.com/stripe/stripe-go/v84/paymentintent"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type PaymentRepository struct {
	collection *mongo.Collection
}

type PaginationRange struct {
	limit int64
	skip  string
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

func (r *PaymentRepository) CreatePayment(ctx context.Context) (*stripe.PaymentIntent, error) {

	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(1588),
		Currency: stripe.String(stripe.CurrencyUSD),
		AutomaticPaymentMethods: &stripe.PaymentIntentAutomaticPaymentMethodsParams{
			Enabled: stripe.Bool(true),
		},
	}

	paymentIntent, err := intent.New(params)
	if err != nil {
		return nil, err
	}

	return paymentIntent, nil
}

func (r *PaymentRepository) UpdatePayment(ctx context.Context, intentId string) (*stripe.PaymentIntent, error) {

	params := &stripe.PaymentIntentParams{
		Amount: stripe.Int64(1800),
	}
	params.AddMetadata("orderId", "e0631fa2-a1e9-41fd-9f29-e1350c243725")

	paymentIntent, err := intent.Update(intentId, params)
	if err != nil {
		return nil, err
	}

	return paymentIntent, nil
}

func (r *PaymentRepository) DeletePayment(ctx context.Context)   {}
func (r *PaymentRepository) RetrievePayment(ctx context.Context) {}
func (r *PaymentRepository) Payments(ctx context.Context, p PaginationRange) *intent.Iter {

	params := &stripe.PaymentIntentListParams{}
	params.Limit = stripe.Int64(p.limit)
	params.StartingAfter = stripe.String(p.skip)

	return intent.List(params)
}

func (r *PaymentRepository) Payment(ctx context.Context, intentId string) (*stripe.PaymentIntent, error) {

	params := &stripe.PaymentIntentParams{}
	paymentIntent, err := intent.Get(intentId, params)

	if err != nil {
		return nil, err
	}

	return paymentIntent, nil
}
