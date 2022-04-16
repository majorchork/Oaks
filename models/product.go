package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Product struct {
	gorm.Model
	Name        string `json:"name" gorm:"name"`
	Price       int    `json:"price" gorm:"price"`
	Quantity    int    `json:"quantity" gorm:"quantity"`
	Description string `json:"description" gorm:"description"`
	Image       string `json:"image" gorm:"Image"`
}
