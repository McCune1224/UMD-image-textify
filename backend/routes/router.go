package routes

import (
	"umd/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func RegisterDefaults(app *fiber.App) {
	// Enable Logging
	app.Use(logger.New())
	app.Use(recover.New())

	// CORS for api development
	app.Use(cors.New(cors.ConfigDefault))
	// Initialize standard Go html template engine

	app.Get("/", func(c *fiber.Ctx) error {
		access := c.Cookies("AccessToken")
		if access != "" {
			return c.Render("dashboard", fiber.Map{"Token": access})
		}
		return c.Render("login", fiber.Map{})
	})
}

func OauthRoutes(app *fiber.App) {
	auth := app.Group("/auth")
	auth.Get("/box", middleware.BoxAuth)
	auth.Get("/box/callback", middleware.BoxOauthRedirect)
}
