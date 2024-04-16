package dto

import (
	"time"
)

type Payment struct {
	Id            uint      `json:"id"`
	UserID        uint      `json:"user_id,omitempty"`
	CoffeeHouseID uint      `json:"coffee_house_id,omitempty"`
	Cost          int64     `json:"cost,omitempty"`
	Date          time.Time `json:"date,omitempty"`
	Items         []Coffee  `json:"items,omitempty"`
}

type CreatePayment struct {
	UserID        uint      `json:"user_id,omitempty"`
	CoffeeHouseID uint      `json:"coffee_house_id,omitempty"`
	Cost          int64     `json:"cost,omitempty"`
	Date          time.Time `json:"date,omitempty"`
	Items         []string  `json:"items,omitempty"`
}
