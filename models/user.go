package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name          string        `gorm:"not null" ;json:"user_name"`
	Email         string        `gorm:"not null unique" ;json:"email"`
	Password      string        `gorm:"not null unique" ;json:"password"`
	LastOrder     string        `json:"last_order"`
	Favourite     string        `json:"favourite"`
	Revenue       string        `gorm:"default:0" ;json:"-"`
	Orders        []Payment     `gorm:"foreignKey:UserID" ;json:"orders"`
	VisitedHouses []CoffeeHouse `gorm:"many2many:user_visited_houses" ;json:"visited_houses"`
}
