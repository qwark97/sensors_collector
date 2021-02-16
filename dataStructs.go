package main

type sensorsConfig []sensorConfig

type sensorConfig struct {
	Address    string `json:"address"`
	SensorName string `json:"sensor_name"`
	Comment    string `json:"comment"`
	Frequency  int32  `json:"frequency"`
	Unit       string `json:"unit"`
	MACAddress string `json:"mac_address"`
}

type dbConfig struct {
	Host     string `json:"host"`
	Port     int32  `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Dbname   string `json:"dbname"`
}

type measure struct {
	Origin   string
	Unit     string
	Value    float32
	SensorId string
}
