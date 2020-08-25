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
// Connects to *Body URL, with *Body Token as a task
//
// @return: Response Status, error
func (auth *Body) Connect() (string, error) {
	// Get response from Body.URL
	response, err := http.Get(auth.URL + auth.Token)
	if err != nil {
		return response.Status, err
	}
	defer response.Body.Close()
	return response.Status, nil
}
