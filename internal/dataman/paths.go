package dataman

import (
	"os"
	"path"
)

func GetTemplateFilePaths(templateName string) ([]string, error) {
	results := []string{}

	tDirPath := path.Join(GetTxpDir(), "/templates", templateName)

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
