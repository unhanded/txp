package txpfiber

import (
	"os"
	"path"

	"github.com/charmbracelet/log"
	"github.com/gofiber/fiber/v2"
	"github.com/unhanded/txp/internal/txpc"
)

func compileWithDir(dirName string, format string) ([]byte, error) {
	mainFileName := path.Join(dirName, "main.typ")
	typFileContent, tfcErr := os.ReadFile(mainFileName)
	if tfcErr != nil {
		return nil, tfcErr
	}
	tx, txErr := txpc.New()
	if txErr != nil {
		return nil, txErr
	}
	result, cErr := tx.Compile(typFileContent, dirName, format)
	if cErr != nil {
		return nil, cErr
	}
	return result, nil
}

func Send(c *fiber.Ctx, b []byte, format string) error {
	_, writeErr := c.Write(b)
	if writeErr != nil {
		log.Error("failed to write response", "err", writeErr.Error())
		return c.SendStatus(500)
	}
	withFormat(c, format)

	return c.SendStatus(200)
}

func withFormat(c *fiber.Ctx, format string) {
	mime := toMimeType(format)
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
