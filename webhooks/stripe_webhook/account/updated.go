package account

import (
	"encoding/json"
	"fmt"

	"github.com/stripe/stripe-go/v84"
)

func Updated(event stripe.Event) {
	var account stripe.Account

	if err := json.Unmarshal(event.Data.Raw, &account); err != nil {
		fmt.Printf("failed to parse stripe.account.available data: %s\n", err)
		return
	}

	fmt.Print(account)
}
