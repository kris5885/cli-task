package cmd

import (
	"fmt"
	"os"
	"strconv"

	"cli-task/task"
	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Mark a task as complete",
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
			err := manager.DoTask(id)
			if err != nil {
				fmt.Printf("Failed to mark task \"%d\" as complete: %v\n", id, err)
			} else {
				fmt.Printf("Marked task \"%d\" as complete.\n", id)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(doCmd)
}
