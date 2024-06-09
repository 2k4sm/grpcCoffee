package routes

import (
	"github.com/2k4sm/httpCoffee/src/main/handlers"
	"github.com/2k4sm/httpCoffee/src/main/repositories"
	"github.com/2k4sm/httpCoffee/src/main/services"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func UserRoutes(router fiber.Router, db *gorm.DB) {

	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	userHandler := handlers.NewUserHandler(userService)

	router.Get("/", userHandler.GetAll)
	router.Get("/:userId", userHandler.GetById)
	router.Get("/name/:name", userHandler.GetByName)
	router.Get("/email/:email", userHandler.GetByEmail)
	router.Post("/", userHandler.CreateUser)
	router.Delete("/:userId", userHandler.DeleteUser)

}
