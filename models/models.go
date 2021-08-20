package models

//Weather provides important information about weather data
type Weather struct {
	LocationName string `json:"name"`
	Sys          struct {
		Country string `json:"country"`
	} `json:"sys"`
	Coord struct {
		Lat  float32 `json:"lat"`
		Long float32 `json:"lon"`
	} `json:"coord"`
	Weather []struct {
		Weather     string `json:"main"`
		Description string `json:"description"`
	} `json:"weather"`
	Temperature struct {
		Temperature    float32 `json:"temp"`
		MinTemperature float32 `json:"temp_min"`
		MaxTemperature float32 `json:"temp_max"`
	} `json:"main"`
	Wind struct {
		Speed float32 `json:"speed"`
		Deg   float32 `json:"deg"`
		Gust  float32 `json:"gust"`
	} `json:"wind"`
	Code int `json:"cod"`
}

//CityList provides the struct for the list
type CityList struct {
	CityList []City `json:"citylist"`
}

//City provides struct for city location information
type City struct {
	ID    int `json:"id"`
	Coord struct {
		Lat  float32 `json:"lat"`
		Long float32 `json:"lon"`
	} `json:"coord"`
}

//LatLong provides struct for Lat and Long values
type LatLong struct {
	Lat  float64 `json:"lat"`
	Long float64 `json:"long"`
}

//ErrReponse provides struct for retrun JSON error message
type ErrReponse struct {
	Error       string
	Description string
	Code        int
}
