package dto

import (
	"github.com/2k4sm/httpCoffee/src/main/entities"
)

type User struct {
	Id            uint          `json:"id"`
	Name          string        `json:"user_name"`
	Email         string        `json:"email"`
	Password      string        `json:"password"`
	LastOrder     string        `json:"last_order"`
	Favourite     string        `json:"favourite"`
	Revenue       string        `json:"-"`
	Orders        []Payment     `json:"orders"`
	VisitedHouses []CoffeeHouse `json:"visited_houses"`
}

type CreateUser struct {
	Name     string `json:"user_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func ParseToUserEntity(createUser CreateUser) entities.User {

	user := entities.User{
		Name:     createUser.Name,
		Email:    createUser.Email,
		Password: createUser.Password,
	}

	return user
}

func ParseFromUserEntity(user entities.User) User {

	payments := []Payment{}
	for _, payment := range user.Orders {
		payments = append(payments, ParseFromPaymentEntity(payment))
	}

	houses := []CoffeeHouse{}
	for _, house := range user.VisitedHouses {
		houses = append(houses, ParseFromHouseEntity(house))
	}

	return User{
		Id:            user.ID,
		Name:          user.Name,
		Email:         user.Email,
		Password:      user.Password,
		LastOrder:     user.LastOrder,
		Favourite:     user.Favourite,
		Revenue:       user.Revenue,
		Orders:        payments,
		VisitedHouses: houses,
	}
}
