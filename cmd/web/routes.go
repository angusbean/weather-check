package main

import (
	"net/http"

	"github.com/angusbean/weather-check/config"
	"github.com/angusbean/weather-check/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	//Use Chi middleware for Logging and Applcaiton recovery
	mux.Use(middleware.Recoverer)
	mux.Use(middleware.Logger)

	//Post Requests
	mux.Post("/", handlers.Repo.RequestWeather)

	return mux
}
