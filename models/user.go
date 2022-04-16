package models

type User struct {
	ID           uint   `json:"ID" gorm:"autoincrement"` // for storage
	Name         string `json:"name" gorm:"name"`
	Email        string `json:"email" gorm:"email"`
	Username     string `json:"Username" gorm:"Username"`
	Password     string `json:"password,omitempty" gorm:"-"`
	PasswordHash string `json:"-" gorm:"password-hash"`
	Address      string `json:"address" gorm:"address"`
}

type Buyer struct {
	User
	UserID  uint //this is the foreignkey
	BuyerID uint `json:"buyerID" gorm:"primarykey, autoincrement"`
}

type Seller struct {
	User
	UserID   uint      `gorm:"foreignkey"`
	SellerID uint      `json:"sellerID" gorm:"SellerID"`
	Product  []Product `json:"product" gorm:"product"`
}
