package benchmark

import "testing"

func BenchmarkTable(b *testing.B) {
	user := []struct {
		Name    string
		request string
	}{
		{
			Name:    "Dimas",
			request: "Dimas",
		},
		{
			Name:    "Eka",
			request: "Eka",
		},
	}

	for _, benchmark := range user {
		b.Run(benchmark.Name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.request)
			}
		})
	}
}
