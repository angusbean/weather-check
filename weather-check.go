package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"time"

	"github.com/angusbean/weather-check/helpers"
	"github.com/angusbean/weather-check/models"
)

func main() {
	startTime := time.Now()

	//Receive Args (Lat and Long) from command line
	if len(os.Args) != 3 {
		fmt.Println("error, usage required:", os.Args[0], "Latitude Value", "Longitude Value")
		os.Exit(1)
	}

	//Convert latitude string value to float and check for errors
	lat, err := strconv.ParseFloat(os.Args[1], 32)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//Convert longitude string value to float and check for errors
	long, err := strconv.ParseFloat(os.Args[2], 32)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Open the city.list json file and handle erros
	jsonFile, err := os.Open("openweather-info/city.list.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer jsonFile.Close()

	//Read opened jsonFile as a byte array
	byteValue, _ := ioutil.ReadAll(jsonFile)

	//Initialise City array
	var citylist models.CityList

	//Unmarshal byteArray into cities struct
	json.Unmarshal(byteValue, &citylist)

	//Create global values for city location calculation
	var closestCityID int
	var closestCityName, cloestCityCountry string
	var latOffSet, longOffSet, tmpTotalOffSet, totalOffSet float64
	totalOffSet = 10000000.00

	// Interate through every city in list to determine which coords are closest
	for i := 0; i < len(citylist.CityList); i++ {
		latOffSet = math.Abs(lat - float64(citylist.CityList[i].Coord.Lat))
		longOffSet = math.Abs(long - float64(citylist.CityList[i].Coord.Long))
		tmpTotalOffSet = latOffSet + longOffSet
		if tmpTotalOffSet < totalOffSet {
			totalOffSet = tmpTotalOffSet
			closestCityID = citylist.CityList[i].ID
			closestCityName = citylist.CityList[i].Name
			cloestCityCountry = citylist.CityList[i].Country
		}
	}
	fmt.Println("Closest City:", closestCityName+",", cloestCityCountry, "ID:", closestCityID)
	helpers.TimeTrack(startTime, "factorial")

	/*
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
