package middleware

import "github.com/gin-gonic/gin"

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

		ctx.Next()
	}
}
