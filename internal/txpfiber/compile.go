package txpfiber

import (
	"os"
	"path"

	"github.com/charmbracelet/log"
	"github.com/gofiber/fiber/v2"
	"github.com/unhanded/txp/internal/dataman"
	"github.com/unhanded/txp/internal/fs"
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

func prepareForCompile(templateName string) (string, error) {
	wd := NewWorkdir()
	templatePath := dataman.GetTemplatePath(templateName)

	err := fs.CopyAll(templatePath, path.Join("./", wd))
	if err != nil {
		log.Error("Failed to copy files", "err", err)
		return "", err
	}

	return wd, nil
}

func populate(c *fiber.Ctx, workDir string) error {
	if err := placeUserData(c, workDir); err != nil {
		return err
	}

	return placeContext(c, workDir)
}
