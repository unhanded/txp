package fs

import (
	"os"

	"github.com/charmbracelet/log"
)

func UnbotheredDelete(fp string) {
	if err := os.RemoveAll(fp); err != nil {
		log.Warn("Delete failed", "err", err)
	}
}
