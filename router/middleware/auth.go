package middleware

import (
	. "apiserver/handler"
	"apiserver/package/errors"
	"apiserver/package/token"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Parse JWT
		if _, err := token.ParseRequest(c); err != nil {
			SendResponse(c, errors.TokenInvalidError, nil)
			c.Abort()
			return
		}

		c.Next()
	}
}
