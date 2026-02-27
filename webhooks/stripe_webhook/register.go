package stripewebhook

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hs622/ecommerce-cart/constants"
	"github.com/hs622/ecommerce-cart/webhooks/stripe_webhook/account"
	"github.com/hs622/ecommerce-cart/webhooks/stripe_webhook/balance"
	"github.com/hs622/ecommerce-cart/webhooks/stripe_webhook/charge"
	paymentintent "github.com/hs622/ecommerce-cart/webhooks/stripe_webhook/payment_intent"
	"github.com/stripe/stripe-go/v84"
)

func StripeWebHookChecker(ctx *gin.Context) {

	payload, err := ctx.GetRawData()
	if err != nil {
		fmt.Println("Error Receving webhook.")
		ctx.AbortWithStatus(400)
		return
	}

	// verifying stripe-signature for incoming webhooks.
	// If it's not working, check the API Version in the Stripe Dashboard.
	// It's mandatory in statically typed langnages, as stated in the Stripe documentation
	// reference link: https://docs.stripe.com/sdks/set-version
	event, err := stripe.ConstructEvent(
		payload,
		ctx.GetHeader("Stripe-Signature"),
		string(constants.STRIPT_WEBHOOK_SECRET),
	)

	if err != nil {
		fmt.Printf("Failed to pasre webhook body json: %v\n", err.Error())
		ctx.Status(http.StatusBadRequest)
		return
	}

	switch event.Type {
	case "payment_intent.created":
		paymentintent.Created(event)
	case "payment_intent.succeeded":
		paymentintent.Succeeded(event)
	case "payment_intent.failed":
		paymentintent.Failed(event)
	case "charge.succeeded":
		charge.Succeeded(event)
	case "charge.failed":
		charge.Failed(event)
	case "charge.refunded":
		charge.Refunded(event)
	case "balance.available":
		balance.Available(event)
	case "account.updated":
		account.Updated(event)
	default:
		fmt.Printf("Unhandle webhook type captured: %s\n", event.Type)
	}

	ctx.Status(http.StatusOK)
}
