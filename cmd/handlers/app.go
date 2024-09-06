package handlers

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v3"
)

func HomeHandler(c fiber.Ctx) error {
	return c.SendString(fmt.Sprintf("Hello world from the golang starter %d", os.Getpid()))
}
