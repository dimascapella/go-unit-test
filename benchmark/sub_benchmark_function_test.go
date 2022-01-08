package benchmark

import "testing"

func BenchmarkSub(b *testing.B) {
	b.Run("Bwang", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Bwang")
		}
	})
	b.Run("Rola", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Rola")
		}
	})
}
