package models

//ErrReponse provides struct for retrun JSON error message
type ErrReponse struct {
	Error       string
	Description string
	Code        int
}
