package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-resty/resty"
	"github.com/joho/godotenv"
)

func authenticate(client *resty.Client) {
	client.SetFormData(map[string]string{
		"grant_type":    "password",
		"client_id":     os.Getenv("CLIENT_ID"),
		"client_secret": os.Getenv("CLIENT_SECRET"),
		"scope":         "read_station",
		"username":      os.Getenv("NETATMO_USERNAME"),
		"password":      os.Getenv("NETATMO_PASSWD")})
	response, error := client.R().Post("/oauth2/token")
	fmt.Println(error)
	fmt.Println(response)
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

	authenticate(client)
}
