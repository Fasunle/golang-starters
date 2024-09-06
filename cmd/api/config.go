package main

import (
	"time"

	"github.com/gofiber/fiber/v3"
)

var appConfig = fiber.Config{
	CaseSensitive: true,
	StrictRouting: true,
	ServerHeader:  "Fiber",
	AppName:       "Fiber",
	BodyLimit:     10 * 1024 * 1024,
	Immutable:     true,
	UnescapePath:  false,
	ReadTimeout:   5 * time.Second,
	WriteTimeout:  10 * time.Second,
}
