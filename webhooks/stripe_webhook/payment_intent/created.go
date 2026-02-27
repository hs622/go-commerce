package paymentintent

import (
	"encoding/json"
	"fmt"

	"github.com/stripe/stripe-go/v84"
)

func Created(event stripe.Event) {
	var paymentintent stripe.PaymentIntent

	if err := json.Unmarshal(event.Data.Raw, &paymentintent); err != nil {
		fmt.Printf("failed to parse stripe.paymentintent.created data: %v\n", err)
		return
	}

	fmt.Println(paymentintent)
}
