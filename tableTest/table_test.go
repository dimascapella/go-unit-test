package tableTest

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHelloWorldTable(t *testing.T) {
	testCase := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Dimas",
			request:  "Dimas",
			expected: "Hello Dimas",
		},
		{
			name:     "Eka",
			request:  "Eka",
			expected: "Hello Eka",
		},
		{
			name:     "Zerg",
			request:  "Zerg",
			expected: "Hello Fuma",
		},
	}

	for _, test := range testCase {
		t.Run(test.name, func(t *testing.T) {
			res := HelloWorld(test.request)
			require.Equal(t, test.expected, res)
		})
	}
}
