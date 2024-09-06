package routes

import (
	"github.com/gofiber/fiber/v3"
)

func New(app *fiber.App) {
	// configure all the middlewares here
	middlewares(app)
	// app routes
	appRoutes(app)
}
