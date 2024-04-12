package main

import (
	"fmt"

	"github.com/2k4sm/httpCoffee/db"
	"github.com/2k4sm/httpCoffee/routes"
	"github.com/2k4sm/httpCoffee/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func main() {
	utils.LoadEnv()

	var dbConf db.Config = db.Config{
		Host:     utils.GetEnv("HOST"),
		Username: utils.GetEnv("USERNAME"),
		Password: utils.GetEnv("PASSWORD"),
		DBName:   utils.GetEnv("DB"),
		Port:     utils.GetEnv("PORT"),
	}

	db.InitDB(&dbConf)

	app := fiber.New(fiber.Config{
		Prefork:      false,
		ServerHeader: "httpCoffee",
		AppName:      fmt.Sprintf("httpCoffee %s", utils.GetEnv("VERSION")),
	})

	api := app.Group("/api/v1")

	coffees := api.Group("/coffees")

	routes.CoffeeRoutes(coffees)

	log.Fatal(app.Listen(":6969"))
}
