package models

import (
	"time"

	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model
	UserID        uint      `gorm:"not null index"`
	CoffeeHouseID uint      `gorm:"not null index"`
	Cost          int64     `gorm:"not null"`
	Date          time.Time `gorm:"not null"`
	Items         []Coffee  `gorm:"many2many:payment_items"`
}
