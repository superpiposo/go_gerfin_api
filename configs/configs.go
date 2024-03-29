package configs

import (
	"github.com/superpiposo/go_gerfin_api/schema"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Init() error {
	db, _ = InitializePostgres()
	db.AutoMigrate(&schema.Usuarios{})
	return nil

}

func GetLogger(p string) *Logger {
	logger := NewLogger(p)
	return logger
}
