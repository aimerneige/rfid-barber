package model

import "gorm.io/gorm"

// Card consumer's vip card
type Card struct {
	gorm.Model
	Name       string
	Phone      string `gorm:"unique"`
	RFIDCardID string `gorm:"unique;column:rfid_card_id"`
	Balance    uint
}
