package handlers

import (
	"net/http"
	"strconv"

	"github.com/2k4sm/httpCoffee/src/main/dto"
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

func (p *payHandler) GetAll(ctx *fiber.Ctx) error {
	payments := p.PayService.GetAll()

	if (len(payments)) == 0 {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  http.StatusNotFound,
			"message": "No payments found",
		})
	}

	paymentsResponse := []dto.Payment{}

	for _, payment := range payments {
		paymentsResponse = append(paymentsResponse, dto.ParseFromPaymentEntity(payment))
	}

	return ctx.JSON(paymentsResponse)
}

func (p *payHandler) GetById(ctx *fiber.Ctx) error {
	reqId := ctx.Params("paymentId")

	reqIntId, err := strconv.Atoi(reqId)

	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
	}

	payment := p.PayService.GetById(reqIntId)

	if payment.ID == 0 {
		return ctx.Status(http.StatusNotFound).JSON(fiber.Map{
			"status":  http.StatusNotFound,
			"message": "Payment not found",
		})
	}

	return ctx.JSON(dto.ParseFromPaymentEntity(payment))
}

func (p *payHandler) GetAllByUserId(ctx *fiber.Ctx) error {
	reqId := ctx.Params("userId")

	reqIntId, err := strconv.Atoi(reqId)

	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
	}

	payment := p.PayService.GetById(reqIntId)

	if payment.ID == 0 {
		return ctx.Status(http.StatusNotFound).JSON(fiber.Map{
			"status":  http.StatusNotFound,
			"message": "Payment not found",
		})
	}

	return ctx.JSON(dto.ParseFromPaymentEntity(payment))
}

func (p *payHandler) GetAllByHouseId(ctx *fiber.Ctx) error {
	reqId := ctx.Params("houseId")

	reqIntId, err := strconv.Atoi(reqId)

	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
	}

	payment := p.PayService.GetById(reqIntId)

	if payment.ID == 0 {
		return ctx.Status(http.StatusNotFound).JSON(fiber.Map{
			"status":  http.StatusNotFound,
			"message": "Payment not found",
		})
	}

	return ctx.JSON(dto.ParseFromPaymentEntity(payment))
}

func (p *payHandler) CreatePayment(ctx *fiber.Ctx) error {

	var newPayment dto.CreatePayment

	if err := ctx.BodyParser(&newPayment); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
	}

	payment := p.PayService.CreatePayment(newPayment)

	return ctx.JSON(dto.ParseFromPaymentEntity(payment))
}

func (p *payHandler) DeletePayment(ctx *fiber.Ctx) error {

	reqId := ctx.Params("paymentId")

	reqIntId, err := strconv.Atoi(reqId)

	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
	}

	payment, err := p.PayService.DeletePayment(reqIntId)

	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
	}

	return ctx.JSON(dto.ParseFromPaymentEntity(payment))
}
