package middlewares

import (
	"api/auth"
	"api/r"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.Request.Header.Get("Authorization")
		if authHeader == "" {
			ctx.JSON(http.StatusUnauthorized, r.Error[string]("authorization header is required"))
			ctx.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			ctx.JSON(http.StatusUnauthorized, r.Error[string]("authorization header is invalid"))
			ctx.Abort()
			return
		}

		_, err := auth.ParseToken(parts[1])
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, r.Error[string]("authorization header is invalid"))
			ctx.Abort()
			return
		}

		ctx.Next()
	}

}
