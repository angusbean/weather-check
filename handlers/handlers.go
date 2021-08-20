package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/angusbean/weather-check/config"
	"github.com/angusbean/weather-check/models"
	weathercalc "github.com/angusbean/weather-check/weather-calc"
)

//Repo the repository used by the handlers
var Repo *Repository

//Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

//PostLocation accepts JSON input and returns JSON weather information for that location
func (m *Repository) GetWeather(w http.ResponseWriter, r *http.Request) {
	var location models.LatLong
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		fmt.Println(err)
	}

	if err := r.Body.Close(); err != nil {
		fmt.Println(err)
	}

	if err := json.Unmarshal(body, &location); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			fmt.Println(err)
		}
	}

	fmt.Printf("Location as struct: %+v\n", location)

	weather := weathercalc.RetrieveWeather(weathercalc.LocateCity(location.Lat, location.Long))

	fmt.Println(weather)

	jsonWeather, err := json.MarshalIndent(weather, "", "     ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(jsonWeather))

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonWeather)
}
