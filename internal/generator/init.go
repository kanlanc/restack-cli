package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"os/exec"
)

func InitProject(projectName string) error {
	// Create project directory structure
	// Clone the quickstart template repository
	cmd := exec.Command("git", "clone", "https://github.com/kanlanc/restack-0.52-quickstart.git", projectName)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to clone template repository: %v\n%s", err, output)
	}

	// Remove the .git directory to start fresh
	err = os.RemoveAll(filepath.Join(projectName, ".git"))
	if err != nil {
		return fmt.Errorf("failed to remove .git directory: %v", err)
	}

	fmt.Printf("Successfully initialized project '%s' from quickstart template\n", projectName)
	return nil
}


