package routes

import (
	"github.com/Fasunle/golang-starters.git/cmd/handlers"
	fiber "github.com/gofiber/fiber/v3"
)

func appRoutes(app *fiber.App) {
	app.Get("/api/v1", handlers.HomeHandler)
}
