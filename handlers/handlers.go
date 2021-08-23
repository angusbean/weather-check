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

//GetWeather accepts lat and long as JSON input and returns JSON weather information for closest city location
func (m *Repository) RequestWeather(w http.ResponseWriter, r *http.Request) {
	var location models.LatLong

	//Read json file into memory with limits on json file size
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		fmt.Println(err)
	}

	//Check for errors in body of json file
	if err := r.Body.Close(); err != nil {
		fmt.Println(err)
	}

	//Unmarshal json into location struct, checking for errors
	if err := json.Unmarshal(body, &location); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		//w.WriteHeader(422)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			fmt.Println(err)
		}
	}

	//Check to see if values are valid lats and longs
	if location.Lat < -90 || location.Lat > 90 || location.Long < -180 || location.Long > 180 {
		var errReponse models.ErrReponse
		errReponse.Error = "useage error"
		errReponse.Description = "requires valid latitude and longitude values"
		errReponse.Code = 406

		jsonErrReponse, err := json.Marshal(errReponse)
		if err != nil {
			fmt.Println(err)
		}

		//Set response headers and write JSON as reponse
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(406)
		w.Write(jsonErrReponse)
		return
	}

	//Check to see if coords were populated correctly
	if location.Lat == 0 || location.Long == 0 {
		var errReponse models.ErrReponse
		errReponse.Error = "useage error"
		errReponse.Description = "requires lat and long fields with float64 as values"
		errReponse.Code = 400

		jsonErrReponse, err := json.Marshal(errReponse)
		if err != nil {
			fmt.Println(err)
		}

		//Set response headers and write JSON as response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		w.Write(jsonErrReponse)
		return
	}

	//Create weather object based on location
	weather := weathercalc.RetrieveWeather(weathercalc.LocateCity(location.Lat, location.Long, m.App.CityList))

	//Marshal new weather object into JSON
	jsonWeather, err := json.MarshalIndent(weather, "", "     ")
	if err != nil {
		fmt.Println(err)
	}

	//Set response headers and write JSON as reponse
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(jsonWeather)
}
