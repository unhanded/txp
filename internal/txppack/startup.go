package txppack

import (
	"os"
	"path"

	"github.com/charmbracelet/log"
	"github.com/unhanded/txp/internal/environ"
)

func StartupCheck() {
	if templates, err := GetTemplates(); err != nil {
		log.Error("Failed to look up templates", "err", err)
	} else {
		for _, t := range templates {
			log.Info("Got template directory", "templateName", t)
		}
	}
}

func GetTemplates() ([]string, error) {
	out := []string{}
	templatesDir := path.Join(environ.TxpDir(), "templates")
	templates, err := os.ReadDir(templatesDir)
	if err != nil {
		return nil, err
	}
	for _, item := range templates {
		if item.IsDir() {
			out = append(out, item.Name())
		}
	}
	return out, nil
}
