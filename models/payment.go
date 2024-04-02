package models

import (
	"time"

	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model
	PayerID uint      `gorm:"not null" ;json:"payer_id"`
	Cost    int64     `gorm:"not null" ;json:"cost"`
	Date    time.Time `gorm:"not null" ;json:"date"`
	Items   []Coffee  `json:"items"`
	HouseID uint      `json:"house_id"`
}
