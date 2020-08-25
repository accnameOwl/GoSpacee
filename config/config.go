package config

import (
	"encoding/json"
	"flag"
	"log"
	"os"
)

// Conf ...
// datacollection from config.json,
//
// use CatchConfigurations to fetch data
type Conf struct {
	Auth   string
	Remote bool
	Debug  bool
}

// CatchConfigurations ...
// Returns Configuration struct with data pulled from config.json and stores it to
// /etc/config.conf
func CatchConfigurations(fileDir string) Conf {
	c := flag.String("c", fileDir, "Specify the configuration file")
	flag.Parse()
	file, err := os.Open(*c)
	if err != nil {
		log.Fatal("Can't open config file: ", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	config := Conf{}
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatal("Can't decode config JSON: ", err)
	}
	return config
}
