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
	"encoding/json"
	"fmt"
	"github.com/alandavd/gig/db"
	"github.com/spf13/cobra"
	"log"
	"os"
)

// exportCmd represents the export command
var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export your data to JSON.",
	Long:  `Export your tasks and categories to a JSON file.`,
	Run: func(cmd *cobra.Command, args []string) {
		data, err := db.ExportData()
		if err != nil {
			log.Fatal(err)
		}
		jsonData, err := json.MarshalIndent(data, "", "    ")
		if err != nil {
			log.Fatal(err)
		}
		f, err := os.Create("gigHistory.json")
		defer f.Close()
		if err != nil {
			log.Fatal(err)
		}
		if _, err = f.Write(jsonData); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Your database history was written to your current folder successfully.\n")
	},
}

func init() {
	rootCmd.AddCommand(exportCmd)
}
