package config

import (
	"gorm.io/driver/sqlserver"
	_ "gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase(connectionString string) {
	database, err := gorm.Open(sqlserver.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}
	DB = database
}

func GetDB() *gorm.DB {
	return DB
}
