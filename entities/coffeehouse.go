package models

import "gorm.io/gorm"

type CoffeeHouse struct {
	gorm.Model
	Name             string `gorm:"not null unique"`
	UserCount        int64  `gorm:"not null default:0 index"`
	TopCoffee        string
	Revenue          int64     `gorm:"not null default:0 index"`
	AvailableCoffees []Coffee  `gorm:"many2many:coffee_house_coffees"`
	Payments         []Payment `gorm:"foreignKey:CoffeeHouseID"`
}
