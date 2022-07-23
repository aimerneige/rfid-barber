package handler

import (
	"barber-server/internal/api/response"

	"github.com/gin-gonic/gin"
)

// NotFound 404 page not found
func NotFound(c *gin.Context) {
	response.NotFound(c, nil, "404 page not found")
}
