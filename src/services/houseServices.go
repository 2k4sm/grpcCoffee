package services

import (
	"github.com/2k4sm/httpCoffee/src/dto"
	"github.com/2k4sm/httpCoffee/src/entities"
	"github.com/2k4sm/httpCoffee/src/repositories"
)

type HouseServiceInterface interface {
	GetHouses() []entities.CoffeeHouse
	GetHouseById(id int) entities.CoffeeHouse
	GetHouseByName(houseName string) entities.CoffeeHouse
	CreateNewHouse(newHouse dto.CreateCoffeeHouse) entities.CoffeeHouse
	DeleteHouseById(id int) (entities.CoffeeHouse, error)
}

type houseService struct {
	HouseRepository repositories.HouseRepositoryInterface
}

func NewHouseService(houseRepository repositories.HouseRepositoryInterface) HouseServiceInterface {
	return &houseService{
		HouseRepository: houseRepository,
	}
}

func (h *houseService) GetHouses() []entities.CoffeeHouse {
	return h.HouseRepository.FindAll()
}

func (h *houseService) GetHouseById(id int) entities.CoffeeHouse {
	return h.HouseRepository.FindById(uint(id))
}

func (h *houseService) GetHouseByName(houseName string) entities.CoffeeHouse {
	return h.HouseRepository.FindByName(houseName)
}

func (h *houseService) CreateNewHouse(newHouse dto.CreateCoffeeHouse) entities.CoffeeHouse {
	house := dto.ParseToHouseEntity(newHouse)

	return h.HouseRepository.Save(&house)
}

func (h *houseService) DeleteHouseById(id int) (entities.CoffeeHouse, error) {

	return h.HouseRepository.DeleteById(uint(id))

}
