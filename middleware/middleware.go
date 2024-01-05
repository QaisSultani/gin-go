package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authenticate(ctx *gin.Context) {
	if !(ctx.Request.Header.Get("Token") == "auth") {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"Message": "Token Not present",
		})
		return
	}
	ctx.Next()
}

// func Authenticate() gin.HandlerFunc {
// 	// now we can write custom logic to be applied before my middleware
// 	return func(ctx *gin.Context) {
// 		if !(ctx.Request.Header.Get("Token") == "auth") {
// 			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
// 				"Message": "Token Not present",
// 			})
// 			return
// 		}
// 		ctx.Next()
// 	}
// }

func AddHeader(ctx *gin.Context) {
	ctx.Writer.Header().Set("key", "value")
	ctx.Next()
}
