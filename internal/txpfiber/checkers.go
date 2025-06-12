package txpfiber

import (
	"strings"

	"github.com/charmbracelet/log"
	"github.com/gofiber/fiber/v2"
	"github.com/unhanded/txp/internal/environ"
	"github.com/unhanded/txp/internal/txppack"
)

func authCheck(c *fiber.Ctx) bool {
	if environ.TxpToken() == "" {
		return true
	}

	reqToken := c.Get("Authorization", "")

	return strings.HasSuffix(reqToken, environ.TxpToken())
}

func IsOkayToCompileForGetMethod(c *fiber.Ctx) bool {
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
