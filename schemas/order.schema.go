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
	ProductID string  `bson:"product_id" json:"product_id,omitempty" binding:"uuid4"`
	Quantity  int64   `bson:"quantity" json:"quantity,omitempty" binding:"required"`
	Price     float64 `bson:"price" json:"price,omitempty" binding:"required"`
}

type CreateOrderRequest struct {
	Id                 bson.ObjectID            `bson:"_id,omitempty" json:"-"`
	OrderID            string                   `bson:"order_id,index" json:"order_id,omitempty" binding:"omitempty,uuid4"`
	UserID             string                   `bson:"user_id" json:"user_id,omitempty" binding:"omitempty,uuid4"`
	ShipAddressID      string                   `bson:"ship_address_id" json:"ship_address_id,omitempty" binding:"omitempty,uuid4"`
	Items              []Item                   `bson:"items" json:"items" binding:"required,min=1,dive"`
	DiscountAmount     *float64                 `bson:"discount_amount" json:"discount_amount,omitempty"`
	DiscountCode       string                   `bson:"discount_code" json:"discount_code,omitempty"`
	TotalPrice         float64                  `bson:"total_price" json:"total_price" binding:"required_if_item_is_available"`
	VAT                float64                  `bson:"vat" json:"vat,omitempty" binding:"required_if_item_is_available"`
	PaymentMethod      *constants.PaymentMethod `bson:"payment_method" json:"payment_method,omitempty"`
	PaymentReferenceID string                   `bson:"payment_reference_id" json:"payment_reference_id,omitempty" binding:"omitempty,uuid4"`
	Status             *OrderStatus             `bson:"status" json:"status,omitempty"`
	OrderedAt          time.Time                `bson:"ordered_at" json:"ordered_at"`
	UpdatedAt          time.Time                `bson:"updated_at" json:"updated_at"`
	DeletedAt          *time.Time               `bson:"deleted_at,omitempty" json:"deleted_at,omitempty"`
}

// update order request struct
type PatchOrderRequest struct {
	Items          []Item      `bson:"items" json:"items" binding:"required,min=1,dive"`
	TotalPrice     float64     `bson:"total_price" json:"total_price" binding:"required_if_item_is_available"`
	VAT            float64     `bson:"vat" json:"vat" binding:"required_if_item_is_available"`
	DiscountAmount float64     `bson:"discount_amount" json:"discount_amount,omitempty"`
	DiscountCode   string      `bson:"discount_code" json:"discount_code,omitempty"`
	Status         OrderStatus `bson:"status" json:"status"`
	UpdatedAt      time.Time   `bson:"updated_at" json:"updated_at"`
	DeletedAt      *time.Time  `bson:"deleted_at,omitempty" json:"deleted_at,omitempty"`
}
