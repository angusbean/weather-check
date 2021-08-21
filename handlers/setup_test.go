package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/angusbean/weather-check/config"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

var app config.AppConfig

func getTestRoutes() http.Handler {
	// change this to true when in production
	app.InProduction = false

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.InfoLog = infoLog

	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = errorLog

	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)

	mux.Post("/", Repo.RequestWeather)

	return mux
}
