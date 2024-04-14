package handlers

import (
	"net/http"
	"strconv"

	"github.com/2k4sm/httpCoffee/src/dto"
	"github.com/2k4sm/httpCoffee/src/services"
	"github.com/gofiber/fiber/v2"
)

type HouseHandlerInterface interface {
	GetHouses(ctx *fiber.Ctx) error
	GetHouseById(ctx *fiber.Ctx) error
	GetHouseByName(ctx *fiber.Ctx) error
	CreateNewHouse(ctx *fiber.Ctx) error
	DeleteHouseById(ctx *fiber.Ctx) error
}

type houseHandler struct {
	HouseService services.HouseServiceInterface
}

func NewHouseHandler(houseService services.HouseServiceInterface) HouseHandlerInterface {
	return &houseHandler{
		HouseService: houseService,
	}
}

func (h *houseHandler) GetHouses(ctx *fiber.Ctx) error {
	houses := h.HouseService.GetHouses()

	if len(houses) == 0 {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  http.StatusNotFound,
			"message": "No houses found",
		})
	}

	houseResponse := []dto.CoffeeHouse{}
	for _, house := range houses {
		houseResponse = append(houseResponse, dto.ParseFromHouseEntity(house))
	}

	return ctx.JSON(houseResponse)
}

func (h *houseHandler) GetHouseById(ctx *fiber.Ctx) error {
	reqId := ctx.Params("houseId")

	reqIntId, err := strconv.Atoi(reqId)

	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
	}

	house := h.HouseService.GetHouseById(reqIntId)

	if house.ID == 0 {

		return ctx.Status(http.StatusNotFound).JSON(fiber.Map{
			"status":  http.StatusNotFound,
			"message": "House not found",
		})
	}

	return ctx.JSON(dto.ParseFromHouseEntity(house))

}

func (h *houseHandler) GetHouseByName(ctx *fiber.Ctx) error {
	reqName := ctx.Params("houseName")

	house := h.HouseService.GetHouseByName(reqName)

	if house.ID == 0 {

		return ctx.Status(http.StatusNotFound).JSON(fiber.Map{
			"status":  http.StatusNotFound,
			"message": "House not found",
		})
	}

	return ctx.JSON(dto.ParseFromHouseEntity(house))
}

func (h *houseHandler) CreateNewHouse(ctx *fiber.Ctx) error {
	var newHouse dto.CreateCoffeeHouse

	if err := ctx.BodyParser(&newHouse); err != nil {

		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
	}

	house := h.HouseService.CreateNewHouse(newHouse)

	return ctx.JSON(dto.ParseFromHouseEntity(house))
}

func (h *houseHandler) DeleteHouseById(ctx *fiber.Ctx) error {

	reqId := ctx.Params("houseId")

	reqIntId, err := strconv.Atoi(reqId)

	if err != nil {

		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
	}

	house, err := h.HouseService.DeleteHouseById(reqIntId)

	if err != nil {

		return ctx.Status(http.StatusNotFound).JSON(fiber.Map{
			"status":  http.StatusNotFound,
			"message": err.Error(),
		})
	}

	return ctx.JSON(dto.ParseFromHouseEntity(house))
}
