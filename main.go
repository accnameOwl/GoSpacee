package main

import (
	"fmt"

	"github.com/accnameowl/GoSpacee/auth"
	"github.com/accnameowl/GoSpacee/config"
	"github.com/accnameowl/GoSpacee/insight"
)

// Config ...
// Catches information from 'config.json'
var Config config.Conf

func init() {
	Config = config.CatchConfigurations("./config/config.json")
}

func main() {

	Authorization := auth.Body{
		URL:   "https://api.nasa.gov/planetary/apod",
		Token: "?api_key=QeHSqI7jLoVAvghkd0SF0ZJ03v1XT2XlfY4dpBLT",
	}

	status, err := Authorization.Connect()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Connected to: %s\nToken: %s\n%s", Authorization.URL, Authorization.Token, status)

	// * new Mars weather station
	InSight := insight.New("1.0", "json", "DEMO_KEY")
	err = InSight.Get()
	if err != nil {
		panic(err.Error())
	}
}
