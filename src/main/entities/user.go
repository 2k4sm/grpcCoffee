package entities

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name          string `gorm:"not null"`
	Email         string `gorm:"not null unique"`
	Password      string `gorm:"not null unique"`
	LastOrder     string
	Favourite     string
	Revenue       string        `gorm:"default:0"`
	Orders        []Payment     `gorm:"foreignKey:UserID"`
	VisitedHouses []CoffeeHouse `gorm:"many2many:user_visited_houses"`
}
