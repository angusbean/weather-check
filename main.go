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
	"github.com/go-redis/redis"
	"github.com/joho/godotenv"
)

const portNumber = ":3000"

var app config.AppConfig
var infoLog *log.Logger
var errorLog *log.Logger
var cityList models.Cities
var redisClient *redis.Client

func init() {
	//Initializing redis
	dsn := os.Getenv("REDIS_DSN")
	if len(dsn) == 0 {
		dsn = "localhost:6379"
	}
	redisClient = redis.NewClient(&redis.Options{
		Addr: dsn, //redis port
	})
	_, err := redisClient.Ping().Result()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("application running on port :3000")

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

	//Set the applicaiton environment variables from .env file
	err = godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	return nil
}
