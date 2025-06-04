package main

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/unhanded/txp/internal/txpfiber"
)

func main() {
	app := fiber.New(fiber.Config{DisableStartupMessage: true})
	app.Use(limiter.New(limiter.Config{Max: 6, Expiration: time.Second * 30}))

	app.Post("/v1/templates/:templateName", txpfiber.HandleCompile)

	addr := ":" + os.Getenv("PORT")
	app.Listen(addr)
}
