// Package cmd holds all the commands the user interacts with.
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

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a done task",
	Long:  `Add a done task to it's proper category.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Handle gig new category without description of task, it currently breaks the app
		category := args[0]
		task := args[1]
		id, err := db.CreateTask(category, task)
		if err != nil {
			log.Fatalf("Error while creating task \"%s\": %v", task, err)
		}
		fmt.Printf("Created new task %d successfully\n", id)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
