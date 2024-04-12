package routes

import (
	"github.com/2k4sm/httpCoffee/handlers"
	"github.com/2k4sm/httpCoffee/repositories"
	"github.com/2k4sm/httpCoffee/services"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CoffeeRoutes(router fiber.Router, db *gorm.DB) {
	coffeeRepository := repositories.NewCoffeeRepository(db)
	coffeeService := services.NewCoffeeService(coffeeRepository)
	coffeeHandler := handlers.NewCoffeeHandler(coffeeService)

	router.Get("/", coffeeHandler.GetCoffees)
	router.Get("/:coffeeId", coffeeHandler.GetCoffeeById)
	router.Get("/names/:coffeeName", coffeeHandler.GetCoffeeByName)
	router.Post("/", coffeeHandler.CreateNewCoffee)
	router.Delete("/:coffeeId", coffeeHandler.DeleteCoffeeById)

}
