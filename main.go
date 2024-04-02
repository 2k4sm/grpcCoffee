package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		Prefork:      true,
		ServerHeader: "httpCoffee",
		AppName:      "httpCoffee V0.0.1",
	})

	log.Fatal(app.Listen(":6969"))
}
