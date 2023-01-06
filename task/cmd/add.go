package cmd

import (
	"fmt"
	"os"
	"strings"
	"task/db"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a new task",
	Run: func(cmd *cobra.Command, arg []string) {
		task := strings.Trim(strings.Join(arg, " "), " ")
		_, err := db.CreateTask(task)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		fmt.Printf("Added '%s' to your task list \n", task)
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
