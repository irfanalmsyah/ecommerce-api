package models

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	UserID    uint
	CartItems []CartItem
}

type CartItem struct {
	gorm.Model
	CartID    uint
	Product   Product
	ProductID uint
	Quantity  int `json:"quantity"`
}

type CartItemDTO struct {
	ID       uint       `json:"id"`
	CartID   uint       `json:"cart_id"`
	Quantity int        `json:"quantity"`
	Product  ProductDTO `json:"product"`
}

type CartDTO struct {
	ID        uint          `json:"id"`
	UserID    uint          `json:"user_id"`
	CartItems []CartItemDTO `json:"cart_items"`
}
