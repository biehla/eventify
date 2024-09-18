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

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Redirect("/booking/2")
	})

	app.Get("/booking/:bookingId", func(c *fiber.Ctx) error {
		if c.Params("bookingId") != "" {
			bookingId, err := strconv.ParseInt(c.Params("bookingId"), 10, 64)
			if err != nil {
				return c.SendString("Invalid booking ID")
			}

			booking := database.GetBooking(bookingId)
			view := views.GetBooking(booking)
			handler := adaptor.HTTPHandler(templ.Handler(view))

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
			event := database.GetEvent(eventId)
			view := views.GetEvent(event)
			handler := adaptor.HTTPHandler(templ.Handler(view))

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
