package routes

import (
	"net/http"

	"github.com/angusbean/weather-check/auth"
	"github.com/angusbean/weather-check/config"
	"github.com/angusbean/weather-check/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func Routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	//Chi middleware for Logging and Applcaiton recovery
	mux.Use(middleware.Recoverer)
	mux.Use(middleware.Logger)

	// Protect routes by JWT
	mux.Group(func(mux chi.Router) {
		mux.Use(auth.TokenAuthMiddleware)

		mux.Post("/request-weather", handlers.Repo.RequestWeather)
		mux.Post("/token/refresh", handlers.Repo.RefreshToken)
		mux.Post("/logout", handlers.Repo.Logout)
	})

	//Public routes
	mux.Post("/login", handlers.Repo.Login)

	return mux
}
