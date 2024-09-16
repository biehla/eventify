package main

import (
	"eventify/database"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/:name?", func(c *fiber.Ctx) error {
		if c.Params("name") != "" {
			return c.SendString("Name: " + c.Params("name"))
		}
		return c.SendString("Who's John?")
	})

	database.SetupDB()

	app.Listen(":3000")
}
