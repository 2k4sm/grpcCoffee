package dto

import (
	"github.com/2k4sm/httpCoffee/src/main/entities"
	"github.com/lib/pq"
)

type Coffee struct {
	Id          uint     `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Origin      string   `json:"origin"`
	Contents    []string `json:"contents"`
	Cost        int64    `json:"cost"`
}

type CreateCoffee struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Origin      string   `json:"origin"`
	Contents    []string `json:"contents"`
	Cost        int64    `json:"cost"`
}

func ParseFromCoffeeEntity(coffee entities.Coffee) Coffee {
	coffees := Coffee{
		Id:          coffee.ID,
		Name:        coffee.Name,
		Description: coffee.Description,
		Origin:      coffee.Origin,
		Contents:    []string(coffee.Contents),
		Cost:        coffee.Cost,
	}
	return coffees
}

func ParseToCoffeeEntity(coffee CreateCoffee) entities.Coffee {
	coffees := entities.Coffee{
		Name:        coffee.Name,
		Description: coffee.Description,
		Origin:      coffee.Origin,
		Contents:    pq.StringArray(coffee.Contents),
		Cost:        coffee.Cost,
	}
	return coffees
}
