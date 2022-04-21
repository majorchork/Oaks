package models

type Status struct {
	Pending   bool `json:"pending,omitempty"`
	Canceled  bool `json:"canceled,omitempty"`
	Completed bool `json:"completed,omitempty"`
}

type Order struct {
	Buyer           Buyer
	User            User
	Product         Product
	UserID          uint      //this is the foreign key linked to the user
	ProductID       uint      //this is the foreign key linked to the product
	OrderID         uint      `json:"order-id" gorm:"primarykey, autoincrement"`
	Quantity        int       `json:"quantity" gorm:"quantity"`
	DeliveryAddress string    `json:"delivery-address" gorm:"delivery-address"`
	TotalCost       int       `json:"total_cost" gorm:"total_cost"`
	Products        []Product `json:"products" gorm:"products"`
	Status          `gorm:"embedded"`
	Notification    string
}

type Cart struct {
	id        uint   `gorm:"primary_key"`
	Name      string `json:"name" gorm:"name"`
	Price     int
	Quantity  int
	ProductID uint
	Buyer     Buyer
	BuyerID   uint
}
