package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/angusbean/weather-check/auth"
	"github.com/angusbean/weather-check/config"
	"github.com/angusbean/weather-check/models"
	weathercalc "github.com/angusbean/weather-check/weather-calc"
	"github.com/dgrijalva/jwt-go"
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

//Login validates the user details and returns JWT Access and Refresh Tokens
func (m *Repository) Login(w http.ResponseWriter, r *http.Request) {

	var validUser = models.User{
		ID:       1,
		Username: "username",
		Password: "password",
	}

	var user models.User

	//Read json file into memory with limits on json file size
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		log.Print(err)
	}

	//Check for errors in body of json file
	if err := r.Body.Close(); err != nil {
		log.Print(err)
	}

	//Unmarshal json into location struct, checking for errors
	if err := json.Unmarshal(body, &user); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		if err := json.NewEncoder(w).Encode(err); err != nil {
			log.Print(err)
		}
	}

	//compare the user from the request, with the one defined:
	if user.Username != validUser.Username || user.Password != validUser.Password {
		var errReponse models.ErrReponse
		errReponse.Error = "invalid user"
		errReponse.Description = "incorrect username or password provided"
		errReponse.Code = 401

		jsonErrReponse, err := json.Marshal(errReponse)
		if err != nil {
			log.Print(err)
		}

		//Set response headers and write JSON as reponse
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(401)
		w.Write(jsonErrReponse)
		return
	}

	ts, err := auth.CreateToken(user.ID)
	if err != nil {
		log.Print(err)
		return
	}

	//store any errors with auth in redis
	saveErr := auth.CreateAuth(user.ID, ts)
	if saveErr != nil {
		log.Print(saveErr)
	}

	tokens := map[string]string{
		"access_token":  ts.AccessToken,
		"refresh_token": ts.RefreshToken,
	}

	jsonTokens, err := json.Marshal(tokens)
	if err != nil {
		log.Print(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write([]byte(jsonTokens))
	return

}

//RequestWeather accepts lat and long as JSON input and returns JSON weather information for closest city location
func (m *Repository) RequestWeather(w http.ResponseWriter, r *http.Request) {
	var location models.Coord

	//Read json file into memory with limits on json file size
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		log.Print(err)
	}

	//Check for errors in body of json file
	if err := r.Body.Close(); err != nil {
		log.Print(err)
	}

	//Unmarshal json into location struct, checking for errors
	if err := json.Unmarshal(body, &location); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		if err := json.NewEncoder(w).Encode(err); err != nil {
			log.Print(err)
		}
	}

	//Validate the JWT token, returning JSON error if not valid
	_, err = auth.ExtractTokenMetadata(r)
	if err != nil {
		var errReponse models.ErrReponse
		errReponse.Error = "authorisation error"
		errReponse.Description = err.Error()
		errReponse.Code = 401

		jsonErrReponse, err := json.Marshal(errReponse)
		if err != nil {
			log.Print(err)
		}

		//Set response headers and write JSON as response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(401)
		w.Write(jsonErrReponse)
		return
	}

	//Check to see if values are valid lats and longs
	if location.Lat < -90 || location.Lat > 90 || location.Long < -180 || location.Long > 180 {
		var errReponse models.ErrReponse
		errReponse.Error = "useage error"
		errReponse.Description = "requires valid latitude and longitude values"
		errReponse.Code = 406

		jsonErrReponse, err := json.Marshal(errReponse)
		if err != nil {
			log.Print(err)
		}

		//Set response headers and write JSON as reponse
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(406)
		w.Write(jsonErrReponse)
		return
	}

	//Check to see if coords were populated correctly, 0 values check for non correct types
	if location.Lat == 0 || location.Long == 0 {
		var errReponse models.ErrReponse
		errReponse.Error = "useage error"
		errReponse.Description = "requires lat and long fields with float64 as values"
		errReponse.Code = 400

		jsonErrReponse, err := json.Marshal(errReponse)
		if err != nil {
			log.Print(err)
		}

		//Set response headers and write JSON as response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		w.Write(jsonErrReponse)
		return
	}

	//Using the lat and long provided return the closest city ID number (OpenWeatherMaps)
	closestCityID := weathercalc.LocateCity(location.Lat, location.Long, m.App.Cities)

	//Create weather object based on location
	weather := weathercalc.RetrieveWeather(closestCityID)

	//Marshal new weather object into JSON
	jsonWeather, err := json.MarshalIndent(weather, "", "     ")
	if err != nil {
		log.Print(err)
	}

	//Set response headers and write JSON as reponse
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(jsonWeather)
}

//Logout validates the user details and returns JWT Access and Refresh Tokens
func (m *Repository) Logout(w http.ResponseWriter, r *http.Request) {
	au, err := auth.ExtractTokenMetadata(r)
	//check if token valid
	if err != nil {
		var errReponse models.ErrReponse
		errReponse.Error = "authorisation error"
		errReponse.Description = err.Error()
		errReponse.Code = 401

		jsonErrReponse, err := json.Marshal(errReponse)
		if err != nil {
			log.Print(err)
		}

		//Set response headers and write JSON as response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(401)
		w.Write(jsonErrReponse)
		return
	}

	//delete Access token and check if delete successfully
	deleted, delErr := auth.DeleteAuth(au.AccessUuid)
	if delErr != nil || deleted == 0 {
		var errReponse models.ErrReponse
		errReponse.Error = "authorisation error"
		errReponse.Description = delErr.Error()
		errReponse.Code = 401

		jsonErrReponse, err := json.Marshal(errReponse)
		if err != nil {
			log.Print(err)
		}

		//Set response headers and write JSON as response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(401)
		w.Write(jsonErrReponse)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write([]byte("Successfully logged out"))
	return
}

//RefreshToken allows user to refresh their Access Token without needing to re-login
func (m *Repository) RefreshToken(w http.ResponseWriter, r *http.Request) {
	mapToken := map[string]string{}

	//Read json file into memory with limits on json file size
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		log.Print(err)
	}

	//Check for errors in body of json file
	if err := r.Body.Close(); err != nil {
		log.Print(err)
	}

	//Unmarshal json into mapToken
	if err := json.Unmarshal(body, &mapToken); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		if err := json.NewEncoder(w).Encode(err); err != nil {
			log.Print(err)
		}
	}
	refreshToken := mapToken["refresh_token"]

	//verify the token
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("REFRESH_SECRET")), nil
	})

	//if there is an error, the token must have expired
	if err != nil {
		var errReponse models.ErrReponse
		errReponse.Error = "authorisation error"
		errReponse.Description = err.Error()
		errReponse.Code = 401

		jsonErrReponse, err := json.Marshal(errReponse)
		if err != nil {
			log.Print(err)
		}

		//Set response headers and write JSON as response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(401)
		w.Write(jsonErrReponse)
		return
	}

	//check if the token is valid
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		var errReponse models.ErrReponse
		errReponse.Error = "authorisation error"
		errReponse.Description = err.Error()
		errReponse.Code = 401

		jsonErrReponse, err := json.Marshal(errReponse)
		if err != nil {
			log.Print(err)
		}

		//Set response headers and write JSON as response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(401)
		w.Write(jsonErrReponse)
		return
	}

	//Since token is valid, get the uuid:
	claims, ok := token.Claims.(jwt.MapClaims) //the token claims should conform to MapClaims
	if ok && token.Valid {
		refreshUuid, ok := claims["refresh_uuid"].(string) //convert the interface to string
		if !ok {
			var errReponse models.ErrReponse
			errReponse.Error = "authorisation error"
			errReponse.Description = err.Error()
			errReponse.Code = 401

			jsonErrReponse, err := json.Marshal(errReponse)
			if err != nil {
				log.Print(err)
			}

			//Set response headers and write JSON as response
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(401)
			w.Write(jsonErrReponse)
			return
		}

		userId, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
		if err != nil {
			var errReponse models.ErrReponse
			errReponse.Error = "authorisation error"
			errReponse.Description = err.Error()
			errReponse.Code = 422

			jsonErrReponse, err := json.Marshal(errReponse)
			if err != nil {
				log.Print(err)
			}

			//Set response headers and write JSON as response
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(422)
			w.Write(jsonErrReponse)
			return
		}

		//Delete the previous Refresh Token
		deleted, delErr := auth.DeleteAuth(refreshUuid)
		if delErr != nil || deleted == 0 { //if any goes wrong
			var errReponse models.ErrReponse
			errReponse.Error = "authorisation error"
			errReponse.Description = err.Error()
			errReponse.Code = 401

			jsonErrReponse, err := json.Marshal(errReponse)
			if err != nil {
				log.Print(err)
			}

			//Set response headers and write JSON as response
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(401)
			w.Write(jsonErrReponse)
			return
		}

		//Create new pairs of refresh and access tokens
		ts, createErr := auth.CreateToken(userId)
		if createErr != nil {
			var errReponse models.ErrReponse
			errReponse.Error = "authorisation error"
			errReponse.Description = err.Error()
			errReponse.Code = 403

			jsonErrReponse, err := json.Marshal(errReponse)
			if err != nil {
				log.Print(err)
			}

			//Set response headers and write JSON as response
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(403)
			w.Write(jsonErrReponse)
			return
		}

		//save the tokens metadata to redis
		saveErr := auth.CreateAuth(userId, ts)
		if saveErr != nil {
			var errReponse models.ErrReponse
			errReponse.Error = "authorisation error"
			errReponse.Description = err.Error()
			errReponse.Code = 403

			jsonErrReponse, err := json.Marshal(errReponse)
			if err != nil {
				log.Print(err)
			}

			//Set response headers and write JSON as response
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(403)
			w.Write(jsonErrReponse)
			return
		}

		tokens := map[string]string{
			"access_token":  ts.AccessToken,
			"refresh_token": ts.RefreshToken,
		}

		jsonTokens, err := json.Marshal(tokens)
		if err != nil {
			log.Print(err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write([]byte(jsonTokens))
		return
	} else {
		var errReponse models.ErrReponse
		errReponse.Error = "authorisation error"
		errReponse.Description = err.Error()
		errReponse.Code = 401

		jsonErrReponse, err := json.Marshal(errReponse)
		if err != nil {
			log.Print(err)
		}

		//Set response headers and write JSON as response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(401)
		w.Write(jsonErrReponse)
		return
	}
}
