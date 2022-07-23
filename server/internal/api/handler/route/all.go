package route

import (
	"barber-server/internal/api/handler"
	"barber-server/internal/api/handler/middleware"

	"github.com/gin-gonic/gin"
)

// AllRouteCollection all of the route
func AllRouteCollection(r *gin.Engine) *gin.Engine {
	r.GET("", handler.Index)
	r.GET("/ping", handler.Ping)

	r.Use(middleware.AuthMiddleware())

	card := r.Group("/card")
	card.GET("", handler.GetAllCard)
	card.POST("", handler.CreateCard)
	cardID := card.Group("/:id")
	cardID.Use(middleware.ParamIDCheckMiddleware())
	cardID.GET("", handler.GetCardByID)
	cardID.POST("", handler.UpdateCardDataByID)
	cardID.POST("/balance", handler.UpdateBalanceByID)
	cardID.DELETE("", handler.DeleteCardByID)
	cardSearch := card.Group("/search")
	cardSearch.GET("/name/:name", handler.GetCardByName)
	cardSearch.GET("/phone/:phone", handler.GetCardByPhone)
	cardSearch.GET("/rfid/:rfid", handler.GetCardByRFIDCardID)

	record := r.Group("/record")
	record.GET("/", handler.GetAllRecord)
	record.POST("/", handler.CreateRecord)
	recordID := record.Group("/:id")
	recordID.Use(middleware.ParamIDCheckMiddleware())
	recordID.GET("", handler.GetRecordByID)

	r.NoRoute(handler.NotFound)

	return r
}
