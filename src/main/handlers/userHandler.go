package handlers

import (
	"net/http"
	"strconv"

	"github.com/2k4sm/httpCoffee/src/main/dto"
	"github.com/2k4sm/httpCoffee/src/main/services"
	"github.com/gofiber/fiber/v2"
)

type UserHandlerInterface interface {
	GetAll(ctx *fiber.Ctx) error
	GetById(ctx *fiber.Ctx) error
	GetByName(ctx *fiber.Ctx) error
	GetByEmail(ctx *fiber.Ctx) error
	CreateUser(ctx *fiber.Ctx) error
	DeleteUser(ctx *fiber.Ctx) error
}

type userHandler struct {
	UserService services.UserServiceInterface
}

func NewUserHandler(userService services.UserServiceInterface) UserHandlerInterface {
	return &userHandler{
		UserService: userService,
	}
}

func (u *userHandler) GetAll(ctx *fiber.Ctx) error {
	users := u.UserService.GetAll()

	if len(users) == 0 {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  http.StatusNotFound,
			"message": "No users found",
		})
	}

	usersResponse := []dto.User{}

	for _, user := range users {
		usersResponse = append(usersResponse, dto.ParseFromUserEntity(user))
	}

	return ctx.JSON(usersResponse)
}

func (u *userHandler) GetById(ctx *fiber.Ctx) error {
	reqId := ctx.Params("userId")

	reqIdInt, err := strconv.Atoi(reqId)

	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
	}

	user := u.UserService.GetById(reqIdInt)

	if user.ID == 0 {
		return ctx.Status(http.StatusNotFound).JSON(fiber.Map{
			"status":  http.StatusNotFound,
			"message": "No user found",
		})
	}

	return ctx.JSON(dto.ParseFromUserEntity(user))
}

func (u *userHandler) GetByName(ctx *fiber.Ctx) error {
	reqName := ctx.Params("name")

	user := u.UserService.GetByName(reqName)

	if user.ID == 0 {
		return ctx.Status(http.StatusNotFound).JSON(fiber.Map{
			"status":  http.StatusNotFound,
			"message": "No user found",
		})
	}

	return ctx.JSON(dto.ParseFromUserEntity(user))
}

func (u *userHandler) GetByEmail(ctx *fiber.Ctx) error {
	reqEmail := ctx.Params("email")

	user := u.UserService.GetByEmail(reqEmail)

	if user.ID == 0 {
		return ctx.Status(http.StatusNotFound).JSON(fiber.Map{
			"status":  http.StatusNotFound,
			"message": "No user found",
		})
	}

	return ctx.JSON(dto.ParseFromUserEntity(user))
}

func (u *userHandler) CreateUser(ctx *fiber.Ctx) error {
	var createUser dto.CreateUser

	if err := ctx.BodyParser(&createUser); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
	}

	newUser := u.UserService.CreateUser(createUser)

	return ctx.JSON(dto.ParseFromUserEntity(newUser))
}

func (u *userHandler) DeleteUser(ctx *fiber.Ctx) error {
	reqId := ctx.Params("userId")

	reqIdInt, err := strconv.Atoi(reqId)

	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
	}

	user, err := u.UserService.DeleteUser(reqIdInt)

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	return ctx.JSON(user)
}
