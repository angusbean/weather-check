package models

//City provides struct for city location information
type City struct {
	ID    int `json:"id"`
	Coord `json:"coord"`
}
