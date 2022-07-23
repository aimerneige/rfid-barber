package service

import (
	"barber-server/internal/api/database"
	"barber-server/internal/api/model"

	"gorm.io/gorm"
)

// CardService service for card
type CardService struct {
	DB *gorm.DB
}

// NewCardService create a CardService
func NewCardService() CardService {
	return CardService{
		DB: database.GetDB(),
	}
}

// Create create a card into database
func (c CardService) Create(name, phone, rfid string, balance uint) error {
	card := model.Card{
		Name:       name,
		Phone:      phone,
		RFIDCardID: rfid,
		Balance:    balance,
	}
	return c.DB.Create(&card).Error
}

// GetAll get all card data
func (c CardService) GetAll() ([]model.Card, error) {
	var cards []model.Card
	err := c.DB.Find(&cards).Error
	return cards, err
}

// GetByID get card data by id
func (c CardService) GetByID(id uint) (model.Card, error) {
	var card model.Card
	err := c.DB.Find(&card, id).Error
	return card, err
}

// GetByName get card data by name
func (c CardService) GetByName(name string) ([]model.Card, error) {
	var cards []model.Card
	err := c.DB.Where("name = ?", name).Find(&cards).Error
	return cards, err
}

// GetByPhone get card data by phone
func (c CardService) GetByPhone(phone string) (model.Card, error) {
	var card model.Card
	err := c.DB.Where("phone = ?", phone).First(&card).Error
	return card, err
}

// GetByRFID get card data by rfid
func (c CardService) GetByRFID(rfid string) (model.Card, error) {
	var card model.Card
	err := c.DB.Where("rfid_card_id = ?", rfid).First(&card).Error
	return card, err
}

// UpdateNameByID update card name by id
func (c CardService) UpdateNameByID(id uint, name string) error {
	return c.DB.First(&model.Card{}, id).Updates(map[string]interface{}{
		"name": name,
	}).Error
}

// UpdatePhoneByID update card phone by id
func (c CardService) UpdatePhoneByID(id uint, phone string) error {
	return c.DB.First(&model.Card{}, id).Updates(map[string]interface{}{
		"phone": phone,
	}).Error
}

// UpdateRFIDByID update card rfid by id
func (c CardService) UpdateRFIDByID(id uint, rfid string) error {
	return c.DB.First(&model.Card{}, id).Updates(map[string]interface{}{
		"rfid_card_id": rfid,
	}).Error
}

// UpdateBalanceByID update card balance by id
func (c CardService) UpdateBalanceByID(id, balance uint) error {
	return c.DB.First(&model.Card{}, id).Updates(map[string]interface{}{
		"balance": balance,
	}).Error
}

// UpdateBalanceWithCountByID update card balance with count by id
func (c CardService) UpdateBalanceWithCountByID(id uint, count int) error {
	return c.DB.First(&model.Card{}, id).Updates(map[string]interface{}{
		"balance": gorm.Expr("balance + ?", count),
	}).Error
}

// UpdateDataByID update card data by id
func (c CardService) UpdateDataByID(id uint, name, phone, rfid string, balance uint) error {
	return c.DB.First(&model.Card{}, id).Updates(map[string]interface{}{
		"name":         name,
		"phone":        phone,
		"rfid_card_id": rfid,
		"balance":      balance,
	}).Error
}

// DeleteByID delete card by id
func (c CardService) DeleteByID(id uint) error {
	return c.DB.Delete(&model.Card{}, id).Error
}
