package schemas

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Product struct {
	Id          bson.ObjectID `bson:"_id,omitempty" json:"-"`
	ProductID   string        `bson:"product_id,index" json:"product_id,omitempty" binding:"omitempty"`
	Title       string        `bson:"title,omitempty" json:"title,omitempty" binding:"required,min=3,max=50" msg:"Invalid title!"`
	Description string        `bson:"description" json:"description,omitempty" binding:"required,min=10,max=500"`
	Price       float64       `bson:"price" json:"price,omitempty" binding:"required"`
	Category    string        `bson:"category" json:"category,omitempty"`
	SKU         string        `bson:"sku,omitempty" json:"sku,omitempty"`
	Stock       int           `bson:"stock,omitempty" json:"stock,omitempty"`
	Images      []string      `bson:"images,omitempty" json:"images,omitempty"`
	Tags        []string      `bson:"tags,omitempty" json:"tags,omitempty"`
	IsActive    bool          `bson:"is_active,omitempty" json:"publish,omitempty"`
	Created_at  time.Time     `bson:"created_at" json:"created_at"`
	Updated_at  time.Time     `bson:"updated_at" json:"updated_at,omitempty"`
}

// Options: MetaData for pagination, filtering, etc.
type ProductFilter struct {
	Category string
}

// Rating      *uint8        `json:"rating" bson:"rating"`
