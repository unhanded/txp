package txpfiber

import (
	"fmt"
	"os"
	"path"

	"github.com/charmbracelet/log"
	"github.com/google/uuid"
	"github.com/unhanded/txp/internal/environ"
)

func NewWorkdir() string {
	workroot := environ.TxpWorkRoot()
	id := fmt.Sprintf("%s", uuid.New().String())
	wd := path.Join(workroot, id)
	if err := os.Mkdir(wd, 0755); err != nil {
		log.Error("Mkdir failure", "err", err)
	}
	return wd
}
