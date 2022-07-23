package handler

import (
	"barber-server/internal/api/response"
	"time"

	"github.com/gin-gonic/gin"
)

// Index returns simple json to shows thant server is running
func Index(c *gin.Context) {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	response.OK(c, time.Now().In(loc).Format("2006-01-02 15:04:05"), "Server is running!")
}
