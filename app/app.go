package app

import "github.com/gofiber/fiber/v2"

var app *fiber.App

func init() {
	app = fiber.New()
}

func StartApp() {

	MapUrls()

	app.Listen(":8080")
}
