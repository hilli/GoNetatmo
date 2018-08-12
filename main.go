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

func authenticate(client *resty.Client) string {
	client.SetFormData(map[string]string{
		"grant_type":    "password",
		"client_id":     os.Getenv("CLIENT_ID"),
		"client_secret": os.Getenv("CLIENT_SECRET"),
		"scope":         "read_station",
		"username":      os.Getenv("NETATMO_USERNAME"),
		"password":      os.Getenv("NETATMO_PASSWD")})

	response, error := client.R().SetResult(AuthToken{}).Post("/oauth2/token")

	if error != nil {
		panic(error)
	}

	return response.Result().(*AuthToken).AccessToken
}

func stationData(client *resty.Client, token string) {
	response, error := client.R().Post("/api/getstationdata")
	fmt.Println(error)
	fmt.Println(response)

}

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	client := resty.New()
	client.SetHeader("Content-Type", "application/x-www-form-urlencoded")
	client.SetHostURL("https://api.netatmo.com")

	fmt.Println(authenticate(client))
}
