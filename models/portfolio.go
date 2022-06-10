package models

import (
	"github.com/jinzhu/gorm"
)

type Portfolio struct {
	gorm.Model
	Username string  `json:"username" gorm:"foreign_key"`
	Scrip    string  `json:"scrip" gorm:"foreign_key"`
	Type     string  `json:"type"`
	Total    float64 `json:"total"`
	Price    float64 `json:"price"`
}
