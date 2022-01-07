package helper

import "testing"

func TestHello(t *testing.T) {
	res := HelloWorld("Dimas")
	if res != "Hello Dimas" {
		panic("Result Not Match")
	}
}

// go test = run all unit test on current folder
// go test -v = run all unit test with detail
// go test -v -run FunctionName = run unit test with spesific function
// go test ./... = run all unit test from parent folder
