package entities

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Coffee struct {
	gorm.Model
	Name        string         `gorm:"not null unique index"`
	Description string         `gorm:"not null"`
	Origin      string         `gorm:"not null"`
	Contents    pq.StringArray `gorm:"type:text[]"`
	Cost        int64          `gorm:"not null index"`
}
