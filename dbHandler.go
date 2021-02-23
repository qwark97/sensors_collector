package main

import (
	"database/sql"
	"encoding/json"
	"io"
	"log"
)

var loadedDBConfig dbConfig

func loadDBConfig(r io.Reader) error {
	decoder := json.NewDecoder(r)

	if err := decoder.Decode(&loadedDBConfig); err != nil {
		return err
	}
	return nil
}

func loadToDB(db *sql.DB, measurement measure) {
	sqlStatement := insertStatements.getStatement(measurement.Category)
	_, err := db.Exec(
		sqlStatement,
		measurement.Origin,
		measurement.Unit,
		measurement.Value,
		measurement.SensorId,
	)
	if err != nil {
		log.Println("ERROR - inserting into DB -", err)
	}
}
