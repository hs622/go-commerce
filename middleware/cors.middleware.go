package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hs622/ecommerce-cart/utils"
)

// high-ordered function
func CORSMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "localhost")
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		ctx.Writer.Header().Set("Content-Type", "application/json")

		if ctx.Request.Method == "Options" {
			ctx.AbortWithStatus(204)
			return
		}

		errors := make(map[string]string)
		if ctx.Request.Header.Get("Content-Type") != "application/json" {
			errors["header"] = "Please attach JSON header."
			utils.ErrorResponse(ctx, http.StatusBadRequest, "Insufficient request headers.", errors)
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
