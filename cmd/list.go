/*
Copyright © 2021 Alan Rojas alandavidrl11@gmail.com

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"github.com/alandavd/gig/db"
	"github.com/spf13/cobra"
	"log"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List categories and tasks",
	Long: `List categories and tasks of a category.`,
	Run: func(cmd *cobra.Command, args []string) {
		sectionFlag := cmd.Flag("tasks")
		switch sectionFlag.Value.String() {
		case "true":
			category := args[0]
			exists := db.BucketExists(category)
			if !exists {
				fmt.Println("This category does not exists.")
				return
			}
			tasks, err := db.ListTasks(category)
			if err != nil {
				log.Fatal(err)
			}
			if len(tasks) == 0 {
				fmt.Println("You don't have tasks yet.")
				return
			}
			fmt.Println("You have the following tasks:")
			for i, task := range tasks {
				fmt.Printf("%d. %s\n", i+1, task.Value)
			}
		case "false":
			categories, err := db.ListCategories()
			if err != nil {
				log.Fatal(err)
			}
			if len(categories) == 0 {
				fmt.Println("You have no categories.")
				return
			}
			fmt.Println("You have the following categories:")
			for i, category := range categories {
				fmt.Printf("%d. %s\n", i+1, category.Key)
			}
		default:
			fmt.Println("Upps I don't know what happened")
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Local flags
	listCmd.Flags().BoolP("tasks", "t", false, "List categories or tasks")
}
