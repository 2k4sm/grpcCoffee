package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name          string `gorm:"not null" ;json:"name"`
	LastOrder     string `gorm:"not null" ;json:"last_order"`
	Favourite     string `gorm:"not null" ;json:"favourite"`
	revenue       string `gorm:"not null default:0"`
	lastOrderDate *time.Time
}
