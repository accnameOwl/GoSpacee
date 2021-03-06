package main

import (
	"flag"
	"fmt"
	"github.com/accnameowl/GoSpacee/auth"
	"github.com/accnameowl/GoSpacee/insight"
	"github.com/accnameowl/GoSpacee/test"
	"time"
)

func main() {

	Configurations := GetConfig("main")

	// * Flags
	flag.BoolVar(&Configurations.Debug, "debug", false, "Debug mode")
	flag.BoolVar(&Configurations.SaveFetched, "save-fetched", true, "Save fetched data locally")
	flag.Parse()

	var Authorization = &auth.TestBody
	// if debug
	if Configurations.Debug {
		fmt.Println("Starting app in Debug mode")
		time.Sleep(time.Second * 1)
		//await server response
		fmt.Println("Connecting...")
		await := <-test.ServerConnection(Authorization)
		fmt.Printf("Connected to: %s\nToken: %s\n%s\n\n", Authorization.URL, Authorization.Token, await)
	}

	// * new Mars weather station
	InSight := insight.New("1.0", "json", Configurations.Auth)
	err := InSight.Get(Configurations.SaveFetched)
	if err != nil {
		panic(err.Error())
	}
}
