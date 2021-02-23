package main

var insertStatements SQLInsertStatements

func (statements *SQLInsertStatements) loadStatements() {
	statements.Statements["temperature"] = `
INSERT INTO temperatures (origin, unit, temperature, sensor_id)
VALUES ($1, $2, $3, $4)`

	statements.Statements["humidity"] = `
INSERT INTO humidities (origin, unit, humidity, sensor_id)
VALUES ($1, $2, $3, $4)`
}

func (statements *SQLInsertStatements) getStatement(category string) string {
	val := statements.Statements[category]
	if val == "" {
		panic("Invalid SQL statement category")
	}
	return val
}
