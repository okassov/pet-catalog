package entity

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Id          uint64 `gorm:"primary_key;column:id" json:"id" `
	Name        string `gorm:"column:name" json:"name" `
	Description string `gorm:"column:description" json:"description"`
	Price       uint64 `gorm:"column:price" json:"price"`
	Rating      uint8  `gorm:"column:rating" json:"rating"`
	ImageUrl    string `gorm:"column:image_url" json:"image_url"`
}
