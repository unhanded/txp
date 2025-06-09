package txpfiber

import (
	"path"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/gofiber/fiber/v2"
	"github.com/unhanded/txp/internal/dataman"
	"github.com/unhanded/txp/internal/environ"
	"github.com/unhanded/txp/internal/fs"
	"github.com/unhanded/txp/internal/txppack"
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

	if !getMethodCheck(c) {
		log.Info("Compile failed becase user data is required")
		return c.SendStatus(fiber.ErrForbidden.Code)
	}

	if !authCheck(c) {
		log.Info("Rejected because of auth fail", "ip", c.IP())
		return c.SendStatus(fiber.ErrUnauthorized.Code)
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

func authCheck(c *fiber.Ctx) bool {
	if environ.TxpToken() == "" {
		return true
	}

	reqToken := c.Get("Authorization", "")

	return strings.HasSuffix(reqToken, environ.TxpToken())
}

func getMethodCheck(c *fiber.Ctx) bool {
	if c.Method() != "GET" {
		return true
	}

	tName := c.Params("templateName", "")
	pack, packErr := txppack.New(tName)
	pack.Validate()
	if packErr != nil {
		log.Error("Pack parsing failed", "err", packErr)
		return false
	}
	if pack.Info == nil {
		return false
	}

	return pack.Info.CompileEmpty
}
