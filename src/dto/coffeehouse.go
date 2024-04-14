package dto

import (
	"github.com/2k4sm/httpCoffee/src/entities"
)

type CoffeeHouse struct {
	Id               uint      `json:"id"`
	Name             string    `json:"house_name"`
	UserCount        int64     `json:"user_count"`
	TopCoffee        string    `json:"top_coffee"`
	Revenue          int64     `json:"revenue"`
	AvailableCoffees []Coffee  `json:"coffees"`
	Payments         []Payment `json:"payments"`
}

type CreateCoffeeHouse struct {
	Name             string            `json:"house_name"`
	AvailableCoffees []entities.Coffee `json:"coffees"`
}

func ParseToHouseEntity(newHouse CreateCoffeeHouse) entities.CoffeeHouse {
	house := entities.CoffeeHouse{
		Name:             newHouse.Name,
		AvailableCoffees: newHouse.AvailableCoffees,
	}

	return house
}

func ParseFromHouseEntity(house entities.CoffeeHouse) CoffeeHouse {
	var Coffees []Coffee

	for _, coffee := range house.AvailableCoffees {
		Coffees = append(Coffees, ParseFromCoffeeEntity(coffee))
	}

	coffees := CoffeeHouse{
		Id:               house.ID,
		Name:             house.Name,
		UserCount:        house.UserCount,
		TopCoffee:        house.TopCoffee,
		Revenue:          house.Revenue,
		AvailableCoffees: Coffees,
	}
	return coffees
}
