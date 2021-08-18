package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/angusbean/weather-check/models"
)

func main() {
	//Receive Args (Lat and Long) from command line
	if len(os.Args) != 3 {
		fmt.Println("error, usage required:", os.Args[0], "Latitude Value", "Longitude Value")
		os.Exit(1)
	}
	lat, long := os.Args[1], os.Args[2]
	fmt.Println(lat, long)

	// Open the city.list json file and handle erros
	jsonFile, err := os.Open("openweather-info/city.list.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Successfully opened the JSON file")
	defer jsonFile.Close()

	// Read opened jsonFile as a byte array
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// Initialise City array
	var citylist models.CityList

	// Unmarshal byteArray into cities struct
	json.Unmarshal(byteValue, &citylist)

	// Interate through every city in list
	for i := 0; i < len(citylist.CityList); i++ {
		if citylist.CityList[i].Name == "Oxford" {
			fmt.Println("Found Oxford")
		}
	}

	/*s
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
	*/
}
