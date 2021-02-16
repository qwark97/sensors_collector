package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

func collectData(config sensorConfig, respChan chan measure) {
	var (
		measurement         string
		measurementResponse measure
		value               float32
	)
	for {
		log.Printf("INFO - measure: %s\n", config.SensorName)
		measurement = func() string {
			resp, err := http.Get(config.Address)
			if err != nil {
				log.Println("ERROR - connection -", err)
				return ""
			}
			defer func() {
				err := resp.Body.Close()
				if err != nil {
					log.Println("ERROR - body closing -", err)
					return
				}
			}()
			if resp.StatusCode != 200 {
				log.Println("ERROR - sensor response -", resp.Status)
				return ""
			}
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Println("ERROR - response reading -", err)
				return ""
			}
			res := string(body)
			return res
		}()
		valueStep, err := strconv.ParseFloat(measurement, 32)
		value = float32(valueStep)
		if err != nil {
			log.Println("ERROR - sensor's response value -", err)
			continue
		}
		measurementResponse = measure{
			Origin:   config.SensorName,
			Unit:     config.Unit,
			Value:    value,
			SensorId: config.MACAddress,
		}
		respChan <- measurementResponse
		sleep := time.Duration(config.Frequency)
		time.Sleep(sleep * time.Second)
	}
}
