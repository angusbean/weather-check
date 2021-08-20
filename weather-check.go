package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/angusbean/weather-check/models"
	weathercalc "github.com/angusbean/weather-check/weather-calc"
)

const portNumber = ":3000"

var weatherInfo models.Weather

func main() {
	toPrint := fmt.Sprintf("Staring application on port %s", portNumber)
	fmt.Println(toPrint)

	//Receive Args (Lat and Long) from command line
	if len(os.Args) != 3 {
		fmt.Println("input error, usage required:", os.Args[0], "Latitude Value", "Longitude Value")
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

	weatherInfo = weathercalc.RetrieveWeather(weathercalc.LocateCity(lat, long))
	fmt.Printf("API Response as struct %+v\n", weatherInfo)

	http.HandleFunc("/", jsonResponse)
	http.ListenAndServe(":3000", nil)

}

func jsonResponse(w http.ResponseWriter, r *http.Request) {
	js, err := json.Marshal(weatherInfo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
