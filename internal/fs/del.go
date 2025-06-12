package fs

import (
	"os"

	"github.com/charmbracelet/log"
)

func FileDelete(fp string) error {
	return os.Remove(fp)
}

func UnbotheredDelete(fp string) {
	if err := os.RemoveAll(fp); err != nil {
		log.Warn("Delete failed", "err", err)
	}
}
