package auth

import (
	"net/http"
)

// TestBody ...
// TestBody body, containing connection URL and token.
// used for testing connection with API token
var TestBody = Body{
	URL:   "https://api.nasa.gov/planetary/apod",
	Token: "QeHSqI7jLoVAvghkd0SF0ZJ03v1XT2XlfY4dpBLT",
}

// Body ...
//
// body of an authorization task
type Body struct {
	URL   string
	Token string
}

// Connect ...
// Connects to *Body URL, with auth *Body.token
//
// @return: Response Status
func (auth *Body) Connect() string {
	// Get response from Body.URL
	response, err := http.Get(auth.URL + "?api_key=" + auth.Token)
	if err != nil {
		panic(err.Error())
	}
	defer response.Body.Close()
	return response.Status
}
