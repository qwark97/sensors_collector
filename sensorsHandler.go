package main

import (
	"encoding/json"
	"io"
)

var loadedSensorsConfig sensorsConfig

func loadSensorsConfig(r io.Reader) error {
	decoder := json.NewDecoder(r)

	if err := decoder.Decode(&loadedSensorsConfig); err != nil {
		return err
	}
	return nil
}
