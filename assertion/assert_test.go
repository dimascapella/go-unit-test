package assertion

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Bwang")
	assert.Equal(t, "Hello Bwang", result, "Must Be Hello Bwang")
	assert.NotEqual(t, "Hi Bwang", result, "Not Same")
	fmt.Println("Test With Assert Done")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Bwang")
	require.Equal(t, "Hi Bwang", result, "Must Be Hello Bwang")
}

// assert = Running test and call Fail(), the program still running while meet the errors
// require = Running test and call FailNow(), programm will stop while meet the errors
