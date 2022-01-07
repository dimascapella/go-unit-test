package subTest

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubTest(t *testing.T) {
	t.Run("Bwang", func(t *testing.T) {
		result := HelloWorld("Bwang")
		assert.Equal(t, "Hello Bwang", result, "Must Be Hello Bwang")
	})

	t.Run("Rodi", func(t *testing.T) {
		result := HelloWorld("Ba")
		assert.Equal(t, "Hello Rodi", result, "Must Be Hello Rodi")
	})
}

// go test -v -run=Func/subFunc = Running Spesific Sub Function
