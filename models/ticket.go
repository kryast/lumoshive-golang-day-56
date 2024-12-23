package models

import (
	"gorm.io/gorm"
)

type Ticket struct {
	gorm.Model
	Name     string `json:"name"`
	Price    uint   `json:"price"`
	Quantity uint   `json:"quantity"`
}
