package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/angusbean/weather-check/config"
	"github.com/angusbean/weather-check/handlers"
	"github.com/angusbean/weather-check/models"
	"github.com/angusbean/weather-check/routes"
	weathercalc "github.com/angusbean/weather-check/weather-calc"
)

const portNumber = ":3000"

var app config.AppConfig
var infoLog *log.Logger
var errorLog *log.Logger
var cityList models.Cities

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}

	toPrint := fmt.Sprintf("Starting application on port %s", portNumber)
	fmt.Println(toPrint)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes.Routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func run() error {
	//InProduction should change this to true when in production
	app.InProduction = false

	//infoLog prints to terminal (update in Production)
	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.InfoLog = infoLog

	//errorLog prints to terminal (update to file in Production)
	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = errorLog

	//cityList loads the JSON file of city information into memory
	jFile, err := os.Open("weather-calc/openweather-info/city.list.json")
	if err != nil {
		log.Fatal(err)
	}
	defer jFile.Close()
	cityList = weathercalc.LoadCityList(jFile)
	app.Cities = cityList

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	return nil
}
