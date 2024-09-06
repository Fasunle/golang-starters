package routes

import (
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cache"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/csrf"
	"github.com/gofiber/fiber/v3/middleware/helmet"
	"github.com/gofiber/fiber/v3/middleware/idempotency"
	"github.com/gofiber/fiber/v3/middleware/limiter"
)

var ConfigDefault = helmet.Config{
	XSSProtection:             "0",
	ContentTypeNosniff:        "nosniff",
	XFrameOptions:             "SAMEORIGIN",
	ReferrerPolicy:            "no-referrer",
	CrossOriginEmbedderPolicy: "require-corp",
	CrossOriginOpenerPolicy:   "same-origin",
	CrossOriginResourcePolicy: "same-origin",
	OriginAgentCluster:        "?1",
	XDNSPrefetchControl:       "off",
	XDownloadOptions:          "noopen",
	XPermittedCrossDomain:     "none",
}

var corsOptions = cors.Config{
	AllowOrigins:     []string{"http://localhost:80"},
	AllowCredentials: true,
	AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	AllowHeaders:     []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
	ExposeHeaders:    []string{"Link"},
	MaxAge:           300, // Maximum value not ignored by any of major browsers
}

var rateLimitOptions = limiter.Config{
	Max:        5,
	Expiration: 1 * time.Minute,
	KeyGenerator: func(c fiber.Ctx) string {
		return c.IP()
	},
	LimitReached: func(c fiber.Ctx) error {
		return c.SendStatus(fiber.StatusTooManyRequests)
	},
	SkipFailedRequests:     false,
	SkipSuccessfulRequests: false,
	LimiterMiddleware:      limiter.FixedWindow{},
}

var cacheOptions = cache.Config{
	Next: func(c fiber.Ctx) bool {
		return c.Query("noCache") == "true"
	},
	Expiration:   30 * time.Minute,
	CacheControl: true,
}

var csrfOptions = csrf.Config{
	KeyLookup:      "header:X-Csrf-Token",
	CookieName:     "csrf_",
	CookieSameSite: "Lax",
	Expiration:     1 * time.Hour,
}

func middlewares(app *fiber.App) {
	// initialize middlewares
	app.Use(cors.New(corsOptions))
	app.Use(cache.New(cacheOptions))
	app.Use(csrf.New(csrfOptions))
	app.Use(limiter.New(rateLimitOptions))
	app.Use(idempotency.New()) // prevent server from making duplicate requests
}
