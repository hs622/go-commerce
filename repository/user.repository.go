package repository

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/hs622/ecommerce-cart/schemas"
	"github.com/hs622/ecommerce-cart/utils/tokens"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type UserRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(db *mongo.Database) *UserRepository {
	return &UserRepository{
		collection: db.Collection("users"),
	}
}

func (h *UserRepository) RegisterIndexesForUser(ctx context.Context) error {

	indexModel := mongo.IndexModel{
		Keys:    bson.D{{Key: "user_id", Value: 1}},
		Options: options.Index().SetUnique(true),
	}

	_, err := h.collection.Indexes().CreateOne(ctx, indexModel)
	if err != nil {
		return err
	}

	return nil
}

func (h *UserRepository) New(ctx context.Context, user *schemas.CreateUserRequest) error {

	hash, err := tokens.HashingPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = hash
	user.UserID = uuid.NewString()
	user.CreatedAt = time.Now()

	if _, err := h.collection.InsertOne(ctx, &user); err != nil {
		return err
	}

	return nil
}
