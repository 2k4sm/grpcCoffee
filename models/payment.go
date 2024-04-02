package models

import (
	"time"

	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model
	Payer string    `gorm:"not null" ;json:"payer"`
	Cost  int64     `gorm:"not null" ;json:"cost"`
	Date  time.Time `gorm:"not null" ;json:"date"`
}
