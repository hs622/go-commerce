package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

type OrderRepository struct {
	collection *mongo.Collection
}

func NewOrderRepository(db *mongo.Database) *OrderRepository {
	return &OrderRepository{
		collection: db.Collection("orders"),
	}
}

func (r *OrderRepository) CreateOrder(ctx *context.Context)  {}
func (r *OrderRepository) UpdateOrder(ctx *context.Context)  {}
func (r *OrderRepository) SuspendOrder(ctx *context.Context) {}
func (r *OrderRepository) FetchOrders(ctx *context.Context)  {}
func (r *OrderRepository) FetchOrder(ctx *context.Context)   {}
