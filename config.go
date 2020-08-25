package main

import (
	"fmt"
	"github.com/tkanos/gonfig"
)

// Configuration ...
// datacollection from config.json,
//
// use CatchConfigurations to fetch data
type Configuration struct {
	Auth        string
	Debug       bool
	SaveFetched bool
}

// GetConfig ...
func GetConfig(params ...string) Configuration {
	configuration := Configuration{}
	env := "main"
	if len(params) > 0 {
		env = params[0]
	}
	// fileName := fmt.Sprintf("./%config.json", env)
	fileName := fmt.Sprintf("./config/%s_config.json", env)
	gonfig.GetConf(fileName, &configuration)
	return configuration
}
