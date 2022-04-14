package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name            string `json:"name" gorm:"name"`
	Price           int    `json:"price" gorm:"price"`
	Quantity        int    `json:"quantity" gorm:"quantity"`
	ProductCategory string `json:"product_category" gorm:"product_category"`
}
