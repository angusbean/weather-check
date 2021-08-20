package main

import (
	"net/http"

	"github.com/angusbean/weather-check/config"
	"github.com/angusbean/weather-check/handlers"
	"github.com/go-chi/chi"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Post("/", handlers.Repo.PostLocation)

	return mux
}
