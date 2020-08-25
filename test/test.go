package test

import (
	"github.com/accnameowl/GoSpacee/auth"
)

// ServerConnection ...
// async function, testing connection with URL and Token
//
// expecting: await := <-ServerConnection()
func ServerConnection(a *auth.Body) <-chan string {
	taskConnectionTest := make(chan string)
	go func() {
		defer close(taskConnectionTest)
		taskConnectionTest <- a.Connect()
	}()
	return taskConnectionTest
}
