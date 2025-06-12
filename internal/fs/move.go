package fs

import (
	"path"

	"github.com/charmbracelet/log"
)

func FileMove(src string, dst string) error {
	if err := FileCopy(src, dst); err != nil {
		return err
	}

	return FileDelete(src)
}

func FileRename(fp string, newName string) error {
	containingDir := path.Dir(fp)
	oldName := path.Base(fp)
	log.Debug("Renaming file", "from", oldName, "to", newName, "in", containingDir)
	newFilePath := path.Join(containingDir, newName)
	if err := FileMove(fp, newFilePath); err != nil {
		log.Debug("Failed to rename file", "err", err)
		return err
	}

	return nil
}
