package handler

import (
	"barber-server/internal/api/response"
	"barber-server/internal/api/service"

	"github.com/gin-gonic/gin"
)

// CreateRecord crate a new record
func CreateRecord(c *gin.Context) {
	type requestDto struct {
		ID    uint `json:"id"`
		Count int  `json:"count"`
	}
	var dto requestDto
	if err := c.ShouldBindJSON(&dto); err != nil {
		response.BadRequest(c, err, "Fail to bind json.")
		return
	}
	cardService := service.NewCardService()
	card, err := cardService.GetByID(dto.ID)
	if err != nil {
		response.InternalServerError(c, err, "Fail to get card info.")
		return
	}
	if card.ID != dto.ID {
		response.BadRequest(c, nil, "Card not found.")
		return
	}
	recordService := service.NewRecordService()
	if err := recordService.Create(dto.ID, dto.Count); err != nil {
		response.InternalServerError(c, err, "Fail to create record.")
		return
	}
	response.Created(c, nil, "Created.")
}

// GetAllRecord get all record data
func GetAllRecord(c *gin.Context) {
	recordService := service.NewRecordService()
	records, err := recordService.GetAll()
	if err != nil {
		response.InternalServerError(c, err, "Fail to get all record data.")
		return
	}
	response.OK(c, records, "Successful.")
}

// GetRecordByID get record data by id
func GetRecordByID(c *gin.Context) {
	id := c.GetUint("id")
	recordService := service.NewRecordService()
	record, err := recordService.GetByID(id)
	if err != nil {
		response.InternalServerError(c, err, "fail to get record data.")
		return
	}
	if record.ID == 0 {
		response.NotFound(c, nil, "Not Found.")
		return
	}
	response.OK(c, record, "Successful.")
}
