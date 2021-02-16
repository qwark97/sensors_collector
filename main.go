package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
)

var (
	dbConfigPath      = ".db_config.json"
	sensorsConfigPath = ".sensors_config.json"
)

func main() {
	// load all configs
	loadConfigs()
	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		loadedDBConfig.Host,
		loadedDBConfig.Port,
		loadedDBConfig.User,
		loadedDBConfig.Password,
		loadedDBConfig.Dbname)

	// connect to DB
	db, err := sql.Open("postgres", psqlInfo)
	errHandle("Could not open DB connection", err)
	defer func() {
		err = db.Close()
		errHandle("Could not close DB connection", err)
	}()
	err = db.Ping()
	errHandle("Could not connect to DB", err)

	dataToDBChan := make(chan measure)

	// based on sensors_config start one or many collectors as a goroutines
	// which will collect data and send it back via channel
	for _, sensor := range loadedSensorsConfig {
		go collectData(sensor, dataToDBChan)
		log.Printf("INFO - start sensor: %s\n", sensor.Comment)
	}

	// infinite loop that will put data into db
	for {
		measurement, _ := <-dataToDBChan
		go loadToDB(measurement)
	}
}

func errHandle(msg string, err error) {
	if err != nil {
		log.Panicln(msg+"\n", err)
	}
}

func loadConfigs() {
	file, err := os.Open(dbConfigPath)
	errHandle("Could not open provided DB config", err)

	err = loadDBConfig(file)
	errHandle("Could not load provided DB config", err)

	file, err = os.Open(sensorsConfigPath)
	errHandle("Could not open provided sensors config", err)

	err = loadSensorsConfig(file)
	errHandle("Could not load provided sensors config", err)
}
