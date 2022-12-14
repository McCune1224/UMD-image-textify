package middleware

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/oauth2"
)

func boxOauthConfig() *oauth2.Config {
	conf := &oauth2.Config{
		ClientID:     os.Getenv("BOX_CLIENT_ID"),
		ClientSecret: os.Getenv("BOX_CLIENT_SECRET"),
		RedirectURL:  os.Getenv("REDIRECT_URI"),
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://account.box.com/api/oauth2/authorize",
			TokenURL: "https://api.box.com/oauth2/token",
		},
	}
	return conf
}

func BoxAuthLogin(c *fiber.Ctx) error {
	path := boxOauthConfig()
	url := path.AuthCodeURL("state")
	return c.Redirect(url)
}

func BoxAuthLogout(c *fiber.Ctx) error {
	c.Cookie(&fiber.Cookie{
		Name:    "RefreshToken",
		Value:   "",
		Expires: time.Unix(0, 0),
	})
	c.Cookie(&fiber.Cookie{
		Name:    "AccessToken",
		Value:   "",
		Expires: time.Unix(0, 0),
	})
	return c.Redirect("/")
}

func BoxOauthRedirect(c *fiber.Ctx) error {
	code := struct {
		Code string `json:"code"`
	}{}

	if err := c.BodyParser(&code); err != nil {
		return (c.JSON(fiber.Map{"Error": err.Error()}))
	}

	payload, err := boxOauthConfig().Exchange(c.Context(), code.Code)
	if err != nil {
		return c.JSON(fiber.Map{"Error": err.Error()})
	}
	if payload == nil {
		return c.JSON(fiber.Map{"Error": "token nil"})
	}

	return c.JSON(payload)
}

// func BoxOauthRedirect(c *fiber.Ctx) error {
// 	// Get AccessToken from endpoints
// 	payload, err := boxOauthConfig().Exchange(c.Context(), c.FormValue("code"))
// 	if err != nil {
// 		return c.JSON(fiber.Map{"Error": err.Error()})
// 	}
// 	if payload == nil {
// 		return c.JSON(fiber.Map{"Error": "token nil"})
// 	}
//
// 	// Set AccessToken and RefreshToken in Browser Cookies
// 	c.Cookie(&fiber.Cookie{
// 		Name:    "AccessToken",
// 		Value:   payload.AccessToken,
// 		Expires: payload.Expiry,
// 	})
//
// 	c.Cookie(&fiber.Cookie{
// 		Name:  "RefreshToken",
// 		Value: payload.RefreshToken,
// 	})
//
// 	// Send user back to home page
// 	return c.Redirect("/")
// }
