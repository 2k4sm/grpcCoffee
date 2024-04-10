package models

import (
	"gorm.io/gorm"
)

type MultiString []string

type Coffee struct {
	gorm.Model
	Name        string      `gorm:"not null unique index"`
	Description string      `gorm:"not null"`
	Origin      string      `gorm:"not null"`
	Contents    MultiString `gorm:"type:text not null"`
	Cost        int64       `gorm:"not null index"`
}
