package main

import (
	"flag"
	"fmt"

	"github.com/accnameowl/GoSpacee/auth"
	"github.com/accnameowl/GoSpacee/config"
	"github.com/accnameowl/GoSpacee/insight"
	"github.com/accnameowl/GoSpacee/test"
)

// Config ...
// Catches information from 'config.json'
var Config config.Conf

func init() {
	Config = config.CatchConfigurations("./config/config.json")
}

func main() {

	// * Flags
	flag.BoolVar(&Config.Remote, "no-rem", false, "false: Download data locally, true(defalt): Fetch remotely")
	flag.BoolVar(&Config.Debug, "debug", true, "Debug mode")
	Authorization := auth.Body{
		URL:   "https://api.nasa.gov/planetary/apod",
		Token: "QeHSqI7jLoVAvghkd0SF0ZJ03v1XT2XlfY4dpBLT",
	}

	// await server response
	await := <-test.ServerConnection(Authorization)

	if Config.Debug {
		fmt.Print("Connecting...")
		fmt.Printf("Connected to: %s\nToken: %s\n%s", Authorization.URL, Authorization.Token, status)
	}

	// * new Mars weather station
	InSight := insight.New("1.0", "json", "DEMO_KEY")
	err = InSight.Get()
	if err != nil {
		panic(err.Error())
	}
}

// arrayFlags ...
type arrayFlags []string

// mainFlags ...
// scraped on main()
var mainFlags arrayFlags

func (i *arrayFlags) Set(value string) error {
	*i = append(*i, value)
	return nil
}
