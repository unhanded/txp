package txppack

import (
	"fmt"
	"os"
	"path"
	"slices"

	"github.com/charmbracelet/log"
	"github.com/unhanded/txp/internal/dataman"
)

type TxpPack struct {
	Name     string
	FileList []string
}

func (tp TxpPack) Validate() error {
	log.Info("Got file list", "itemCount", len(tp.FileList))
	alwaysRequiredFiles := []string{"main.typ"}
	warnMissingFiles := []string{"txpfile.yml"}
	for _, req := range alwaysRequiredFiles {
		if !slices.ContainsFunc(
			tp.FileList,
			func(e string) bool {
				return path.Base(e) == req
			}) {
			return fmt.Errorf("missing required file '%s'", req)
		}
	}

	for _, file := range warnMissingFiles {
		if !slices.ContainsFunc(
			tp.FileList,
			func(e string) bool {
				return path.Base(e) == file
			}) {
			log.Warn("Missing expected file", "filename", file)
		}
	}

	return nil
}

func New(templateName string) (*TxpPack, error) {
	tPath := dataman.GetTemplatePath(templateName)
	log.Infof("Attempting to get template at path %s", tPath)
	if fi, err := os.Stat(tPath); err != nil {
		return nil, fmt.Errorf("Failed before enumerating with error: %s", err.Error())
	} else if !fi.IsDir() {
		return nil, fmt.Errorf("NOT_A_DIR")
	}

	fileList := unpackDirs([]string{}, tPath)
	return &TxpPack{Name: templateName, FileList: fileList}, nil
}

func unpackDirs(list []string, parentPath string) []string {
	files, err := os.ReadDir(parentPath)

	if err != nil {
		return list
	}

	for _, file := range files {
		pathSoFar := path.Join(parentPath, file.Name())
		log.Debugf("Inspecting path %s", pathSoFar)
		if file.IsDir() {
			list = unpackDirs(list, pathSoFar)
			continue
		}
		list = append(list, pathSoFar)
	}

	return list
}
