package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go"
	"github.com/sirupsen/logrus"
)

func mountedConnectionParameters() map[string]string {
	connectionParams := make(map[string]string)
	basePath := "/app/influxdb"
	files, err := ioutil.ReadDir(basePath)
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		if strings.HasPrefix(file.Name(), ".") == false {
			fileContent, err := ioutil.ReadFile(fmt.Sprintf("%v/%v", basePath, file.Name()))
			if err != nil {
				logrus.Errorf("Could not read file %v", file.Name())
			}
			connectionParams[file.Name()] = string(fileContent)
		}
	}
	return connectionParams
}

// ConnectionParameters contains all the required parameters for talking to InfluxDB
var ConnectionParameters = mountedConnectionParameters()

// Write function writes temeperature measurement in in bucket defined in .env file
func Write(t Temperature) {
	client := influxdb2.NewClient(ConnectionParameters["url"], ConnectionParameters["token"])
	writeAPI := client.WriteApi(ConnectionParameters["org"], ConnectionParameters["bucket"])
	p := influxdb2.NewPoint(Measurement(t), Tags(t), Fields(t), time.Now())
	writeAPI.WritePoint(p)
}

// Read functions reads all the temperatures saved inside of InfluxDB and returns them as array
func Read(measurement string) [][]byte {
	client := influxdb2.NewClient(ConnectionParameters["url"], ConnectionParameters["token"])
	queryAPI := client.QueryApi(ConnectionParameters["org"])
	fluxQuery := fmt.Sprintf(`from(bucket:"%v") |> range(start:-5) |> filter(fn:(r) => r._measurement == "%v")`,
		ConnectionParameters["bucket"], measurement)
	logrus.Infof("FLUX QUERY: %v", fluxQuery)
	result, err := queryAPI.Query(context.Background(), fluxQuery)
	if err != nil {
		logrus.Errorf("Error :%v", err)
		panic(err)
	}
	var temperaturesArray [][]byte
	for result.Next() {
		jsonn, err := json.Marshal(result.Record().Values())
		if err != nil {
			panic(err)
		}
		temperaturesArray = append(temperaturesArray, jsonn)
	}

	return temperaturesArray
}
