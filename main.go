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

	Configuration := GetConfig("main")

	// * Flags
	flag.BoolVar(&Configuration.Remote, "remote", false, "Download data locally, instead of using remote, true or false")
	flag.BoolVar(&Configuration.Debug, "debug", false, "Debug mode")
	flag.Parse()

	var Authorization = &auth.TestBody
	// if debug
	if Configuration.Debug {
		fmt.Println("Starting app in Debug mode")
		time.Sleep(time.Second * 1)
		//await server response
		fmt.Println("Connecting...")
		await := <-test.ServerConnection(Authorization)
		fmt.Printf("Connected to: %s\nToken: %s\n%s\n\n", Authorization.URL, Authorization.Token, await)
	}

	// * new Mars weather station
	InSight := insight.New("1.0", "json", "DEMO_KEY")
	err := InSight.Get(Configuration.Remote)
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
