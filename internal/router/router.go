package router

import "github.com/gofiber/fiber/v3"

func SetupRoutes(app *fiber.App) {
	app.Get("/ping", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "pong",
		})
	})

	setupCustomerRoutes(app)
	setupCarRoutes(app)
}
