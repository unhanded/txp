package main

import (
	"os"
	"time"

	"github.com/charmbracelet/log"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/unhanded/txp/internal/environ"
	"github.com/unhanded/txp/internal/txpfiber"
)

func main() {
	if environ.TxpDebug() {
		log.SetLevel(log.DebugLevel)
	}
	app := fiber.New(fiber.Config{DisableStartupMessage: true})
	app.Use(limiter.New(limiter.Config{Max: 6, Expiration: time.Second * 30}))

	templ := app.Group("/template")

	templ.Use(compress.New(compress.ConfigDefault))

	templ.All("/:templateName",
		func(c *fiber.Ctx) error {
			mth := c.Method()
			if mth == "POST" || mth == "GET" {
				return txpfiber.HandleCompile(c)
			}
			return c.Next()
		},
	)

	addr := ":" + os.Getenv("PORT")
	app.Listen(addr)
}
