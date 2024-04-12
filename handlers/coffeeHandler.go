package handlers

import (
	"net/http"
	"strconv"

	"github.com/2k4sm/httpCoffee/dto"
	"github.com/2k4sm/httpCoffee/services"
	"github.com/gofiber/fiber/v2"
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

	if len(coffees) == 0 {
		return ctx.Status(http.StatusNotFound).JSON(fiber.Map{
			"status":  http.StatusNotFound,
			"message": "No coffees found",
		})
	}

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
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
	}

	coffee := c.CoffeeService.GetCoffeeById(reqIntParam)

	if coffee.ID == 0 {
		return ctx.Status(http.StatusNotFound).JSON(fiber.Map{
			"status":  http.StatusNotFound,
			"message": "coffee not found.",
		})
	}

	coffeeResponse := dto.ParseFromEntities(coffee)

	return ctx.JSON(coffeeResponse)
}

func (c *coffeeHandler) GetCoffeeByName(ctx *fiber.Ctx) error {
	reqParam := ctx.Params("coffeeName")

	coffee := c.CoffeeService.GetCoffeeByName(reqParam)

	if coffee.ID == 0 {
		return ctx.Status(http.StatusNotFound).JSON(fiber.Map{
			"status":  http.StatusNotFound,
			"message": "coffee not found.",
		})
	}

	coffeeResponse := dto.ParseFromEntities(coffee)

	return ctx.JSON(coffeeResponse)
}

func (c *coffeeHandler) CreateNewCoffee(ctx *fiber.Ctx) error {
	var newCoffee dto.CreateCoffee

	if err := ctx.BodyParser(&newCoffee); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
	}

	coffee := c.CoffeeService.CreateNewCoffee(newCoffee)

	coffeeResponse := dto.ParseFromEntities(coffee)

	return ctx.Status(http.StatusCreated).JSON(coffeeResponse)

}

func (c *coffeeHandler) DeleteCoffeeById(ctx *fiber.Ctx) error {
	reqParam := ctx.Params("coffeeId")

	reqIntParam, err := strconv.Atoi(reqParam)

	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
	}

	err = c.CoffeeService.DeleteCoffee(reqIntParam)

	if err != nil {
		return ctx.Status(http.StatusNotFound).JSON(fiber.Map{
			"status":  http.StatusNotFound,
			"message": err.Error(),
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": "coffee deleted successfully",
	})
}
