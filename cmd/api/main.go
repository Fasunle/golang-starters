package main

import (
	"fmt"
	"os"
	"time"

	"github.com/Fasunle/golang-starters.git/cmd/routes"
	"github.com/gofiber/fiber/v3"
)

type Config struct{}

func main() {
	app := Config{}
	//
	app.serve()
}

// start and serve http server
func (app *Config) serve() {
	addr := fmt.Sprintf(":%s", getPort())

	server := fiber.New(appConfig)

	// register all routes
	routes.New(server)

	server.Listen(addr, fiber.ListenConfig{
		EnablePrefork: true,
	})

	defer server.ShutdownWithTimeout(10 * time.Second)
}

// use the set port in env or default to 80
func getPort() string {
	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "80"
	}

	return PORT
}
