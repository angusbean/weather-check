package routes

import (
	"net/http"

	"github.com/angusbean/weather-check/config"
	"github.com/angusbean/weather-check/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func Routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	//Use Chi middleware for Logging and Applcaiton recovery
	mux.Use(middleware.Recoverer)
	mux.Use(middleware.Logger)

	//Post Requests
	mux.Post("/request-weather", handlers.Repo.RequestWeather)
	mux.Post("/login", handlers.Repo.Login)

	return mux
}
