package bench

import (
	"runtime"
	"testing"

	dto "github.com/javiertelioz/clean_architecture/pkg/application/dto/hello"
	usecase "github.com/javiertelioz/clean_architecture/pkg/application/use_cases/hello"
)

func BenchmarkHelloUseCase(b *testing.B) {
	// Memory stats before
	var stats runtime.MemStats
	runtime.ReadMemStats(&stats)
	heapBefore := stats.HeapAlloc

	// Setup
	uc := usecase.NewHelloUseCase()
	input := &dto.HelloInput{
		Name: "Joe",
	}

	// Initial validation
	result, err := uc.Execute(input)
	if err != nil {
		b.Fatalf("Initial execution failed: %v", err)
	}
	if result == nil {
		b.Fatal("Expected non-nil result")
	}

	// Benchmark
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result, err := uc.Execute(input)
		if err != nil {
			b.Fatalf("Execution failed at iteration %d: %v", i, err)
		}
		if result == nil {
			b.Fatalf("Got nil result at iteration %d", i)
		}
	}

	// Memory stats after
	runtime.ReadMemStats(&stats)
	heapAfter := stats.HeapAlloc

	b.ReportMetric(float64(heapAfter-heapBefore)/float64(b.N), "heap-delta-bytes/op")
	b.ReportMetric(float64(stats.NumGC), "GC-cycles")
}
