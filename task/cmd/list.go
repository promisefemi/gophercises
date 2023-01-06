/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"task/db"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all available tasks",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.AllTask()
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		if len(tasks) == 0 {
			fmt.Println("You have no task to complete.")
			return
		}

		fmt.Println("You have the following tasks:")
		for i, task := range tasks {
			fmt.Printf("%d.  %s\n", i+1, task.Value)
		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
