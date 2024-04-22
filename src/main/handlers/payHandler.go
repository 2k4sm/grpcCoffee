package handlers

import (
	"github.com/2k4sm/httpCoffee/src/main/services"
	"github.com/gofiber/fiber/v2"
)

type PayHandlerInterface interface {
	GetAll(ctx *fiber.Ctx) error
	GetById(ctx *fiber.Ctx) error
	GetAllByUserId(ctx *fiber.Ctx) error
	GetAllByHouseId(ctx *fiber.Ctx) error
	CreatePayment(ctx *fiber.Ctx) error
	DeletePayment(ctx *fiber.Ctx) error
}

type payHandler struct {
	PayService services.PayServiceInterface
}

func NewPayHandler(payService services.PayServiceInterface) PayHandlerInterface {
	return &payHandler{
		PayService: payService,
	}
}




