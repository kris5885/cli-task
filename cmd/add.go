package cmd

import (
	"fmt"
	"os"
	"strings"

	"cli-task/task"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task to your task list",
	Run: func(cmd *cobra.Command, args []string) {
		manager, err := task.NewTaskManager()
		if err != nil {
			fmt.Println("Error initializing task manager:", err)
			os.Exit(1)
		}

		description := strings.Join(args, " ")
		if description == "" {
			fmt.Println("Please provide a task description.")
			return
		}

		t, err := manager.AddTask(description)
		if err != nil {
			fmt.Println("Error adding task:", err)
			os.Exit(1)
		}

		fmt.Printf("Added \"%s\" to your task list.\n", t.Description)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
