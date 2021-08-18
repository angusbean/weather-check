package models

//Weather provides important information about weather data
type Weather struct {
	LocationName string `json:"name"`
	Coord        struct {
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
	Code int `json:"cod"`
}
