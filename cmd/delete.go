package cmd

import (
	"fmt"
	"os"
	"strconv"

	"cli-task/task"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a task from your task list",
	Run: func(cmd *cobra.Command, args []string) {
		manager, err := task.NewTaskManager()
		if err != nil {
			fmt.Println("Error initializing task manager:", err)
			os.Exit(1)
		}

		var ids []int
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Printf("Failed to parse the argument %s as integer\n", arg)
				continue
			}
			ids = append(ids, id)
		}

		for _, id := range ids {
			err := manager.DeleteTask(id)
			if err != nil {
				fmt.Printf("Failed to delete task \"%d\": %v\n", id, err)
			} else {
				fmt.Printf("Deleted task \"%d\".\n", id)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
