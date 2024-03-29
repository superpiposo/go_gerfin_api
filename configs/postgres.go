package configs

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializePostgres() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=admin dbname=go_gerfin_api port=7000 sslmode=disable TimeZone=America/Sao_Paulo"
	if db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{}); err != nil {
		return nil, err
	} else {
		return db, nil
	}
}
