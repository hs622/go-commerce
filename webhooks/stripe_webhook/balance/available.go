package balance

import (
	"encoding/json"
	"fmt"

	"github.com/stripe/stripe-go/v84"
)

func Available(event stripe.Event) {
	var account stripe.Balance

	if err := json.Unmarshal(event.Data.Raw, &account); err != nil {
		fmt.Printf("failed to parse stripe.balance.available data: %s\n", err)
		return
	}

	fmt.Print(account)
}
