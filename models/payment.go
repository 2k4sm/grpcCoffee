package models

import (
	"time"

	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model
	UserID        uint      `gorm:"not null" ;json:"user_id,omitempty"`
	CoffeeHouseID uint      `gorm:"not null" ;json:"coffee_house_id,omitempty"`
	Cost          int64     `gorm:"not null" ;json:"cost,omitempty"`
	Date          time.Time `gorm:"not null" ;json:"date,omitempty"`
	Items         []Coffee  `gorm:"many2many:payment_items" ;json:"items,omitempty"`
}
