package txppack

import (
	"os"

	"github.com/goccy/go-yaml"
)

type TxpFileInfo struct {
	CompileEmpty bool `yaml:"compile_empty"`
}

func (tfi *TxpFileInfo) defaults() error {
	tfi.CompileEmpty = false
	return nil
}

func (tfi *TxpFileInfo) FromFile(fp string) error {
	tfi.defaults()
	f, err := os.ReadFile(fp)
	if err != nil {
		return err
	}
	dummy := TxpFileInfo{}
	if err := yaml.Unmarshal(f, &dummy); err != nil {
		return err
	}

	yaml.Unmarshal(f, tfi)

	return nil
}
