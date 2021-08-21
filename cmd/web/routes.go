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

	mux.Use(middleware.Recoverer)
	mux.Use(middleware.Logger)

	mux.Post("/", handlers.Repo.RequestWeather)

	return mux
}
