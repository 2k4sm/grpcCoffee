package models

import "gorm.io/gorm"

type CoffeeHouse struct {
	gorm.Model
	Name             string    `gorm:"not null unique" ;json:"house_name"`
	UserCount        int64     `gorm:"not null default:0" ;json:"user_count"`
	TopCoffee        string    `gorm:"not null" ;json:"top_coffee"`
	Revenue          int64     `gorm:"not null default:0" ;json:"revenue"`
	AvailableCoffees []Coffee  `gorm:"many2many:coffee_house_coffees" ;json:"coffees"`
	Payments         []Payment `gorm:"foreignKey:CoffeeHouseID" ;json:"payments"`
}
