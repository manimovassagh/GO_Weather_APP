package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type apiConfigData struct {
	openWeatherMapApiKey string `json:"openWeatherMapApiKey`
}

type weatherData struct {
	Name string `json:"name"`
	Main struct {
		Kelvin float64 `json:"temp"`
	} `json:"main"`
}

func loadApiConfig(filename string) (apiConfigData, error) {

	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return apiConfigData{}, err
	}
	var c apiConfigData
	err = json.Unmarshal(bytes, &c)
	if err != nil {
		return apiConfigData{}, err
	}
	return c, nil
}


func main() {
	http.HandleFunc("/hello", hello)

	http.HandleFunc("/weather", hello)

	http.ListenAndServe(":8080",nil)



}


