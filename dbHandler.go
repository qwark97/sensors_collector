package main

import (
	"encoding/json"
	"io"
)

var loadedDBConfig dbConfig

func loadDBConfig(r io.Reader) error {
	decoder := json.NewDecoder(r)

	if err := decoder.Decode(&loadedDBConfig); err != nil {
		return err
	}
	return nil
}
