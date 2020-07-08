package main

import "time"

type Temperature struct {
	City             string  `json:"city"`
	Country          string  `json:"country"`
	TemperatureScale string  `json:"temperature_scale"`
	TemperatureValue float64 `json:"temperature_value"`
}

type FluxTemperature struct {
	City             string    `json:"city"`
	Country          string    `json:"country"`
	TemperatureScale string    `json:"temperature_scale"`
	Time             time.Time `json:"_time"`
	Start            time.Time `json:"_start"`
	Stop             time.Time `json:"_stop"`
	Measurement      string    `json:"_measurement"`
	Field            string    `json:"_field"`
	Value            float64   `json:"_value"`
	Table            int       `json:"table"`
}

var Temperatures = []Temperature{
	{City: "Sanski Most", TemperatureScale: "Celsius", TemperatureValue: 19.28},
}

func Tags(t Temperature) map[string]string {
	newMap := make(map[string]string)
	newMap["city"] = t.City
	newMap["country"] = t.Country
	newMap["temperature_scale"] = t.TemperatureScale

	return newMap
}

func Fields(t Temperature) map[string]interface{} {
	newMap := make(map[string]interface{})
	newMap["temperature_value"] = t.TemperatureValue

	return newMap
}

func Measurement(t Temperature) string {
	return "temperatures"
}
