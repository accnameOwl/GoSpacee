package test

import (
	"fmt"
	"net/http"
	"github.com/accnameowl/gospacee/auth"
)

// Run_Test ...
// Runs a test on connection.
func (a *Auth) Connection() res string {
	response, err := http.Get(a.URL + a.Token)
	if err != nil {
		return err
	}

	defer response.Body.Close()
	res = "Successfully connected!\nURL: " + authTest.URL + "\nAuth token: " + authTest.Token + "\nResponse Status: " + response.Status
}
