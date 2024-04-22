package dto

import (
	"time"

	"github.com/2k4sm/httpCoffee/src/main/entities"
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
	UserID        uint              `json:"user_id,omitempty"`
	CoffeeHouseID uint              `json:"coffee_house_id,omitempty"`
	Cost          int64             `json:"cost,omitempty"`
	Date          time.Time         `json:"date,omitempty"`
	Items         []entities.Coffee `json:"items,omitempty"`
}

func ParseFromPaymentEntity(payment entities.Payment) Payment {

	items := []Coffee{}
	for _, item := range payment.Items {
		items = append(items, ParseFromCoffeeEntity(item))
	}
	payments := Payment{
		Id:            payment.ID,
		UserID:        payment.UserID,
		CoffeeHouseID: payment.CoffeeHouseID,
		Cost:          payment.Cost,
		Date:          payment.Date,
		Items:         items,
	}
	return payments
}

func ParseToPaymentEntity(payment CreatePayment) entities.Payment {

	payments := entities.Payment{
		UserID:        payment.UserID,
		CoffeeHouseID: payment.CoffeeHouseID,
		Cost:          payment.Cost,
		Date:          payment.Date,
		Items:         payment.Items,
	}
	return payments
}
