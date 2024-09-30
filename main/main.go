package main

import (
	"eventify/database"
	"eventify/models"
	"eventify/views"

	"slices"
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

	app.Static("/", "/static")

	app.Get("/", func(c *fiber.Ctx) error {
		view := views.GetOuterHtml[models.Booking](database.GetBooking(2), "Home")
		handler := adaptor.HTTPHandler(templ.Handler(view))
		return handler(c)
	})

	app.Get("/booking/:bookingId", func(c *fiber.Ctx) error {
		if c.Params("bookingId") != "" {
			bookingId, err := strconv.ParseInt(c.Params("bookingId"), 10, 64)
			if err != nil {
				return c.SendString("Invalid booking ID")
			}

			booking := database.GetBooking(bookingId)
			var view templ.Component
			if c.Context().Referer() != nil {
				view = views.GetOuterHtml[models.Booking](booking, "Event: "+c.Params("bookingId"))
			} else if slices.Equal(c.Context().Referer(), []byte("/")) {

			} else {
				view = views.FormatBooking(booking)
			}
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
			event := database.GetEvent(eventId)
			var view templ.Component
			if c.Context().Referer() != nil {
				view = views.FormatEvent(event)
			} else {
				view = views.GetOuterHtml[models.Event](event, "Event: "+c.Params("eventId"))
			}
			handler := adaptor.HTTPHandler(templ.Handler(view))

			return handler(c)
		}
		return c.SendString("Invalid booking ID")
	})

	app.Delete("/booking/:bookingId", func(c *fiber.Ctx) error {
		if c.Params("bookingId") != "" {
			bookingId, err := strconv.ParseInt(c.Params("bookingId"), 10, 64)
			if err != nil {
				return c.SendString("Invalid booking ID")
			}

			database.DeleteBooking(bookingId)

			return c.SendString("Booking deleted")
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
