package database

import (
	"barber-server/internal/api/model"

	"gorm.io/gorm"
)

// DB is a global variable for database
var DB *gorm.DB

// DBInterface is an interface for database
type DBInterface interface {
	InitDB(migrateDst ...interface{}) (*gorm.DB, error)
}

// InitDatabase init database
func InitDatabase(dbi DBInterface) {
	db, err := dbi.InitDB(
		&model.Card{},
		&model.Record{},
	)
	if err != nil {
		panic(err)
	}

	DB = db
}

// GetDB get database
func GetDB() *gorm.DB {
	return DB
}
