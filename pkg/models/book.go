package models

import (
	"mybookstore/pkg/config"

	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	DB = config.GetDB()
}

type Book struct {
	ID    uint   `gorm:"primaryKey" json:"id"`
	Title string `json:"title"`
}
