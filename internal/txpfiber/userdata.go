package txpfiber

import (
	"os"
	"path"
	"time"

	"github.com/charmbracelet/log"
	"github.com/gofiber/fiber/v2"
	"github.com/unhanded/txp/internal/fs"
)

var userDataFileName = "data.json"
var defaultDataFileName = "default.json"

func placeUserData(c *fiber.Ctx, workDir string) error {
	b := c.Body()
	if len(b) == 0 {
		log.Debug("No user data")
		if hasDefaultFallback(workDir) {
			log.Debug("Populating using default data")
			if defaultErr := useDefaultFallback(workDir); defaultErr != nil { // Oh dear, couldn't use default
				log.Error("Default data populate fail, falling back on zero-field data file", "err", defaultErr)
				b = []byte("{}")
			} else {
				return nil
			}
		} else {
			return nil
		}
	}

	f, createErr := os.Create(path.Join(workDir, userDataFileName))
	if createErr != nil {
		log.Error("failed to create file", "err", createErr.Error())
		return createErr
	}
	_, writeErr := f.Write(b)
	if writeErr != nil {
		log.Error("failed to write file", "err", writeErr.Error())
		return writeErr
	} else {
		time.Sleep(time.Millisecond * 10)
	}
	return nil
}

func hasDefaultFallback(workDir string) bool {
	defaultFilepath := path.Join(workDir, defaultDataFileName)
	return fs.FileExist(defaultFilepath)
}

func useDefaultFallback(workDir string) error {
	defaultFilepath := path.Join(workDir, defaultDataFileName)
	return fs.FileRename(defaultFilepath, userDataFileName)
}
