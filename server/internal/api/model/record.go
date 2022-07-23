package model

import "gorm.io/gorm"

// Record card comsume record
type Record struct {
	gorm.Model
	CardID uint
	Time   string
	Count  int
}
