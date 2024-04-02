package models

import (
	"gorm.io/gorm"
)

type Coffee struct {
	gorm.Model
	Name        string   `gorm:"not null unique " ;json:"name"`
	Description string   `gorm:"not null" ;json:"description"`
	Favourite   string   `gorm:"not null" ;json:"favourite"`
	Contents    []string `gorm:"not null" ;json:"contents"`
	Cost        int64    `gorm:"not null" ;json:"cost"`
}
