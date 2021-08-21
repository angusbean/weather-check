package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type postData struct {
	key   string
	value float64
}

var theTests = []struct {
	name               string
	url                string
	method             string
	params             []postData
	expectedStatusCode int
}{
	{"RequestWeather-valid", "/", "POST", []postData{
		{key: "lat", value: 151.35},
		{key: "long", value: 35.25},
	}, http.StatusOK},
	{"RequestWeather-not-valid", "/", "POST", []postData{
		{key: "lat", value: 10000.00},
		{key: "long", value: 10000.00},
	}, http.StatusNotAcceptable},
	{"RequestWeather-0-values", "/", "POST", []postData{
		{key: "lat", value: 0},
		{key: "long", value: 0},
	}, http.StatusNotAcceptable},
	{"RequestWeather-null", "/", "POST", []postData{
		{key: "", value: 0},
		{key: "", value: 0},
	}, http.StatusNotAcceptable},
}

func TestHandlers(t *testing.T) {
	routes := getTestRoutes()
	ts := httptest.NewTLSServer(routes)
	defer ts.Close()

	for _, e := range theTests {
		if e.method == "GET" { //Checks GET Methods
			resp, err := ts.Client().Get(ts.URL + e.url)
			if err != nil {
				t.Log()
				t.Fatal(err)
			}

			if resp.StatusCode != e.expectedStatusCode {
				t.Errorf("for %s, expected %d but got %d", e.name, e.expectedStatusCode, resp.StatusCode)
			}
		}
	}
}
