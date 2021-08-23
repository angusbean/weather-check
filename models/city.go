package models

//City provides struct for city location information
type City struct {
	ID    int `json:"id"`
	Coord struct {
		Lat  float32 `json:"lat"`
		Long float32 `json:"lon"`
	} `json:"coord"`
}
