package main

import (
	"database/sql"
	"encoding/json"
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
