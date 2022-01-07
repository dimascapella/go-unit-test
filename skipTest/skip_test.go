package skipTest

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSkip(t *testing.T) {
	if runtime.GOARCH == "amd64" {
		t.Skip("Cannot Run on Windows")
	}

	result := HelloWorld("Bwang")
	assert.Equal(t, "Hello Bwang", result, "Must Be Hello Bwang")
}

func TestMain(m *testing.M) {
	fmt.Println("Before Test")
	m.Run()
	fmt.Println("After Test")
}
