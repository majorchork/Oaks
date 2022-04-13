package models

type Status struct {
	Pending   bool `json:"pending,omitempty"`
	Canceled  bool `json:"canceled,omitempty"`
	Completed bool `json:"completed,omitempty"`
}

type Order struct {
	Buyer
	Products        []Product `json:"products" gorm:"product"`
	OrderId         uint      `json:"order-id" gorm:"primarykey, autoincrement"`
	Quantity        int       `json:"quantity" gorm:"quantity"`
	DeliveryAddress string    `json:"delivery-address" gorm:"delivery-address"`
	TotalCost       int       `json:"total_cost" gorm:"total_cost"`
	Status
}
