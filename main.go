package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/2k4sm/httpCoffee/src/db"
	"github.com/2k4sm/httpCoffee/src/routes"
	"github.com/2k4sm/httpCoffee/src/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/redirect"
)

func main() {
	utils.LoadEnv()

	var dbConf db.Config = db.Config{
		Host:     utils.GetEnv("HOST"),
		Username: utils.GetEnv("USERNAME"),
		Password: utils.GetEnv("PASSWORD"),
		DBName:   utils.GetEnv("DB"),
		DBPORT:   utils.GetEnv("DBPORT"),
		SSLMODE:  utils.GetEnv("SSLMODE"),
	}

	DB := db.InitDB(&dbConf)

	app := fiber.New(fiber.Config{
		Prefork:      false,
		ServerHeader: "httpCoffee",
		AppName:      fmt.Sprintf("httpCoffee %s", utils.GetEnv("VERSION")),
	})

	app.Use(redirect.New(redirect.Config{
		Rules: map[string]string{
			"/": "/v0",
		},
		StatusCode: 301,
	}))

	app.Get("/"+utils.GetEnv("VERSION"), checkServer)

	api := app.Group("/" + utils.GetEnv("VERSION"))

	coffees := api.Group("/coffees")
	houses := api.Group("/houses")

	routes.CoffeeRoutes(coffees, DB)
	routes.HouseRoutes(houses, DB)
	log.Fatal(app.Listen(":6969"))
}

func checkServer(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(
		fiber.Map{
			"status":  http.StatusOK,
			"message": "server started successfully",
			"time":    time.Now(),
		},
	)
}
