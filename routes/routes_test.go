package routes

import (
	"fmt"
	"testing"

	"github.com/angusbean/weather-check/config"
	"github.com/go-chi/chi"
)

func TestRoutes(t *testing.T) {
	var app config.AppConfig

	mux := Routes(&app)

	switch v := mux.(type) {
	case *chi.Mux:
		//do nothing test passed
	default:
		t.Error((fmt.Sprintf("type is not *chi.Mux, type is %T", v)))
	}

}
