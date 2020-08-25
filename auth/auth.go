package auth

import (
	"net/http"
)

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
