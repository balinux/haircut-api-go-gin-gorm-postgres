package models

import "gorm.io/gorm"

type Haircut struct {
	gorm.Model
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Image       *string `json:"image"`
}
