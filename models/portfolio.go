package models

import (
	"github.com/jinzhu/gorm"
)

type Portfolio struct {
	gorm.Model
	Username string  `json:"username" gorm:"primaryKey"`
	Scrip    string  `json:"scrip" gorm:"primaryKey"`
	Total    float64 `json:"total"`
	Price    float64 `json:"price"`
}
