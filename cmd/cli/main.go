// cmd/cli/main.go
package main

import (
	"fmt"
	"os"
	"path/filepath"
	"github.com/spf13/cobra"
	"github.com/kanlanc/restack-cli/internal/generator"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "restack",
		Short: "Restack.io workflow generator",
	}

	var workflowCmd = &cobra.Command{
		Use:   "workflow [name]",
		Short: "Generate a new workflow",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			err := generator.CreateWorkflow(args[0])
			if err != nil {
				fmt.Println("Error creating workflow:", err)
				os.Exit(1)
			}
			fmt.Printf("Workflow %s created successfully\n", args[0])
		},
	}

	var functionCmd = &cobra.Command{
		Use:   "function [name]",
		Short: "Generate a new function",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			err := generator.CreateFunction(args[0])
			if err != nil {
				fmt.Println("Error creating function:", err)
				os.Exit(1)
			}
			fmt.Printf("Function %s created successfully\n", args[0])
		},
	}

	var initCmd = &cobra.Command{
		Use:   "init [project-name]",
		Short: "Initialize a new Restack project",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			projectName := args[0]
			err := generator.InitProject(projectName)
			if err != nil {
				fmt.Println("Error initializing project:", err)
				os.Exit(1)
			}
			fmt.Printf("Project %s initialized successfully\n", projectName)
		},
	}

	rootCmd.AddCommand(workflowCmd, functionCmd, initCmd)
	rootCmd.Execute()
}
