package handler

import (
	"barber-server/internal/api/response"
	"barber-server/internal/api/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// CreateCard create a newe card
func CreateCard(c *gin.Context) {
	type requestDto struct {
		Name    string `json:"name"`
		Phone   string `json:"phone"`
		RFID    string `json:"rfid"`
		Balance uint   `json:"balance"`
	}
	var dto requestDto
	if err := c.ShouldBindJSON(&dto); err != nil {
		response.BadRequest(c, err, "Fail to bind json.")
		return
	}
	if dto.Name == "" || dto.Phone == "" || dto.RFID == "" || dto.Balance < 0 {
		response.BadRequest(c, nil, "Not a valid data.")
		return
	}
	cardService := service.NewCardService()
	if err := cardService.Create(dto.Name, dto.Phone, dto.RFID, dto.Balance); err != nil {
		response.InternalServerError(c, err, "Fail to create card.")
		return
	}
	response.Created(c, nil, "Successful.")
}

// GetAllCard get all card data
func GetAllCard(c *gin.Context) {
	cardService := service.NewCardService()
	cards, err := cardService.GetAll()
	if err != nil {
		response.InternalServerError(c, err, "Fail to get all card data.")
		return
	}
	response.OK(c, cards, "Successful.")
}

// GetCardByID get card data by id
func GetCardByID(c *gin.Context) {
	id := c.GetUint("id")
	cardService := service.NewCardService()
	card, err := cardService.GetByID(id)
	if err != nil {
		response.InternalServerError(c, err, "Fail to get card data.")
		return
	}
	if card.ID == 0 {
		response.NotFound(c, nil, "Not Found")
		return
	}
	response.OK(c, card, "Successful.")
}

// GetCardByName get card data by name
func GetCardByName(c *gin.Context) {
	name := c.Param("name")
	cardService := service.NewCardService()
	cards, err := cardService.GetByName(name)
	if err != nil {
		response.InternalServerError(c, err, "Fail to get card data.")
		return
	}
	if len(cards) == 0 {
		response.NotFound(c, nil, "Not Found")
		return
	}
	response.OK(c, cards, "Successful.")
}

// GetCardByPhone get card data by phone
func GetCardByPhone(c *gin.Context) {
	phone := c.Param("phone")
	cardService := service.NewCardService()
	card, err := cardService.GetByPhone(phone)
	if err == gorm.ErrRecordNotFound {
		response.NotFound(c, nil, "Not Found.")
		return
	}
	if err != nil {
		response.InternalServerError(c, err, "Fail to get card data.")
		return
	}
	response.OK(c, card, "Successful.")
}

// GetCardByRFIDCardID get card data by rfid
func GetCardByRFIDCardID(c *gin.Context) {
	rfid := c.Param("rfid")
	cardService := service.NewCardService()
	card, err := cardService.GetByRFID(rfid)
	if err == gorm.ErrRecordNotFound {
		response.NotFound(c, nil, "Not Found.")
		return
	}
	if err != nil {
		response.InternalServerError(c, err, "Fail to get card data.")
		return
	}
	response.OK(c, card, "Successful.")
}

// UpdateCardDataByID update card by id
func UpdateCardDataByID(c *gin.Context) {
	type requestDto struct {
		Name    string `json:"name"`
		Phone   string `json:"phone"`
		RFID    string `json:"rfid"`
		Balance uint   `json:"balance"`
	}
	var dto requestDto
	if err := c.ShouldBindJSON(&dto); err != nil {
		response.BadRequest(c, err, "Fail to bind json.")
		return
	}
	id := c.GetUint("id")
	cardService := service.NewCardService()
	if err := cardService.UpdateDataByID(id, dto.Name, dto.Phone, dto.RFID, dto.Balance); err != nil {
		response.InternalServerError(c, err, "Fail to update card data.")
		return
	}
	response.OK(c, nil, "Successful.")
}

// UpdateBalanceByID update card balance by id
func UpdateBalanceByID(c *gin.Context) {
	type requestDto struct {
		Count int `json:"count"`
	}
	var dto requestDto
	if err := c.ShouldBindJSON(&dto); err != nil {
		response.BadRequest(c, err, "Fail to bind json,")
		return
	}
	id := c.GetUint("id")
	recordService := service.NewRecordService()
	if err := recordService.Create(id, dto.Count); err != nil {
		response.InternalServerError(c, err, "Fail to create record.")
		return
	}
	cardService := service.NewCardService()
	if err := cardService.UpdateBalanceWithCountByID(id, dto.Count); err != nil {
		response.InternalServerError(c, err, "Fail to update card data.")
		return
	}
	response.OK(c, nil, "Successful.")
}

// DeleteCardByID delete card by id
func DeleteCardByID(c *gin.Context) {
	cardService := service.NewCardService()
	id := c.GetUint("id")
	err := cardService.DeleteByID(id)
	if err != nil {
		response.InternalServerError(c, err, "Fail to delete card.")
		return
	}
	response.Deleted(c)
}
