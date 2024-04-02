package models

import "gorm.io/gorm"

type CoffeeHouse struct {
	gorm.Model
	Name      string `gorm:"not null unique" ;json:"name"`
	UserCount int64  `gorm:"not null default:0" ;json:"user_count"`
	Revenue   int64  `gorm:"not null default:0" ;json:"revenue"`
}
