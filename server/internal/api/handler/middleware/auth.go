package middleware

import (
	"barber-server/internal/api/response"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// AuthMiddleware middleware to check auth info
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		key := viper.GetString("secret.key")

		if auth != key {
			response.Forbidden(c, nil, "Not a right key.")
			c.Abort()
			return
		}

		c.Next()
	}
}
