package cmd

import (
	"fmt"
	"os"

	"cli-task/task"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all of your tasks",
	Run: func(cmd *cobra.Command, args []string) {
		manager, err := task.NewTaskManager()
		if err != nil {
			fmt.Println("Error initializing task manager:", err)
			os.Exit(1)
		}

		tasks, err := manager.LoadTasks()
		if err != nil {
			fmt.Println("Error loading tasks:", err)
			os.Exit(1)
		}

		if len(tasks) == 0 {
			fmt.Println("You have no tasks! Start by adding one.")
			return
		}

		fmt.Println("Your tasks:")
		for _, t := range tasks {
			status := " "
			if t.Completed {
				status = "x"
			}
			fmt.Printf("%d. [%s] %s\n", t.ID, status, t.Description)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
