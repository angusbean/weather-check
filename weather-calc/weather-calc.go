package weathercalc

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"

	"github.com/angusbean/weather-check/models"
)

//LoadCityList loads the JSON city list into memory
func LoadCityList(jFile *os.File) models.Cities {
	// Open the city.list.json file and handle erros
	jsonFile := jFile

	//Read opened jsonFile as a byte array
	byteValue, _ := ioutil.ReadAll(jsonFile)

	//Initialise City array
	var cityList models.Cities

	//Unmarshal byteArray into cities struct
	json.Unmarshal(byteValue, &cityList)
	return cityList
}

//LocateCity returns the closest City ID (based on OpenWeather file from lat and long provided)
func LocateCity(lat float64, long float64, cityList models.Cities) int {
	//Create global values for city location calculation
	var closestCityID int
	var latOffSet, longOffSet, tmpTotalOffSet, totalOffSet float64
	totalOffSet = 10000000.00

	//Interate through every city in list to determine which coords are closest
	for i := 0; i < len(cityList.Cities); i++ {
		latOffSet = math.Abs(lat - float64(cityList.Cities[i].Coord.Lat))
		longOffSet = math.Abs(long - float64(cityList.Cities[i].Coord.Long))
		tmpTotalOffSet = latOffSet + longOffSet
		if tmpTotalOffSet < totalOffSet {
			totalOffSet = tmpTotalOffSet
			closestCityID = cityList.Cities[i].ID
		}
	}
	return closestCityID
}

//RetrieveWeather returns the weather information based on the city ID from OpenWeather
func RetrieveWeather(closestCityID int) models.WeatherUpdate {
	//Recall the API Key from OS Environment variables set in main run()
	API_KEY := os.Getenv("API_KEY")

	//Create API call
	APICall := "http://api.openweathermap.org/data/2.5/weather?id=" + strconv.Itoa(closestCityID) + "&appid=" + API_KEY

	//Create client & request
	client := &http.Client{}
	req, err := http.NewRequest("GET", APICall, nil)
	if err != nil {
		log.Print(err)
	}

	//Add Request Headers and send
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Print(err)
	}
	defer resp.Body.Close()

	//Read Response Body into Memory as bytes
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Print(err)
	}

	//Unmarshal the bytes as json into the weather model
	var weatherModel models.WeatherUpdate
	json.Unmarshal(bodyBytes, &weatherModel)

	return weatherModel
}
