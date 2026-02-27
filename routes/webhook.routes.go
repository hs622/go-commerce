package routes

import (
	"github.com/gin-gonic/gin"
	stripewebhook "github.com/hs622/ecommerce-cart/webhooks/stripe_webhook"
)

func WebhookRoute(rg *gin.RouterGroup) {

	wb := rg.Group("/v1")

	wb.POST("", stripewebhook.StripeWebHookChecker)
}
