package repositories

import (
	"github.com/2k4sm/httpCoffee/entities"
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
	return coffees
}

func (c *coffeeRepository) FindById(id uint) entities.Coffee {
	var coffee entities.Coffee
	c.Db.First(&coffee, id)
	return coffee
}

func (c *coffeeRepository) FindByName(coffeeName string) entities.Coffee {
	var coffee entities.Coffee
	c.Db.First(&coffee, "name = ?", coffeeName)
	return coffee
}

func (c *coffeeRepository) Save(newCoffee *entities.Coffee) entities.Coffee {
	c.Db.Save(newCoffee)
	return *newCoffee
}

func (c *coffeeRepository) Delete(coffee *entities.Coffee) {
	c.Db.Delete(coffee)
}
