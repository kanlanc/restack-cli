// internal/generator/function.go
package generator

import (
	"fmt"
	"github.com/kanlanc/restack-cli/internal/templates"
)

func CreateFunction(name string) error {
	content := fmt.Sprintf(templates.FunctionTemplate, name)
	return createFile(fmt.Sprintf("src/functions/%s.py", name), content)
}
