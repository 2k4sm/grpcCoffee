package routes

import (
	"github.com/2k4sm/httpCoffee/src/handlers"
	"github.com/2k4sm/httpCoffee/src/repositories"
	"github.com/2k4sm/httpCoffee/src/services"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func HouseRoutes(router fiber.Router, db *gorm.DB) {
	houseRepository := repositories.NewHouseRepository(db)
	houseService := services.NewHouseService(houseRepository)
	houseHandler := handlers.NewHouseHandler(houseService)

	router.Get("/", houseHandler.GetHouses)
	router.Get("/:houseId", houseHandler.GetHouseById)
	router.Get("/names/:houseName", houseHandler.GetHouseByName)
	router.Post("/", houseHandler.CreateNewHouse)
	router.Delete("/:houseId", houseHandler.DeleteHouseById)

}
