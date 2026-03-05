package repository

import (
	"context"
	"fmt"
	"net/url"
	"time"

	"github.com/google/uuid"
	"github.com/hs622/ecommerce-cart/schemas"
	"github.com/hs622/ecommerce-cart/utils/database"
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

	tt := time.Now()
	hash, err := tokens.HashingPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = hash
	user.UserID = uuid.NewString()
	user.CreatedAt = &tt

	if _, err := h.collection.InsertOne(ctx, &user); err != nil {
		return err
	}
	user.Password = ""
	return nil
}

func (h *UserRepository) FetchUsers(ctx context.Context, url *url.URL, users *[]schemas.CreateUserRequest) error {

	opts, err := database.FindOptionsParams(url)
	if err != nil {
		return err
	}

	cursor, err := h.collection.Find(ctx, bson.D{}, opts)
	if err != nil {
		return err
	}
	defer cursor.Close(ctx)

	// for debbuging purpose only
	// for cursor.Next(ctx) {
	// 	fmt.Println(cursor.Current)
	// }

	if err := cursor.All(ctx, users); err != nil {
		return err
	}
	return nil
}

func (h *UserRepository) FetchUser(ctx context.Context, user *schemas.CreateUserRequest, opts database.UserFetchOptions) error {

	param, err := database.FindOneOptionsParams(opts.Url)
	filter, err := database.FindOneOptionsFilters(opts)
	if err != nil {
		return err
	}

	if err := h.collection.FindOne(ctx, filter, param).Decode(&user); err != nil {
		return fmt.Errorf("Error fetching users: %v", err)
	}

	return nil
}

func (h *UserRepository) UpdateUser(
	ctx context.Context, userId string, field *schemas.PatchUserRequest, user *schemas.CreateUserRequest,
) error {

	tt := time.Now()
	filter := bson.D{{Key: "user_id", Value: userId}}
	opts := options.FindOneAndUpdate().SetProjection(bson.D{{Key: "password", Value: 0}}).SetReturnDocument(options.After)

	field.UpdatedAt = &tt
	if err := h.collection.FindOneAndUpdate(
		ctx, filter, bson.D{{Key: "$set", Value: &field}}, opts).Decode(&user); err != nil {
		return fmt.Errorf("UpdateUser: %w", err)
	}

	return nil
}

func (h *UserRepository) DeleteUser(ctx context.Context, userId string) error { return nil }

func (h *UserRepository) SuspendUser(ctx context.Context, userId string) error { return nil }
