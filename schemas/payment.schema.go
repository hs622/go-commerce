package schemas

type Stripe struct {
}

type CreatePaymentIntentRequest struct {
	Id        string `bson:"-" json:"-"`
	PaymentID string `json:"payment_id" binding:"uuid4"`
	UserID    string `json:"user_id" binding:"required,uuid4"`

	// stripe
	Stripe Stripe `bson:"stripe" json:"-"`
}

type PatchPaymentIntentRequest struct {
}
