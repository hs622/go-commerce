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

func (r *ProductRepository) InitIndexes(ctx context.Context) error {
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

func (r *ProductRepository) Create(ctx context.Context, product *schemas.Product) error {
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

func (r *ProductRepository) FetchProdcutById(ctx context.Context, productId string, product *schemas.Product) error {
	err := r.collection.FindOne(ctx, bson.M{"product_id": productId}).Decode(product)
	if err != nil {
		return err
	}
	return nil
}

func (r *ProductRepository) FetchProductsWithQuery(ctx context.Context, products *[]schemas.Product, url *url.URL) error {

	var filter bson.D
	var findOptions options.FindOptionsBuilder

	database.QParamsConvertIntoFindOptions(&findOptions, url)
	database.QParamsConvertIntoFilters(&filter, url)

	cursor, err := r.collection.Find(ctx, filter, &findOptions)
	if err != nil {
		return err
	}

	if err := cursor.All(ctx, products); err != nil {
		log.Println(err)
	}
	return nil
}

func (r *ProductRepository) Update(ctx context.Context, product *schemas.Product, productId string) error {

	product.Updated_at = time.Now()

	// insert modified fields

	return nil
}
