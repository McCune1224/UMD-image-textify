package main

import (
	"os"

	router "umd/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	_ "github.com/joho/godotenv/autoload"
)

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":3000"
	} else {
		port = ":" + port
	}

	return port
}

func main() {
	// Initialize standard Go html template engine
	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	router.RegisterDefaults(app)
	router.OauthRoutes(app)

	app.Listen(getPort())
}
