// internal/generator/workflow.go
package generator

import (
	"fmt"
	"github.com/kanlanc/restack-cli/internal/templates"
)

func CreateWorkflow(name string) error {
   // Takes a template (WorkflowTemplate) and injects the 'name' parameter twice into it
   content := fmt.Sprintf(templates.WorkflowTemplate, name, name)
   
   // Creates a new file in the src/workflows directory with the generated content
   return createFile(fmt.Sprintf("src/workflows/%s.ts", name), content)
}
