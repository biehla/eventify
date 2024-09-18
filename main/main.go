package main

import (
	"eventify/database"
	"eventify/models"
	"fmt"

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

	var booking models.Booking
	fmt.Println(database.GetBooking(17).ToString())
	fmt.Println()
	booking = database.GetBooking(17)
	eventID := booking.GetEventIds()
	event := database.GetEvent(eventID[0])
	fmt.Println(event.ToString())
	err := app.Listen(":3000")
	if err != nil {
		return
	}
}
