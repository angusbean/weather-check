package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/angusbean/weather-check/config"
	"github.com/angusbean/weather-check/handlers"
)

const portNumber = ":3000"

var app config.AppConfig
var infoLog *log.Logger
var errorLog *log.Logger

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}

	toPrint := fmt.Sprintf("Staring application on port %s", portNumber)
	fmt.Println(toPrint)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func run() error {
	// change this to true when in production
	app.InProduction = false

	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.InfoLog = infoLog

	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = errorLog

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	return nil
}

//weatherInfo = weathercalc.RetrieveWeather(weathercalc.LocateCity(lat, long))
//fmt.Printf("API Response as struct %+v\n", weatherInfo)
