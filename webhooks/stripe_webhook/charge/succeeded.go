package charge

import (
	"encoding/json"
	"fmt"

	"github.com/stripe/stripe-go/v84"
)

func Succeeded(event stripe.Event) {
	var charge stripe.Charge

	if err := json.Unmarshal(event.Data.Raw, &charge); err != nil {
		fmt.Printf("failed to parse stripe.charge.succeeded data: %s\n", err.Error())
		return
	}

	fmt.Println(charge)
}
