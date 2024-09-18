package main

import (
	"eventify/database"
	"eventify/views"

	"strconv"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

func main() {
	app := fiber.New(fiber.Config{
		ViewsLayout:       "views",
		PassLocalsToViews: true,
	})

	app.Get("/booking/:bookingId", func(c *fiber.Ctx) error {
		if c.Params("bookingId") != "" {
			bookingId, err := strconv.ParseInt(c.Params("bookingId"), 10, 64)
			if err != nil {
				return c.SendString("Invalid booking ID")
			}
			// return views.GetBooking(bookingId).Render(c, fiber.AcquireResponse().BodyWriter())
			booking := views.GetBooking(bookingId)
			handler := adaptor.HTTPHandler(templ.Handler(booking))
			return handler(c)
		}
		return c.SendString("Invalid booking ID")
	})

	app.Get("/event/:eventId", func(c *fiber.Ctx) error {
		if c.Params("eventId") != "" {
			eventId, err := strconv.ParseInt(c.Params("eventId"), 10, 64)
			if err != nil {
				return c.SendString("Invalid booking ID")
			}
			// return views.GetBooking(eventId).Render(c, fiber.AcquireResponse().BodyWriter())
			event := views.GetEvent(eventId)
			handler := adaptor.HTTPHandler(templ.Handler(event))
			return handler(c)
		}
		return c.SendString("Invalid booking ID")
	})

	database.SetupDB()

	// var booking models.Booking
	// fmt.Println(database.GetBooking(17).ToString())
	// fmt.Println()
	// booking = database.GetBooking(17)
	// eventID := booking.GetEventIds()
	// event := database.GetEvent(eventID[0])
	// fmt.Println(event.ToString())

	err := app.Listen(":3000")
	if err != nil {
		return
	}
}
