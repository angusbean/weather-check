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
func (m *Repository) GetWeather(w http.ResponseWriter, r *http.Request) {
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
		w.WriteHeader(422)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			fmt.Println(err)
		}
	}

	if location.Lat == 0 && location.Long == 0 {
		var errReponse models.ErrReponse
		errReponse.Error = "useage error"
		errReponse.Description = "requires fields lat and long with float64 values"
		errReponse.Code = 204

		jsonErrReponse, err := json.Marshal(errReponse)
		if err != nil {
			fmt.Println(err)
		}

		//Set response headers and write JSON as reponse
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonErrReponse)
		return
	}

	//Create weather object based on location
	weather := weathercalc.RetrieveWeather(weathercalc.LocateCity(location.Lat, location.Long))

	//Marshal new weather object into JSON
	jsonWeather, err := json.MarshalIndent(weather, "", "     ")
	if err != nil {
		fmt.Println(err)
	}

	//Set response headers and write JSON as reponse
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonWeather)
}
