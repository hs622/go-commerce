package repository

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/hs622/ecommerce-cart/schemas"
	"github.com/hs622/ecommerce-cart/utils/database"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type ProductRepository struct {
	collection *mongo.Collection
}

func NewProductRepository(db *mongo.Database) *ProductRepository {
	return &ProductRepository{
		collection: db.Collection("products"),
	}
}

func (r *ProductRepository) InitIndexes(
	ctx context.Context) error {
	// Initialising Indexes

	indexModel := mongo.IndexModel{
		Keys:    bson.D{{Key: "product_id", Value: 1}},
		Options: options.Index().SetUnique(true),
	}

	_, err := r.collection.Indexes().CreateOne(ctx, indexModel)
	if err != nil {
		return fmt.Errorf("Failed to create unique index: %w", err)
	}

	return nil
}

func (r *ProductRepository) CreateSingleProduct(
	ctx context.Context,
	product *schemas.CreateProductRequest) error {
	product.Created_at = time.Now()
	product.Updated_at = time.Now()
	product.ProductID = uuid.NewString()
	product.Title = strings.ToLower(product.Title)

	// insertion
	result, err := r.collection.InsertOne(ctx, product)
	if err != nil {
		return err
	}

	product.Id = result.InsertedID.(bson.ObjectID)
	return nil
}

func (r *ProductRepository) FetchProdcutById(
	ctx context.Context,
	productId string,
	product *schemas.CreateProductRequest,
	url *url.URL) error {

	// var findOptions options.FindOneOptions

	// database.QParamsConvertIntoFindOptions(&findOptions, url)

	err := r.collection.FindOne(ctx, bson.M{"product_id": productId}).Decode(product)
	if err != nil {
		return err
	}
	return nil
}

func (r *ProductRepository) FetchProductsWithQuery(
	ctx context.Context,
	products *[]schemas.CreateProductRequest,
	url *url.URL) error {

	var filter bson.D
	var findOptions options.FindOptionsBuilder

	database.FindOptionsParams(&findOptions, url)
	database.FindOptionsFilters(&filter, url)

	cursor, err := r.collection.Find(ctx, &filter, &findOptions)
	if err != nil {
		return err
	}

	if err := cursor.All(ctx, products); err != nil {
		log.Println(err)
	}
	return nil
}

func (r *ProductRepository) UpdateSingleProduct(
	ctx context.Context,
	incomingProduct *schemas.PatchProductRequest,
	productId *string,
	outgoingProduct *schemas.CreateProductRequest) error {

	filter := bson.D{{
		Key: "product_id", Value: &productId,
	}}
	options := options.FindOneAndUpdate().SetReturnDocument(options.After)
	incomingProduct.Updated_at = time.Now()

	// insert modified fields
	if err := r.collection.FindOneAndUpdate(
		ctx, filter, bson.D{{Key: "$set", Value: &incomingProduct}}, options).Decode(&outgoingProduct); err != nil {
		return err
	}

	return nil
}
