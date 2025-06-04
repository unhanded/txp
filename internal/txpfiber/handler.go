package txpfiber

import (
	"path"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/gofiber/fiber/v2"
	"github.com/unhanded/txp/internal/dataman"
	"github.com/unhanded/txp/internal/fs"
)

func HandleCompile(c *fiber.Ctx) error {
	reqFormat := c.Query("format", "pdf")
	if reqFormat != "pdf" && reqFormat != "png" {
		c.SendStatus(fiber.ErrNotImplemented.Code)
	}
	tName := c.Params("templateName", "")
	if tName == "" {
		return c.SendStatus(400)
	}
	wd, prepErr := prepCompile(tName)
	if prepErr != nil {
		fs.UnbotheredDelete(wd)
		return c.SendStatus(500)
	}
	if err := populate(c, wd); err != nil {
		log.Error("Failed to populate for request.", "err", err)
		return c.SendStatus(500)
	}
	b, err := compileWithDir(wd, reqFormat)

	if err != nil {
		log.Error("Failed on compile", "err", err)
		return c.SendStatus(500)
	}

	log.Info("Sucessfully generated a document, proceeding to send", "work_id", strings.TrimPrefix(wd, "."))
	if err := SendPdf(c, b); err != nil {
		log.Error("Failed to send response", "err", err)
	}
	fs.UnbotheredDelete(wd)

	return nil
}

func prepCompile(templateName string) (string, error) {
	wd := NewWorkdir()
	templatePath := dataman.GetTemplatePath(templateName)

	err := fs.CopyAll(templatePath, path.Join("./", wd))
	if err != nil {
		log.Error("Failed to copy files", "err", err)
		return "", err
	}
	return wd, nil
}

func populate(c *fiber.Ctx, targetDir string) error {
	if err := placeUserData(c, targetDir); err != nil {
		return err
	}

	return placeContext(c, targetDir)
}
