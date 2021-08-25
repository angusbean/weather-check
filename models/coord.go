package models

//LatLong provides struct for Lat and Long values
type Coord struct {
	Lat  float64 `json:"lat"`
	Long float64 `json:"long"`
}
