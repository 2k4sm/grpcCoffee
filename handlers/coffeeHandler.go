package handlers

import (
	"net/http"
	"strconv"

	"github.com/2k4sm/httpCoffee/dto"
	"github.com/2k4sm/httpCoffee/services"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

type CoffeeHandlerInterface interface {
	GetCoffees(ctx *fiber.Ctx) error
	GetCoffeeById(ctx *fiber.Ctx) error
	GetCoffeeByName(ctx *fiber.Ctx) error
	CreateNewCoffee(ctx *fiber.Ctx) error
	DeleteCoffeeById(ctx *fiber.Ctx) error
}

type coffeeHandler struct {
	CoffeeService services.CoffeeServiceInterface
}

func NewCoffeeHandler(coffeeService services.CoffeeServiceInterface) CoffeeHandlerInterface {
	return &coffeeHandler{
		CoffeeService: coffeeService,
	}
}

func (c *coffeeHandler) GetCoffees(ctx *fiber.Ctx) error {
	coffees := c.CoffeeService.GetCoffees()

	coffeeResponse := []dto.Coffee{}

	for _, coffee := range coffees {
		coffeeResponse = append(coffeeResponse, dto.ParseFromEntities(coffee))
	}

	return ctx.JSON(coffeeResponse)
}

func (c *coffeeHandler) GetCoffeeById(ctx *fiber.Ctx) error {
	reqParam := ctx.Params("coffeeId")

	reqIntParam, err := strconv.Atoi(reqParam)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	coffee := c.CoffeeService.GetCoffeeById(reqIntParam)

	coffeeResponse := dto.ParseFromEntities(coffee)

	return ctx.JSON(coffeeResponse)
}

func (c *coffeeHandler) GetCoffeeByName(ctx *fiber.Ctx) error {
	reqParam := ctx.Params("coffeeName")

	coffee := c.CoffeeService.GetCoffeeByName(reqParam)

	coffeeResponse := dto.ParseFromEntities(coffee)

	return ctx.JSON(coffeeResponse)
}

func (c *coffeeHandler) CreateNewCoffee(ctx *fiber.Ctx) error {
	var newCoffee dto.CreateCoffee

	if err := ctx.BodyParser(&newCoffee); err != nil {
		log.Error(err)
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	coffee := c.CoffeeService.CreateNewCoffee(newCoffee)

	coffeeResponse := dto.ParseFromEntities(coffee)

	return ctx.JSON(coffeeResponse)

}

func (c *coffeeHandler) DeleteCoffeeById(ctx *fiber.Ctx) error {
	reqParam := ctx.Params("coffeeId")

	reqIntParam, err := strconv.Atoi(reqParam)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	err = c.CoffeeService.DeleteCoffee(reqIntParam)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	return ctx.JSON(http.StatusOK, "Coffee deleted successfully")
}
