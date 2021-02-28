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
	"strings"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Add new category",
	Long:  `Add new category to store done tasks.`,
	Run: func(cmd *cobra.Command, args []string) {
		catName := strings.Join(args, " ")
		id, err := db.CreateCategory(catName)
		if err != nil {
			log.Fatalf("Error while creating the \"%s\" category: %v", catName, err)
		}
		fmt.Printf("Created new category \"%s\" with ID %d\n", catName, id)
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
}
