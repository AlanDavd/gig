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
	"github.com/boltdb/bolt"
	"log"
	"time"
)

var (
	db *bolt.DB
	categoriesBucket = []byte("gigdb")
)

// Category is the representation of a nested bucket that holds tasks
type Category struct {
	Key string
}

// Task is the representation of a task that is done by the being time
type Task struct {
	Key int
	Value string
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
