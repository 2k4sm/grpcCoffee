package main

import (
	"fmt"

	"github.com/2k4sm/httpCoffee/db"
	"github.com/2k4sm/httpCoffee/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func main() {
	utils.LoadEnv()
	app := fiber.New(fiber.Config{
		Prefork:      true,
		ServerHeader: "httpCoffee",
		AppName:      fmt.Sprintf("httpCoffee %s", utils.GetEnv("VERSION")),
	})

	var dbConf db.Config = db.Config{
		Host:     utils.GetEnv("HOST"),
		Username: utils.GetEnv("USERNAME"),
		Password: utils.GetEnv("PASSWORD"),
		DBName:   utils.GetEnv("DB"),
		Port:     utils.GetEnv("PORT"),
	}

	db.InitDB(&dbConf)

	log.Fatal(app.Listen(":6969"))
}
