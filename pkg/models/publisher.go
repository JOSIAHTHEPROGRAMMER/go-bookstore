package models

type Publisher struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
}
