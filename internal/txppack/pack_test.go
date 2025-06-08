package txppack

import (
	"os"
	"path"
	"testing"
)

func TestPackNew(t *testing.T) {
	if wd, err := os.Getwd(); err != nil {
		t.Fatal(err)
	} else {
		txpDir := path.Join(wd, "../", "../", "/docs/example")
		t.Logf("Setting TXP_DIR to %s", txpDir)
		t.Setenv("TXP_DIR", txpDir)
	}
	pck, creationErr := New("demo")
	if creationErr != nil {
		t.Fatal(creationErr)
	}
	if err := pck.Validate(); err != nil {
		t.Error(err)
	}
}
