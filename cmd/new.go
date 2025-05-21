package cmd

import (
	"fmt"
	"github.com/petar-arandjic/gik/internal/scaffolder"
	"github.com/spf13/cobra"
)

var newCmd = &cobra.Command{
	Use:   "new [project name]",
	Short: "Create a new project",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		projectName := args[0]
		err := scaffolder.CreateProject(projectName)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("Project created:", projectName)
		}
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
}
