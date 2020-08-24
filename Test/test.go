package test

import "fmt"

// Test ...
type Test interface {
	Connection() string
}

// RunTests ...
// runs a series of tests of given interface
func RunTests(t *Test) {
	fmt.Println(t.Connection())
}
