package routes

import (
	"eventify/database"
	"eventify/models"
	"eventify/views"
	components "eventify/views/components/display"
	"strings"

	"strconv"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

func EventRouter(app *fiber.App) fiber.Router {
	router := app.Group("/event", func(c *fiber.Ctx) error {
		return c.Next()
	})

	router.Get("/:eventId", func(c *fiber.Ctx) error {
		if c.Params("eventId") != "" {
			eventId, err := strconv.ParseInt(c.Params("eventId"), 10, 64)
			if err != nil {
				return c.SendString("Invalid booking ID")
			}
			event := database.GetEvent(eventId)
			var view templ.Component
			if c.Context().Referer() == nil {
				view = views.GetOuterHtml[models.Event](event, "Event: "+c.Params("eventId"))
			} else if strings.Contains(string(c.Context().Referer()), string(c.Context().Path())) {
				view = views.GetOuterHtml[models.Event](event, "Event: "+c.Params("eventId"))
			} else {
				view = components.FormatEvent(event)
			}
			handler := adaptor.HTTPHandler(templ.Handler(view))

			return handler(c)
		}
		return c.SendString("Invalid booking ID")
	})

	router.Get("/edit/:eventId/", func(c *fiber.Ctx) error {
		if c.Params("eventId") != "" {
			eventId, err := strconv.ParseInt(c.Params("eventId"), 10, 64)
			if err != nil {
				return c.SendString("Invalid booking ID")
			}
			event := database.GetEvent(eventId)
			var view templ.Component
			if c.Context().Referer() == nil {
				view = views.EditOuterHtml[models.Event](event, "Event: "+c.Params("eventId"))
			} else if strings.Contains(string(c.Context().Referer()), string(c.Context().Path())) {
				view = views.EditOuterHtml[models.Event](event, "Event: "+c.Params("eventId"))
			} else {
				view = components.EditEvent(event)
			}
			handler := adaptor.HTTPHandler(templ.Handler(view))

			return handler(c)
		}
		return c.SendString("Invalid booking ID")
	})

	router.Post("/edit/:eventId/", func(c *fiber.Ctx) error {
		if c.Params("eventId") != "" {
			eventId, err := strconv.ParseInt(c.Params("eventId"), 10, 64)
			if err != nil {
				return c.SendString("Invalid event ID")
			}

			var event models.BaseEvent
			if err := c.BodyParser(&event); err != nil {
				return c.Status(fiber.StatusBadRequest).SendString("Failed to parse request body")
			}

			event.Id = eventId
			if err := database.SetEvent(event); err {
				return c.Status(fiber.StatusInternalServerError).SendString("Failed to update event")
			}

			return c.SendString("Event updated successfully")
		}
		return c.SendString("Invalid event ID")
	})

	return router
}
