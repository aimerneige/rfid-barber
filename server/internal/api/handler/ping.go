package handler

import (
	"barber-server/internal/api/response"
	"time"

	"github.com/gin-gonic/gin"
)

// Ping is a simple ping function
func Ping(c *gin.Context) {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	response.OK(c, gin.H{
		"ip":   c.ClientIP(),
		"time": time.Now().In(loc).Format("2006-01-02 15:04:05"),
	}, "pong")
}
