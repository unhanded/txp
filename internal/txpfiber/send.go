package txpfiber

import (
	"github.com/charmbracelet/log"
	"github.com/gofiber/fiber/v2"
)

func Send(c *fiber.Ctx, b []byte, format string) error {
	_, writeErr := c.Write(b)
	if writeErr != nil {
		log.Error("failed to write response", "err", writeErr.Error())
		return c.SendStatus(500)
	}
	withFormatFromAlias(c, format)

	return c.SendStatus(200)
}

func withFormatFromAlias(c *fiber.Ctx, format string) {
	mime := toMimeType(format)
	withFormat(c, mime)
}

func withFormat(c *fiber.Ctx, mime string) {
	c.Response().Header.SetContentType(mime)
}

func toMimeType(format string) string {
	switch format {
	case "png":
		return "image/png"
	case "svg":
		return "image/svg+xml"
	default:
		return "application/pdf"
	}
}
