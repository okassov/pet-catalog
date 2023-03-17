package entity

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Id          uint64 `gorm:"primary_key;column:id" json:"id"`
	Name        string `gorm:"column:name;UNIQUE" json:"name"`
	Description string `gorm:"column:description" json:"description"`
	ImageUrl    string `gorm:"column:image_url" json:"image_url"`
}
