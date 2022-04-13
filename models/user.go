package models

import "time"

type User struct {
	UserID       uint      `json:"userID" gorm:"primarykey, autoincrement"` // for storage
	Name         string    `json:"name" gorm:"name"`
	Email        string    `json:"email" gorm:"email"`
	Username     string    `json:"Username" gorm:"Username"`
	Password     string    `json:"password,omitempty" gorm:"-"`
	PasswordHash string    `json:"-" gorm:"password-hash"`
	Address      string    `json:"address" gorm:"address"`
	TimeCreated  time.Time `json:"timeCreated" gorm:"timeCreated"`
}

/*

 */
type Buyer struct {
	User
	BuyerId uint `json:"buyer_id" gorm:"primarykey, autoincrement"` // for purchases
}
type Seller struct {
	User
	SellerId uint `json:"seller_id" gorm:"primarykey, autoincrement"` // for purchases
	product  []Product
}
