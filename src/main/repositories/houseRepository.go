package repositories

import (
	"github.com/2k4sm/httpCoffee/src/main/entities"
	"github.com/gofiber/fiber/v2/log"
	"gorm.io/gorm"
)

type HouseRepositoryInterface interface {
	FindAll() []entities.CoffeeHouse
	FindById(id uint) entities.CoffeeHouse
	FindByName(houseName string) entities.CoffeeHouse
	Save(newHouse *entities.CoffeeHouse) entities.CoffeeHouse
	DeleteById(id uint) (entities.CoffeeHouse, error)
}

type houseRepository struct {
	Db *gorm.DB
}

func NewHouseRepository(db *gorm.DB) HouseRepositoryInterface {
	return &houseRepository{
		Db: db,
	}
}

func (h *houseRepository) FindAll() []entities.CoffeeHouse {
	var houses []entities.CoffeeHouse
	h.Db.Preload("AvailableCoffees").Preload("Payments").Find(&houses)

	if len(houses) == 0 {
		log.Info("No houses found")
	}

	return houses
}

func (h *houseRepository) FindById(id uint) entities.CoffeeHouse {
	var house entities.CoffeeHouse

	h.Db.Preload("AvailableCoffees").Preload("Payments").First(&house, id)

	if house.ID == 0 {
		log.Info(gorm.ErrRecordNotFound)
	}

	return house
}

func (h *houseRepository) FindByName(houseName string) entities.CoffeeHouse {
	var house entities.CoffeeHouse

	h.Db.Preload("AvailableCoffees").Preload("Payments").First(&house, "name = ?", houseName)

	if house.ID == 0 {
		log.Info(gorm.ErrRecordNotFound)
	}
	return house
}

func (h *houseRepository) Save(newHouse *entities.CoffeeHouse) entities.CoffeeHouse {
	err := h.Db.Save(newHouse)
	if err.Error != nil {
		log.Info(err.Error)
	}
	var createdHouse entities.CoffeeHouse
	h.Db.Preload("AvailableCoffees").Preload("Payments").First(&createdHouse, "name = ?", newHouse.Name)

	return createdHouse
}

func (h *houseRepository) DeleteById(id uint) (entities.CoffeeHouse, error) {

	var houseToDel entities.CoffeeHouse

	h.Db.Preload("AvailableCoffees").Preload("Payments").First(&houseToDel, id)

	if houseToDel.ID == 0 {
		return houseToDel, gorm.ErrRecordNotFound
	}

	h.Db.Delete(&houseToDel)

	return houseToDel, nil
}
