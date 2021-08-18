package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/angusbean/weather-check/models"
	"github.com/angusbean/weather-check/secrets"
)

func main() {
	//Recall API Key from secrets
	APICall := "http://api.openweathermap.org/data/2.5/weather?q=Sydney&appid=" + secrets.API_key

	// Create client & request
	client := &http.Client{}
	req, err := http.NewRequest("GET", APICall, nil)
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

	var responseObject models.Weather
	json.Unmarshal(bodyBytes, &responseObject)
	fmt.Printf("API Response as struct %+v\n", responseObject)

}
