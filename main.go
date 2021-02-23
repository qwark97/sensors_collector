package main

import (
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
)

var (
	dbConfigPath      *string
	sensorsConfigPath *string
)

func main() {
	dbConfigPath = flag.String("dbConfigPath", ".db_config.json", "Path to JSON DB config")
	sensorsConfigPath = flag.String("sensorsConfigPath", ".sensors_config.json", "Path to JSON sensors config")
	flag.Parse()

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
	errHandle("Could not ping DB", err)

	// load DB insert statements
	insertStatements.Statements = make(map[string]string)
	insertStatements.loadStatements()

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
		go loadToDB(db, measurement)
	}
}

func errHandle(msg string, err error) {
	if err != nil {
		log.Panicln(msg+"\n", err)
	}
}

func loadConfigs() {
	file, err := os.Open(*dbConfigPath)
	errHandle("Could not open provided DB config", err)

	err = loadDBConfig(file)
	errHandle("Could not load provided DB config", err)

	file, err = os.Open(*sensorsConfigPath)
	errHandle("Could not open provided sensors config", err)

	err = loadSensorsConfig(file)
	errHandle("Could not load provided sensors config", err)
}
