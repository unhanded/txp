package txppack

import (
	"fmt"
	"os"
	"path"
	"slices"
)

type TxpPack struct {
	FileList []string
}

func (tp TxpPack) Validate() error {
	alwaysRequiredFiles := []string{"main.typ"}

	for _, req := range alwaysRequiredFiles {
		if !slices.ContainsFunc(
			tp.FileList,
			func(e string) bool {
				return path.Base(e) == req
			}) {
			return fmt.Errorf("missing required file '%s'", req)
		}
	}
	return nil
}

func Load(dirpath string) (*TxpPack, error) {
	if fi, err := os.Stat(dirpath); err != nil {
		return nil, fmt.Errorf("Failed before enumerating with error: %s", err.Error())
	} else if !fi.IsDir() {
		return nil, fmt.Errorf("NOT_A_DIR")
	}

	fileList := unpackDirs([]string{}, dirpath)
	return &TxpPack{FileList: fileList}, nil
}

func unpackDirs(list []string, parentPath string) []string {
	files, err := os.ReadDir(parentPath)

	if err != nil {
		return list
	}

	for _, file := range files {
		pathSoFar := path.Join(parentPath, file.Name())

		if file.IsDir() {
			list = unpackDirs(list, pathSoFar)
			continue
		}

		list = append(list, pathSoFar)
	}

	return list
}
