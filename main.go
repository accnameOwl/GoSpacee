package main

import (
	"fmt"

	"github.com/accnameowl/Calendb/auth"
)

func main() {

	Config := GetConfigurations("config.json")
	fmt.Println(Config.Auth)
	authTest := auth.Body{
		URL:   "https://api.nasa.gov/planetary/apod",
		Token: "?api_key=QeHSqI7jLoVAvghkd0SF0ZJ03v1XT2XlfY4dpBLT",
	}

	res, err := authTest.ConnectToURL()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Successfully connected!\nURL: %s\nAuth token: %s\nResponse Status: %s", authTest.URL, authTest.Token, res)
}
