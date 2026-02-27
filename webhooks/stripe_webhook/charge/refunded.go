package charge

import (
	"encoding/json"
	"fmt"

	"github.com/stripe/stripe-go/v84"
)

func Refunded(event stripe.Event) {
	var charge stripe.Charge

	if err := json.Unmarshal(event.Data.Raw, &charge); err != nil {
		fmt.Printf("failed to parse stripe.charge.refunded data: %s\n", err.Error())
		return
	}

	fmt.Println(charge)

}
