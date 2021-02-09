package main

import (
	"log"
	"os"
)

var (
	dbConfigPath string = ".db_config.json"
)

func main() {
	file, err := os.Open(dbConfigPath)
	errHandle("Could not open provided DB config", err)

	err = loadDBConfig(file)
	errHandle("Could not load provided DB config", err)

}

func errHandle(msg string, err error) {
	if err != nil {
		log.Panic(msg, err)
	}
}
