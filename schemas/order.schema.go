package schemas

import (
	"time"

	"github.com/hs622/ecommerce-cart/constants"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type OrderStatus struct {
	OrderHold     bool                    `bson:"order" json:"order"`
	PaymemtStatus constants.PaymemtStatus `bson:"payment_status" json:"payment_status"`
}

type Item struct {
	ProductID string  `bson:"product_id" json:"product_id,omitempty" binding:"required"`
	Quantity  int64   `bson:"quantity" json:"quantity,omitempty" binding:"required"`
	Price     float64 `bson:"price" json:"price,omitempty" binding:"required"`
}

type CreateOrderRequest struct {
	Id                 bson.ObjectID            `bson:"_id,omitempty" json:"-"`
	OrderID            string                   `bson:"order_id,index" json:"order_id,omitempty"`
	UserID             string                   `bson:"user_id" json:"user_id,omitempty"`
	ShipAddressID      string                   `bson:"ship_address_id" json:"ship_address_id,omitempty"`
	Items              []Item                   `bson:"items" json:"items" binding:"required,min=1,dive"`
	DiscountAmount     *float64                 `bson:"discount_amount" json:"discount_amount,omitempty"`
	DiscountCode       string                   `bson:"discount_code" json:"discount_code,omitempty"`
	TotalPrice         *float64                 `bson:"total_price" json:"total_price" binding:"required_if=,min=0"`
	VAT                *float64                 `bson:"vat" json:"vat,omitempty" binding:"required,min=0"`
	PaymentMethod      *constants.PaymentMethod `bson:"payment_method" json:"payment_method,omitempty"`
	PaymentReferenceID string                   `bson:"payment_reference_id" json:"payment_reference_id,omitempty"`
	Status             *OrderStatus             `bson:"status" json:"status,omitempty"`
	OrderedAt          time.Time                `bson:"ordered_at" json:"ordered_at"`
}

// update order request struct
type PatchOrderRequest struct{}
