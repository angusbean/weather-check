package models

//Weather provides important information about weather data
type WeatherUpdate struct {
	LocationName string `json:"name"`
	Sys          struct {
		Country string `json:"country"`
		Sunrise int    `json:"sunrise"`
		Sunset  int    `json:"sunset"`
	} `json:"sys"`
	Coord   `json:"coord"`
	Weather []struct {
		Main        string `json:"main"`
		Description string `json:"description"`
	} `json:"weather"`
	Temperature struct {
		Temperature    float32 `json:"temp"`
		MinTemperature float32 `json:"temp_min"`
		MaxTemperature float32 `json:"temp_max"`
		Pressure       float32 `json:"pressure"`
		Humidity       float32 `json:"humidity"`
	} `json:"main"`
	Visiblity int `json:"visiblity"`
	Clouds    struct {
		All int `json:"all"`
	} `json:"clouds"`
	Wind struct {
		Speed float32 `json:"speed"`
		Deg   float32 `json:"deg"`
		Gust  float32 `json:"gust"`
	} `json:"wind"`
	Code int `json:"cod"`
}
