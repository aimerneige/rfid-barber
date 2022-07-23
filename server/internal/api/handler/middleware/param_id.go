package middleware

import (
	"barber-server/internal/api/response"
	"barber-server/internal/pkg"

	"github.com/gin-gonic/gin"
)

// ParamIDCheckMiddleware middleware to check param arguments id
func ParamIDCheckMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		idU64, err := pkg.String2Uint64(id)
		if err != nil {
			response.BadRequest(c, err, "Given id is not a number.")
			c.Abort()
			return
		}

		if idU64 <= 0 {
			response.BadRequest(c, err, "Not a valid id.")
			c.Abort()
			return
		}

		c.Set("id", uint(idU64))
		c.Next()
	}
}
