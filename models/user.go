package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name          string        `gorm:"not null" ;json:"name"`
	Email         string        `gorm:"not null unique" ;json:"email"`
	Password      string        `gorm:"not null unique" ;json:"password"`
	LastOrder     string        `json:"last_order"`
	Favourite     string        `json:"favourite"`
	revenue       string        `gorm:"default:0"`
	Orders        []Payment     `json:"orders"`
	VisitedHouses []CoffeeHouse `json:"visited_houses"`
}
