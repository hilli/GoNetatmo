package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-resty/resty"
	"github.com/joho/godotenv"
)

// AuthToken comment
type AuthToken struct {
	AccessToken  string   `json:"access_token"`
	RefreshToken string   `json:"refresh_token"`
	Scope        []string `json:"scope"`
	ExpiresIn    int      `json:"expires_in"`
	ExpireIn     int      `json:"expire_in"`
}

// StationData comment
type StationData struct {
	Body Body `json:"body"`
}

// Body comment
type Body struct {
	Devices []Devices `json:"devices"`
}

// Devices comment
type Devices struct {
	StationName string `json:"station_name"`
}

func authenticate(client *resty.Client) string {
	client.SetFormData(map[string]string{
		"grant_type":    "password",
		"scope":         "read_station",
		"client_id":     os.Getenv("CLIENT_ID"),
		"client_secret": os.Getenv("CLIENT_SECRET"),
		"username":      os.Getenv("NETATMO_USERNAME"),
		"password":      os.Getenv("NETATMO_PASSWD")})

	response, error := client.R().SetResult(AuthToken{}).Post("/oauth2/token")

	if error != nil {
		panic(error)
	}

	return response.Result().(*AuthToken).AccessToken
}

func stationData(client *resty.Client, token string) string {
	client.SetFormData(map[string]string{
		"access_token": token})
	response, error := client.R().SetResult(StationData{}).Post("/api/getstationsdata")

	if error != nil {
		panic(error)
	}
	return response.Result().(*StationData).Body.Devices[0].StationName

}

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	client := resty.New()
	client.SetHeader("Content-Type", "application/x-www-form-urlencoded")
	client.SetHostURL("https://api.netatmo.com")

	token := authenticate(client)

	fmt.Println(stationData(client, token))
}
