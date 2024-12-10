package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name     string `json:"name" gorm:"unique"`
	Products []Product
}

type CategoryDTO struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
