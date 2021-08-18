package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//Weather provides important information about weather data
type Weather struct {
	LocationName string `json:"name"`
	Coord        struct {
		Lat  float32 `json:"lat"`
		Long float32 `json:"lon"`
	} `json:"coord"`
	Weather []struct {
		Weather     string `json:"main"`
		Description string `json:"description"`
	} `json:"weather"`
	Temperature struct {
		Temperature    float32 `json:"temp"`
		MinTemperature float32 `json:"temp_min"`
		MaxTemperature float32 `json:"temp_max"`
	} `json:"main"`
	Code int `json:"cod"`
}

//API Key provided by Angus Bean. Expires:
var API_key = "8116a3514f24c7e9bc73a2b96bfeaa70"

func main() {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://api.openweathermap.org/data/2.5/weather?q=Sydney&appid=8116a3514f24c7e9bc73a2b96bfeaa70", nil)
	if err != nil {
		fmt.Print(err.Error())
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err.Error())
	}

	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err.Error())
	}

	var responseObject Weather
	json.Unmarshal(bodyBytes, &responseObject)
	fmt.Printf("API Response as struct %+v\n", responseObject)
}
