package txppack

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/log"
	yaml "github.com/goccy/go-yaml"
	"github.com/unhanded/txp/internal/fs"
)

// TODO: Get handlers running to accept GET requests if it's static etc
type TxpConf struct {
	Static bool `json:"static" yaml:"static"`
}

func loadConfig(fp string) (*TxpConf, error) {
	if !fs.FileExist(fp) {
		log.Error("File not found", "filepath", fp)
		return nil, fmt.Errorf("file \"%s\" not found", fp)
	}
	b, readErr := os.ReadFile(fp)
	if readErr != nil {
		return nil, readErr
	}
	d := TxpConf{}
	if strings.HasSuffix(fp, ".yml") || strings.HasSuffix(fp, ".yaml") {
		err := yaml.Unmarshal(b, &d)
		if err != nil {
			return nil, err
		}
		return &d, nil
	}
	err := json.Unmarshal(b, &d)
	if err != nil {
		return nil, err
	}
	return &d, nil
}
