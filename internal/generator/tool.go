// internal/generator/function.go
package generator

import (
	"fmt"
	"github.com/kanlanc/restack-cli/internal/tools"
)

func CreateTool(name string) error {
	var content string
	switch name {
	case "mailgun":
		content = tools.MailgunTemplate
	default:
		return fmt.Errorf("tool template not found for: %s", name)
	}
	return createFile(fmt.Sprintf("src/functions/%s.py", name), content)
}




