package routes

import (
	"encoding/json"
	"fmt"
	"umd/middleware"
	"umd/models"

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
			a := fiber.AcquireAgent()
			defer fiber.ReleaseAgent(a)
			req := a.Request()
			req.Header.SetMethod(fiber.MethodGet)
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", access))
			req.Header.SetRequestURI("https://api.box.com/2.0/users/me")
			if err := a.Parse(); err != nil {
				c.JSON(fiber.Map{"Error": err.Error()})
			}
			_, body, errs := a.Bytes()
			// log.Println(code, string(body))
			if errs != nil {
				return c.JSON(fiber.Map{"Error": errs})
			}
			user := &models.BoxUserResponse{}
			err := json.Unmarshal(body, user)
			if err != nil {
				return c.JSON(fiber.Map{"Error": err})
			}

			return c.Render("dashboard", fiber.Map{
				"Name":      user.Name,
				"AvatarURL": user.AvatarURL,
				"ID":        user.ID,
			})
		}
		return c.Render("login", fiber.Map{})
	})
}

func OauthRoutes(app *fiber.App) {
	auth := app.Group("/auth/box")
	auth.Get("/login", middleware.BoxAuthLogin)
	auth.Get("/logout", middleware.BoxAuthLogout)
	auth.Get("/callback", middleware.BoxOauthRedirect)
}
