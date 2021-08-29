package models

//AccessDetails is used by the auth package to validate JWTs
type AccessDetails struct {
	AccessUuid string
	UserId     uint64
}
