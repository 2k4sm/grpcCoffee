package services

import (
	"github.com/2k4sm/httpCoffee/dto"
	"github.com/2k4sm/httpCoffee/entities"
	"github.com/2k4sm/httpCoffee/repositories"
	"github.com/gofiber/fiber/v2/log"
	"gorm.io/gorm"
)

type CoffeeServiceInterface interface {
	GetCoffees() []entities.Coffee
	GetCoffeeById(id int) entities.Coffee
	GetCoffeeByName(coffeeName string) entities.Coffee
	CreateNewCoffee(newCoffee dto.CreateCoffee) entities.Coffee
	DeleteCoffee(id int) error
}

type coffeeService struct {
	CoffeeRepository repositories.CoffeeRepositoryInterface
}

func NewCoffeeService(coffeeRepository repositories.CoffeeRepositoryInterface) CoffeeServiceInterface {
	return &coffeeService{
		CoffeeRepository: coffeeRepository,
	}
}

func (c coffeeService) GetCoffees() []entities.Coffee {
	return c.CoffeeRepository.FindAll()
}

func (c coffeeService) GetCoffeeById(id int) entities.Coffee {
	return c.CoffeeRepository.FindById(uint(id))
}

func (c coffeeService) GetCoffeeByName(coffeeName string) entities.Coffee {
	return c.CoffeeRepository.FindByName(coffeeName)
}

func (c coffeeService) CreateNewCoffee(newCoffee dto.CreateCoffee) entities.Coffee {
	coffee := entities.Coffee{
		Name:        newCoffee.Name,
		Description: newCoffee.Description,
		Origin:      newCoffee.Origin,
		Contents:    entities.MultiString(newCoffee.Contents),
		Cost:        newCoffee.Cost,
	}

	return c.CoffeeRepository.Save(coffee)
}

func (c coffeeService) DeleteCoffee(id int) error {
	coffee := c.CoffeeRepository.FindById(uint(id))

	if coffee.ID == 0 {
		log.Error(gorm.ErrRecordNotFound)
		return gorm.ErrRecordNotFound
	}

	c.CoffeeRepository.Delete(coffee)

	return nil
}
