package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserID     uint
	Total      float64
	Status     string
	OrderItems []OrderItem
}

type OrderItem struct {
	gorm.Model
	OrderID   uint
	Product   Product
	ProductID uint
	Quantity  int
	Price     float64
}
