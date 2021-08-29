package auth

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/angusbean/weather-check/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-redis/redis"
	"github.com/twinj/uuid"
)

//CreateToken returns token details assosicated with a unique user
func CreateToken(userid uint64) (*models.TokenDetails, error) {
	td := &models.TokenDetails{}

	//set expiry of Access Token
	td.AtExpires = time.Now().Add(time.Minute * 5).Unix()
	td.AccessUuid = uuid.NewV4().String()

	//set expiry of Refresh Token
	td.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	td.RefreshUuid = uuid.NewV4().String()

	var err error

	//Creating Access Token
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["access_uuid"] = td.AccessUuid
	atClaims["user_id"] = userid
	atClaims["exp"] = td.AtExpires
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	td.AccessToken, err = at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return nil, err
	}

	//Creating Refresh Token
	rtClaims := jwt.MapClaims{}
	rtClaims["refresh_uuid"] = td.RefreshUuid
	rtClaims["user_id"] = userid
	rtClaims["exp"] = td.RtExpires
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	td.RefreshToken, err = rt.SignedString([]byte(os.Getenv("REFRESH_SECRET")))
	if err != nil {
		return nil, err
	}
	return td, nil
}

//CreateAuth returns any error information from the authentication process
func CreateAuth(userid uint64, td *models.TokenDetails) error {
	//
	dsn := os.Getenv("REDIS_DSN")
	client := redis.NewClient(&redis.Options{
		Addr: dsn, //redis port
	})

	_, err := client.Ping().Result()
	if err != nil {
		log.Fatal(err)
	}

	at := time.Unix(td.AtExpires, 0) //converting Unix to UTC(to Time object)
	rt := time.Unix(td.RtExpires, 0)
	now := time.Now()

	errAccess := client.Set(td.AccessUuid, strconv.Itoa(int(userid)), at.Sub(now)).Err()
	if errAccess != nil {
		return errAccess
	}
	errRefresh := client.Set(td.RefreshUuid, strconv.Itoa(int(userid)), rt.Sub(now)).Err()
	if errRefresh != nil {
		return errRefresh
	}
	return nil
}
