package benchmark

import "testing"

func BenchmarkHelloWorld(b *testing.B) {
	// TODO: Initialize
	for i := 0; i < b.N; i++ {
		HelloWorld("Bwang")
	}
}

// Command Running Benchmark
// go test -v -bench=. >> Running All Benchmark Code
// go test -v -run=NotMatchUnitTest -bench=. >> Running Benchmark without unit test
// go test -v -run=NotMatchUnitTest -bench=FunctionName >> Running Spesific Benchmark without unit test
// go test -v -bench. ./.. >> Running all benchmark and unit test (include root module)
// go test -v -bench=BenchmarkFunctionName/SubTestBenchmark >> Running Spesific Sub Benchmark
