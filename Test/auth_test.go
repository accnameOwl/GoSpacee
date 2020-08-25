package test

import (
	"github.com/accnameowl/GoSpacee/auth"
)

func ServerConnection(a *auth.Body) <-chan string {
	taskConnectionTest := make(chan string)
	go func() {
		defer close(taskConnectionTest)
		taskConnectionTest <- a.Connect()
	}()
	return taskConnectionTest
}
