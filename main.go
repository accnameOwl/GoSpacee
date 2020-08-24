package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/accnameowl/GoSpacee/auth"
)

// Config ...
// Catches information from 'config.json'
var Config Configuration

func init() {
	Config = CatchConfigurations("./config/config.json")
	fmt.Println(Config.Auth)
}

func main() {

	authTest := auth.Body{
		URL:   "https://api.nasa.gov/planetary/apod",
		Token: "?api_key=QeHSqI7jLoVAvghkd0SF0ZJ03v1XT2XlfY4dpBLT",
	}

	RunTests(authTest)
	res := authTest.Test()
}

// Configuration ...
// holds the value of every json type in config.json
type Configuration struct {
	Auth string
}

// CatchConfigurations ...
// Returns Configuration struct with data pulled from config.json and stores it to
// /etc/config.conf
func CatchConfigurations(fileDir string) Configuration {
	c := flag.String("c", fileDir, "Specify the configuration file")
	flag.Parse()
	file, err := os.Open(*c)
	if err != nil {
		log.Fatal("Can't open config file: ", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	config := Configuration{}
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatal("Can't decode config JSON: ", err)
	}
	return config
}
