package schemas

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type OrderStatus struct {
	OrderHold     bool `bson:"order" json:"order"`
	PaymemtStatus bool `bson:"payment_status" json:"payment_status"`
}

type CreateOrderRequest struct {
	Id                 bson.ObjectID `bson:"_id,omitempty" json:"-"`
	OrderID            string        `bson:"product_id,index" json:"product_id,omitempty" binding:"omitempty"`
	UserID             string        `bson:"user_id" json:"user_id"`
	TotalPrice         float64       `bson:"price" json:"price"`
	DiscountAmount     float64       `bson:"discount_amount" json:"discount_amount"`
	DiscountPrice      float64       `bson:"discount_price" json:"discount_price"`
	DiscountCode       string        `bson:"discount_code" json:"discount_code"`
	PaymentService     string        `bson:"payment_service" json:"payment_service"`
	PaymentReferenceId string        `bson:"payment_reference_id" json:"payment_reference_id"`
	Status             OrderStatus   `bson:"status" json:"status"`
	OrderedAt          time.Time     `bson:"ordered_at" json:"ordered_at"`
}
