package weathercalc

import (
	"log"
	"os"
	"testing"
)

func TestLoadCityList(t *testing.T) {
	jFile, err := os.Open("openweather-info/city.list.json")
	if err != nil {
		t.Error(err)
	}
	defer jFile.Close()
}

var LocationTests = []struct {
	name               string
	lat                float64
	long               float64
	expectedReturnedID int
}{
	{"Protaras", 35.012501, 34.058331, 18918},
	{"Judaydah", 15.07512, 45.299622, 30616},
}

func TestLocateCity(t *testing.T) {
	jFile, err := os.Open("openweather-info/city.list.json")
	if err != nil {
		log.Fatal(err)
	}
	defer jFile.Close()
	cityList := LoadCityList(jFile)
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
	{"Protaras", 35.012501, 34.058331, 18918},
	{"Judaydah", 15.07512, 45.299622, 30616},
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
