package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func getAllTemperatures(w http.ResponseWriter, r *http.Request) {
	records := Read("temperatures")

	var temperatures []FluxTemperature
	for _, data := range records {
		var temp FluxTemperature
		if err := json.Unmarshal(data, &temp); err != nil {
			panic(err)
		}
		temperatures = append(temperatures, temp)
	}
	enc := json.NewEncoder(w)
	enc.SetIndent("", "\t")
	enc.Encode(temperatures)
}

func postTemperature(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()
	if err = checkEmptyBody(string(body)); err != nil {
		json.NewEncoder(w).Encode(fmt.Sprintf(`Error: %v`, err.Error()))
		return
	}
	var temperature Temperature
	if err = json.Unmarshal(body, &temperature); err != nil {
		json.NewEncoder(w).Encode(fmt.Sprintf(`Error in request body: %v`, err.Error()))
		return
	}
	Write(temperature)
	json.NewEncoder(w).Encode("New point successfully written !")
}
