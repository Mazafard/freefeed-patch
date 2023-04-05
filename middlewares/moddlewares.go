package middleware

import (
	"feed/utils/freefeed"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FreeFeedMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := freefeed.TokenValid(ctx)
		if err != nil {
			ctx.String(http.StatusUnauthorized, "Unauthorized")
			ctx.Abort()
			return
		}
		ctx.Next()
	}

}
