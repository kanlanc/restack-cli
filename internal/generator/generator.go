// internal/generator/generator.go
package generator

import (
	"fmt"
	"os"
	"path/filepath"
)

func createFile(path string, content string) error {
	err := os.MkdirAll(filepath.Dir(path), 0755)
	if err != nil {
		return err
	}
	return os.WriteFile(path, []byte(content), 0644)
}