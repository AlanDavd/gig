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
package db

import (
	"encoding/json"
	"github.com/boltdb/bolt"
	"log"
	"time"
)

var (
	db               *bolt.DB
	categoriesBucket = []byte("gigdb")
)

// Category is the representation of a nested bucket that holds tasks
type Category struct {
	Key string `json:"category"`
}

// Task is the representation of a task that is done by the being time
type Task struct {
	Key     int    `json:"id"`
	Value   string `json:"description"`
	Created string `json:"created"`
}

// JSONCategory holds a slice of tasks as the only key.
// It is also the representation of a JSON schema of categories
// and their tasks.
type JSONCategory struct {
	Category string `json:"category"`
	Tasks    []Task `json:"tasks"`
}

// Init exposes database initialization to external clients
func Init(dbPath string) {
	if err := initDb(dbPath); err != nil {
		log.Fatal(err)
	}
}

// initDb opens Bolt connection and creates database file
func initDb(dbPath string) error {
	var err error
	db, err = bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return err
	}

	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(categoriesBucket)
		return err
	})
}

// CreateCategory creates a nested bucket within the main database bucket
func CreateCategory(name string) (int, error) {
	var id int
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(categoriesBucket)
		id64, err := b.NextSequence()
		if err != nil {
			return err
		}
		id = int(id64)
		_, err = b.CreateBucketIfNotExists([]byte(name))
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return -1, err
	}
	return id, nil
}

// ListCategories returns a list of all the buckets in database
func ListCategories() ([]Category, error) {
	var categories []Category
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(categoriesBucket)
		c := b.Cursor()
		for k, _ := c.First(); k != nil; k, _ = c.Next() {
			categories = append(categories, Category{Key: string(k)})
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return categories, nil
}

// CreateTask adds done task to it's proper category
func CreateTask(category, payload string) (int, error) {
	var id int
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(categoriesBucket)
		cb := b.Bucket([]byte(category))
		id64, err := cb.NextSequence()
		if err != nil {
			return err
		}
		id = int(id64)
		key := itob(id)
		now := currentDate()
		task := Task{
			Key:     id,
			Value:   payload,
			Created: now,
		}
		data, err := json.Marshal(task)
		if err != nil {
			return err
		}
		return cb.Put(key, data)
	})
	if err != nil {
		return -1, err
	}
	return id, nil
}

// ListTasks returns all the done tasks related to a category
func ListTasks(category string) ([]Task, error) {
	var tasks []Task
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(categoriesBucket)
		cb := b.Bucket([]byte(category))
		c := cb.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			var data Task
			json.Unmarshal(v, &data)
			tasks = append(tasks, Task{
				Key:     btoi(k),
				Value:   data.Value,
				Created: data.Created,
			})
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

// ExportData returns JSON data from all database data
func ExportData() ([]JSONCategory, error) {
	var categories []JSONCategory
	// Open database connection
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(categoriesBucket) // Main database file
		c := b.Cursor()
		// Loop through all categories
		for k, _ := c.First(); k != nil; k, _ = c.Next() {
			var categoryTasks []Task // List of tasks by category
			var jsonCategory JSONCategory
			cb := b.Bucket(k) // Get inside the bucket of each category
			tc := cb.Cursor() // Tasks cursor
			for tk, v := tc.First(); tk != nil; tk, v = tc.Next() {
				categoryTasks = append(categoryTasks, Task{
					Key:   btoi(tk),
					Value: string(v),
				})
			}
			jsonCategory.Category = string(k)
			jsonCategory.Tasks = categoryTasks
			categories = append(categories, jsonCategory)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return categories, nil
}

// BucketExists returns whether a Bucket exists
func BucketExists(name string) bool {
	exists, err := bucketExists(name)
	if err != nil {
		log.Fatal(err)
	}
	return exists
}

// bucketExists returns whether a Bucket exists and handles errors
func bucketExists(name string) (bool, error) {
	var exists bool
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(categoriesBucket)
		cb := b.Bucket([]byte(name))
		if cb != nil {
			exists = true
		}
		return nil
	})
	if err != nil {
		return exists, err
	}
	return exists, err
}
