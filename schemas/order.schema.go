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
	ProductID string  `bson:"product_id" json:"product_id" validate:"required"`
	Quantity  int64   `bson:"quantity" json:"quantity" validate:"required,number,gt=0,format=1"`
	Price     float64 `bson:"price" json:"price" validate:"number,gt=0,format=158.25"`
}

type CreateOrderRequest struct {
	Id                 bson.ObjectID            `bson:"_id,omitempty" json:"-"`
	OrderID            string                   `bson:"order_id,index" json:"order_id,omitempty" binding:"omitempty"`
	UserID             string                   `bson:"user_id" json:"user_id"`
	ShipAddressID      string                   `bson:"ship_address_id" json:"ship_address_id"`
	Items              []*Item                  `bson:"items" json:"items" validate:"required"`
	DiscountAmount     float64                  `bson:"discount_amount,omitempty" json:"discount_amount" validate:"number,gt=0,format=158.25"`
	DiscountCode       string                   `bson:"discount_code,omitempty" json:"discount_code" validate:"min=4,max=20"`
	TotalPrice         float64                  `bson:"total_price" json:"total_price" validate:"number,gtl=0,format=158.25"`
	VAT                float64                  `bson:"vat" json:"vat" validate:"number,gt=0,format=158.25"`
	PaymentMethod      *constants.PaymentMethod `bson:"payment_method,omitempty" json:"payment_method,omitempty"`
	PaymentReferenceID string                   `bson:"payment_reference_id,omitempty" json:"payment_reference_id,omitempty"`
	Status             *OrderStatus             `bson:"status,omitempty" json:"status,omitempty"`
	OrderedAt          time.Time                `bson:"ordered_at" json:"ordered_at"`
}

// update order request struct
type PatchOrderRequest struct{}
