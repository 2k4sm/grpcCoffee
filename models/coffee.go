package models

import (
	"gorm.io/gorm"
)

type MultiString []string

type Coffee struct {
	gorm.Model
	Name        string      `gorm:"not null unique " ;json:"coffee_name"`
	Description string      `gorm:"not null" ;json:"description"`
	Origin      string      `gorm:"not null" ;json:"origin"`
	Contents    MultiString `gorm:"type:text not null" ;json:"contents"`
	Cost        int64       `gorm:"not null" ;json:"cost"`
}
