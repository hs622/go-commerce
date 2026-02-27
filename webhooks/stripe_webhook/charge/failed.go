package charge

import (
	"encoding/json"
	"fmt"

	"github.com/stripe/stripe-go/v84"
)

func Failed(event stripe.Event) {
	var charge stripe.Charge

	if err := json.Unmarshal(event.Data.Raw, &charge); err != nil {
		fmt.Printf("failed to parse stripe.charge.failed data: %s\n", err.Error())
		return
	}

	fmt.Println(charge)
}
