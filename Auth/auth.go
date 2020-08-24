package auth

import (
	"net/http"
)

// Auth ...
//
// body of an authorization task
type Auth struct {
	URL   string
	Token string
}

// ConnectToURL ...
// Connects to *Body URL, with *Body Token as a task
//
// @return: Response Status, error
func (auth *Auth) ConnectToURL() (string, error) {
	// Get response from Body.URL
	response, err := http.Get(auth.URL + auth.Token)
	if err != nil {
		return response.Status, err
	}
	defer response.Body.Close()
	return response.Status, nil
}
