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
	"github.com/unhanded/txp/internal/txppack"
)

func main() {
	if environ.TxpDebug() {
		log.SetLevel(log.DebugLevel)
	}

	txppack.StartupCheck()

	app := fiber.New(fiber.Config{DisableStartupMessage: true})
	app.Use(limiter.New(limiter.Config{Max: 6, Expiration: time.Second * 30}))

	gTemplate := app.Group("/template")

	gTemplate.Use(compress.New(compress.ConfigDefault))

	gTemplate.Get("/:templateName/parameters", txpfiber.HandleParametersGet)

	gTemplate.All("/:templateName",
		func(c *fiber.Ctx) error {
			mth := c.Method()
			if mth == "POST" || mth == "GET" {
				return txpfiber.HandleCompile(c)
			}
			return c.Next()
		},
	)

	app.Get("/healthcheck", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	addr := ":" + os.Getenv("PORT")
	log.Infof("Starting server on port %s", addr)
	app.Listen(addr)
}
