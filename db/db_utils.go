// Package db manages all the interactions related to the database.
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
	"encoding/binary"
	"time"
)

// itob turns integer into a byte slice
func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

// btoi turns a byte slice into an integer
func btoi(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}

// currentDate sets MM-DD-YYYY hh:mm:ss format to the current day
func currentDate() string {
	now := time.Now()
	return now.Format("01-02-2006 15:04:05")
}
