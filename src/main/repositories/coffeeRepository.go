package repositories

import (
	"github.com/2k4sm/httpCoffee/src/main/entities"
	"github.com/gofiber/fiber/v2/log"
	"gorm.io/gorm"
)

type CoffeeRepositoryInterface interface {
	FindAll() []entities.Coffee
	FindById(id uint) entities.Coffee
	FindByName(coffeeName string) entities.Coffee
	Save(newCoffee *entities.Coffee) entities.Coffee
	Delete(coffee *entities.Coffee)
}

type coffeeRepository struct {
	Db *gorm.DB
}

func NewCoffeeRepository(db *gorm.DB) CoffeeRepositoryInterface {
	return &coffeeRepository{
		Db: db,
	}
}

func (c *coffeeRepository) FindAll() []entities.Coffee {
	var coffees []entities.Coffee
	c.Db.Find(&coffees)

	if len(coffees) == 0 {
		log.Info(gorm.ErrRecordNotFound.Error())
	}
	return coffees
}

func (c *coffeeRepository) FindById(id uint) entities.Coffee {
	var coffee entities.Coffee
	c.Db.First(&coffee, id)

	if coffee.ID == 0 {
		log.Info(gorm.ErrRecordNotFound)
	}
	return coffee
}

func (c *coffeeRepository) FindByName(coffeeName string) entities.Coffee {
	var coffee entities.Coffee
	c.Db.First(&coffee, "name = ?", coffeeName)

	if coffee.ID == 0 {
		log.Info(gorm.ErrRecordNotFound)
	}
	return coffee
}

func (c *coffeeRepository) Save(newCoffee *entities.Coffee) entities.Coffee {
	err := c.Db.Save(newCoffee)

	if err.Error != nil {
		log.Info(err.Error)
	}
	return *newCoffee
}

func (c *coffeeRepository) Delete(coffee *entities.Coffee) {
	err := c.Db.Delete(coffee)

	if err.Error != nil {
		log.Info(err.Error)
	}
}
