package fs

import (
	"fmt"
	"os"
	"path"

	"github.com/charmbracelet/log"
)

type FileManager struct {
	BaseDir string
	ready   bool
}

func (fm *FileManager) Init() {
	if fi, err := os.Stat(fm.BaseDir); err != nil {
		log.Error("Failure in base directory check", "err", err)
	} else {
		if fi.IsDir() {
			fm.ready = true
		}
	}
}

func CopyAll(srcDir string, dstDir string) error {
	if !IsValidDir(dstDir) || !IsValidDir(dstDir) {
		return fmt.Errorf("INVALID_DIR")
	}
	dInfo, _ := os.ReadDir(srcDir)
	for _, item := range dInfo {
		fullSrc := path.Join(srcDir, item.Name())
		fullDst := path.Join(dstDir, item.Name())
		err := FileCopy(fullSrc, fullDst)
		if err != nil {
			log.Error("Failed to copy file",
				"src", fullSrc,
				"dst", fullDst,
				"err", err,
			)
		}
	}
	return nil
}

func FileCopy(srcFp string, dstFp string) error {
	if b, err := os.ReadFile(srcFp); err != nil {
		return err
	} else {
		return os.WriteFile(dstFp, b, 0755)
	}
}

func IsValidDir(fp string) bool {
	if fi, err := os.Stat(fp); err != nil {
		return false
	} else if !fi.IsDir() {
		return false
	}

	return true
}

func FileExist(fp string) bool {
	if fi, err := os.Stat(fp); err != nil {
		return false
	} else {
		if fi.IsDir() {
			return false
		}
	}
	return true
}
