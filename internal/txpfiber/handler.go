package txpfiber

import (
	"strings"

	"github.com/charmbracelet/log"
	"github.com/gofiber/fiber/v2"
	"github.com/unhanded/txp/internal/fs"
)

func HandleCompile(c *fiber.Ctx) error {
	reqFormat := c.Query("format", "pdf")
	if reqFormat != "pdf" && reqFormat != "png" && reqFormat != "svg" {
		c.SendStatus(fiber.ErrNotImplemented.Code)
	}
	tName := c.Params("templateName", "")
	if tName == "" {
		return c.SendStatus(400)
	}

	if !IsOkayToCompileForGetMethod(c) {
		log.Warn("Compile failed becase user data is required")
		return c.SendStatus(fiber.ErrForbidden.Code)
	}

	if !authCheck(c) {
		log.Warn("Rejected because of auth fail", "ip", c.IP())
		return c.SendStatus(fiber.ErrUnauthorized.Code)
	}

	workingDir, prepErr := prepareForCompile(tName)
	if prepErr != nil {
		log.Error("Failed to prepare for compile", "err", prepErr)
		fs.UnbotheredDelete(workingDir)
		return c.SendStatus(500)
	}

	if err := populate(c, workingDir); err != nil {
		log.Error("Failed to populate for request.", "err", err)
		return c.SendStatus(500)
	}

	b, err := compileWithDir(workingDir, reqFormat)
	if err != nil {
		log.Error("Failed on compile", "err", err)
		return c.SendStatus(500)
	}

	log.Info("Sucessfully generated a document, proceeding to send", "work_id", strings.TrimPrefix(workingDir, "."))
	if err := Send(c, b, reqFormat); err != nil {
		log.Error("Failed to send response", "err", err)
	}
	fs.UnbotheredDelete(workingDir)

	return nil
}
