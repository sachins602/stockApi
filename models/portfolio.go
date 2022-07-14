package models

import (
	"github.com/jinzhu/gorm"
)

type Portfolio struct {
	gorm.Model
	Username string  `json:"username" gorm:"foreign_key"`
	Scrip    string  `json:"scrip" gorm:"foreign_key"`
	Total    float64 `json:"total"`
	Price    float64 `json:"price"`
}

type PortfolioResponseForTable struct {
	gorm.Model
	Username  string  `json:"username"`
	Scrip     string  `json:"scrip"`
	Total     float64 `json:"total"`
	Price     float64 `json:"price"`
	LastPrice float64 `json:"lastPrice" gorm:"column:LastPrice"`
	Open      float64 `json:"open" gorm:"column:Open"`
}
