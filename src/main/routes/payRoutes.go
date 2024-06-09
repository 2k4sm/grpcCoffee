package routes

import (
	"github.com/2k4sm/httpCoffee/src/main/handlers"
	"github.com/2k4sm/httpCoffee/src/main/repositories"
	"github.com/2k4sm/httpCoffee/src/main/services"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func PaymentRoutes(router fiber.Router, db *gorm.DB) {
	paymentRepository := repositories.NewPayRepository(db)
	paymentService := services.NewPayService(paymentRepository)
	paymentHandler := handlers.NewPayHandler(paymentService)

	router.Get("/", paymentHandler.GetAll)
	router.Get("/:paymentId", paymentHandler.GetById)
	router.Get("/names/:houseId", paymentHandler.GetAllByHouseId)
	router.Get("/names/:userId", paymentHandler.GetAllByUserId)
	router.Post("/", paymentHandler.CreatePayment)
	router.Delete("/:paymentId", paymentHandler.DeletePayment)

}
