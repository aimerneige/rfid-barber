package service

import (
	"barber-server/internal/api/database"
	"barber-server/internal/api/model"
	"time"

	"gorm.io/gorm"
)

// RecordService service for record
type RecordService struct {
	DB *gorm.DB
}

// NewRecordService create a RecordService
func NewRecordService() RecordService {
	return RecordService{
		DB: database.GetDB(),
	}
}

// Create crate a record into database
func (r RecordService) Create(cardID uint, count int) error {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	timeString := time.Now().In(loc).Format("2006-01-02 15:04:05")
	record := model.Record{
		CardID: cardID,
		Time:   timeString,
		Count:  count,
	}
	return r.DB.Create(&record).Error
}

// GetAll get all record data
func (r RecordService) GetAll() ([]model.Record, error) {
	var records []model.Record
	err := r.DB.Find(&records).Error
	return records, err
}

// GetByID get record data by id
func (r RecordService) GetByID(id uint) (model.Record, error) {
	var record model.Record
	err := r.DB.Find(&record, id).Error
	return record, err
}
