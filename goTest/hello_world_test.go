package goTest

import (
	"fmt"
	"testing"
)

func TestHello(t *testing.T) {
	res := HelloWorld("Dimas")
	if res != "Hello Dimas" {
		t.Fatal("Must Be `Hello Bwang`")
	}
	fmt.Println("Bleh Bleh Bleh")
}

// Command for running test
// go test = run all unit test on current folder
// go test -v = run all unit test with detail
// go test -v -run FunctionName = run unit test with spesific function
// go test ./... = run all unit test from parent folder

// Handle Error Test
// Fail() handle error but still running until finish the code
// FailNow() handle error but stop when running code while error
// Error() handle error and call method Fail() with details error
// Fatal() handler error and call method FailNow() with details error
