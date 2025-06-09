package dataman

import (
	"os"
	"path"

	"github.com/unhanded/txp/internal/environ"
)

func GetTemplateFilePaths(templateName string) ([]string, error) {
	results := []string{}

	tDirPath := GetTemplatePath(templateName)

	if dirInfo, err := os.ReadDir(tDirPath); err != nil {
		return nil, err
	} else {
		for _, entry := range dirInfo {
			if entry.IsDir() {
				continue
			}
			resultingPath := path.Join(tDirPath, entry.Name())
			results = append(results, resultingPath)
		}
	}
	return results, nil
}

func GetTemplatePath(templateName string) string {
	return path.Join(environ.TxpDir(), "/templates", templateName)

}
