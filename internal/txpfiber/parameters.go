package txpfiber

import (
	"os"
	"slices"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/gofiber/fiber/v2"
	"github.com/unhanded/txp/internal/dataman"
)

func HandleParametersGet(c *fiber.Ctx) error {
	tName := c.Params("templateName", "")
	if tName == "" {
		return c.SendStatus(400)
	}

	tPaths, err := dataman.GetTemplateFilePaths(tName)
	if err != nil {
		log.Error("Could not get template filepaths", "templateName", tName, "err", err)
		return c.SendStatus(404)
	}

	if idx := slices.IndexFunc(
		tPaths,
		func(item string) bool {
			return strings.HasSuffix(item, "/data.json")
		},
	); idx == -1 {
		log.Error("Could not find data.json for template", "templateName", tName)
		return c.SendStatus(404)
	} else {
		return sendParametersFromFile(c, tPaths[idx])
	}
}

func sendParametersFromFile(c *fiber.Ctx, fp string) error {
	b, err := os.ReadFile(fp)
	if err != nil {
		c.SendStatus(500)
	}
	c.Write(b)

	withFormat(c, "application/json")

	return c.SendStatus(200)
}
