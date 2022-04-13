package models

type Product struct {
	Id              int    `json:"id" gorm:"primarykey, autoincrement"`
	Name            string `json:"name" gorm:"name"`
	Price           int    `json:"price" gorm:"price"`
	Quantity        int    `json:"quantity" gorm:"quantity"`
	ProductCategory string `json:"product_category" gorm:"product_category"`
}
