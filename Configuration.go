package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"
)

// Configuration ...
// holds the value of every json type in config.json
type Configuration struct {
	Auth string
}

// CatchConfigurations ...
// Returns Configuration struct with data pulled from config.json and stores it to
// /etc/config.conf
func CatchConfigurations(fileDir string) Configuration {
	c := flag.String("c", "/etc/config.conf", "config.json")
	flag.Parse()
	file, err := os.Open(*c)
	if err != nil {
		log.Fatal("Can't open config file: ", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	config := Configuration{}
	err = decoder.Decone(&Config)
	if err != nil {
		log.Fatal("Can't decode config JSON: ", err)
	}
	return config
}
