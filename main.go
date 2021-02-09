package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
)

var (
	dbConfigPath string = ".db_config.json"
)

func main() {
	// load configs
	file, err := os.Open(dbConfigPath)
	errHandle("Could not open provided DB config", err)

	err = loadDBConfig(file)
	errHandle("Could not load provided DB config", err)

	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		loadedDBConfig.Host,
		loadedDBConfig.Port,
		loadedDBConfig.User,
		loadedDBConfig.Password,
		loadedDBConfig.Dbname)

	// connect to DB
	db, err = sql.Open("postgres", psqlInfo)
	errHandle("Could not open DB connection", err)
	defer func() {
		err = db.Close()
		errHandle("Could not close DB connection", err)
	}()

	err = db.Ping()
	errHandle("Could not connect to DB", err)

	fmt.Println("Successfully connected!")

	// load sensors_config.json

	// based on sensors_config start one or many collectors as a goroutines
	// which will collect data and send it back via channel

	// infinite loop that will put data into db

}

func errHandle(msg string, err error) {
	if err != nil {
		log.Panicln(msg+"\n", err)
	}
}
