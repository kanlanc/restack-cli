package generator

// import (
// 	"os"
// 	"path/filepath"
// )

func InitProject(projectName string) error {
	// Create project directory structure
	// dirs := []string{
	// 	"src/functions",
	// 	"src/workflows",
	// }

	// files := map[string]string{
	// 	"src/__init__.py":          "",
	// 	"src/functions/__init__.py": "",
	// 	"src/workflows/__init__.py":  "",
	// 	"src/workflows/workflow.py":   getBasicWorkflowTemplate(),
	// 	"src/functions/function.py":  getBasicFunctionTemplate(),
	// 	"src/app.py":                  getFastAPIAppTemplate(),
	// 	"src/client.py":               getClientTemplate(),
	// 	"src/services.py":            getServicesTemplate(),
	// 	// ".env":                    getEnvTemplate(),
	// 	".env.example":            getEnvExampleTemplate(),
	// 	".gitignore":              getGitignoreTemplate(),
	// 	"Dockerfile":              getDockerfileTemplate(),
	// 	// "frontend.py":             getFrontendTemplate(),
	// 	"nginx.conf":              getNginxConfTemplate(),
	// 	"pyproject.toml":          getPyprojectTemplate(),
	// 	"README.md":               getReadmeTemplate(projectName),
	// 	"schedule_workflow.py":    getScheduleWorkflowTemplate(),
	// }

	// // Create base project directory
	// if err := os.MkdirAll(projectName, 0755); err != nil {
	// 	return err
	// }

	// // Create directory structure
	// for _, dir := range dirs {
	// 	if err := os.MkdirAll(filepath.Join(projectName, dir), 0755); err != nil {
	// 		return err
	// 	}
	// }

	// // Create files
	// for file, content := range files {
	// 	if err := os.WriteFile(filepath.Join(projectName, file), []byte(content), 0644); err != nil {
	// 		return err
	// 	}
	// }

	return nil
}

// Add template functions below:
// func getAppTemplate() string {
// 	return `# Your app.py template content here
// `
// }

// func getClientTemplate() string {
// 	return `# Your client.py template content here
// `
// }

// Add similar functions for other templates...
