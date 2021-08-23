package weathercalc

import (
	"testing"
)

func TestLoadCityList(t *testing.T) {
	//TODO
}

var LocationTests = []struct {
	name               string
	lat                float64
	long               float64
	expectedReturnedID int
}{
	{"sydney-test", -33.86, 151.29, 251},
}

func TestLocateCity(t *testing.T) {
	cityList := LoadCityList()
	for _, e := range LocationTests {
		result := LocateCity(e.lat, e.long, cityList)
		if result != e.expectedReturnedID {
			t.Errorf("Doesnt Match")
		}
	}
}

var TheWeatherTests = []struct {
	name               string
	lat                float64
	long               float64
	expectedReturnedID int
}{
	{"sydney-test", -33.86, 151.29, 251},
}

var RetrieveWeatherTests = []struct {
	ID                int
	ExpectedErrorCode int
}{
	{833, 200},  //Valid
	{2960, 200}, //Valid
	{800, 0},    //ID Does not exsist
}

func TestRetrieveWeather(t *testing.T) {
	for _, e := range RetrieveWeatherTests {
		result := RetrieveWeather(e.ID)
		if result.Code != e.ExpectedErrorCode {
			t.Errorf("for %d, expected %d but got %d", e.ID, e.ExpectedErrorCode, result.Code)
		}
	}
}
