package middleware

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"slices"

	"github.com/gin-gonic/gin"
	"github.com/hs622/ecommerce-cart/middleware/exemption"
	"github.com/hs622/ecommerce-cart/utils"
)

func HandleRequestMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		errors := make(map[string]string)
		// fmt.Printf("Going through %s:\n", utils.StringFuncName(0))
		methods := []string{
			http.MethodPost,
			http.MethodPatch,
			http.MethodPut,
		}

		fmt.Println(slices.Contains(exemption.APIExemptions, ctx.FullPath()))
		if slices.Contains(exemption.APIExemptions, ctx.FullPath()) { // high priority exemptions.
			ctx.Next()
			return
		} else if slices.Contains(methods, ctx.Request.Method) {
			if ctx.Request.Header.Get("Content-type") == "application/json" {
				body, err := ctx.GetRawData()
				if len(body) == 0 || err != nil {
					errors["error"] = "unable to parse json."

					utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request.", errors)
					ctx.Abort()
					return
				}

				ctx.Request.Body = io.NopCloser(bytes.NewBuffer(body))
			}
		}

		ctx.Next()
	}
}
