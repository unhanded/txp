package fs

import (
	"os"
	"path"
	"testing"
)

func TestFileRename(t *testing.T) {
	testStr := "{\"msg\":\"guten tag and hello\"}"
	testBytes := []byte(testStr)
	wd := t.TempDir()

	srcFileName := "file_a.json"
	srcFilePath := path.Join(wd, srcFileName)
	dstFileName := "file_b.json"
	dstFilePath := path.Join(wd, dstFileName)

	if err := os.WriteFile(srcFilePath, testBytes, 0755); err != nil {
		t.Fatalf("Failed to even write first file, err: %s", err.Error())
	}
	t.Log("File A in place")
	if err := FileRename(srcFilePath, dstFileName); err != nil {
		t.Fatal(err)
	}
	if b, err := os.ReadFile(dstFilePath); err != nil {
		t.Fatal(err)
	} else {
		if b[5] != testBytes[5] {
			t.Fatal("Content did not copy between files correctly.")
		}
	}

	t.Logf("Successfully renamed %s to %s", srcFileName, dstFileName)
}
