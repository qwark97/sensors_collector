package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
)

var loadedDBConfig dbConfig
var db *sql.DB

func loadDBConfig(r io.Reader) error {
	decoder := json.NewDecoder(r)

	if err := decoder.Decode(&loadedDBConfig); err != nil {
		return err
	}
	return nil
}

func loadToDB(measurement measure) {
	fmt.Println(measurement)
}
