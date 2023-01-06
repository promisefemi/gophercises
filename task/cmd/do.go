/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"
	"task/db"

	"github.com/spf13/cobra"
)

// doCmd represents the do command
var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Mark a task as completed",
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int
		for _, i := range args {
			id, err := strconv.Atoi(i)
			if err != nil {
				fmt.Println("Could not parse argument:", i)
			} else {
				ids = append(ids, id)
			}
		}

		tasks, err := db.AllTask()
		if err != nil {
			fmt.Printf("Something went wrong: %s", err.Error())
			return
		}

		for _, id := range ids {
			if id <= 0 || id > len(tasks) {
				fmt.Printf("Invalid task number: %d \n", id)
				continue
			}

			task := tasks[id-1]
			err := db.DeleteTask(task.Key)
			if err != nil {
				fmt.Printf("Something went wrong, while deleting task #%d: Error %s", id, err)
				continue
			} else {
				fmt.Printf("Marked task  #%d as complete", id)
			}
		}

	},
}

func init() {
	RootCmd.AddCommand(doCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// doCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// doCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
