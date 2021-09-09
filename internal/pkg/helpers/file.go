package helpers

import (
	"bytes"
	"os"
)

var DefaultPerm os.FileMode = 0755

func FileReplaceContent(path string, old, new []byte) error {
	content, e := os.ReadFile(path)

	if e == nil {
		new := bytes.ReplaceAll(content, old, new)
		return os.WriteFile(path, new, DefaultPerm)
	}

	return e
}
