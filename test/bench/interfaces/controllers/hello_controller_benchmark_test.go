package controllers

import (
	"github.com/javiertelioz/clean_architecture/test/mocks/services"
	"net/http"
	"net/http/httptest"
	"runtime"
	"sync"
	"testing"

	"github.com/go-chi/chi/v5"

	usecase "github.com/javiertelioz/clean_architecture/pkg/application/use_cases/hello"
	controller "github.com/javiertelioz/clean_architecture/pkg/interfaces/controllers"
)

func BenchmarkHelloController(b *testing.B) {
	// Memory stats before
	var stats runtime.MemStats
	runtime.ReadMemStats(&stats)
	heapBefore := stats.HeapAlloc

	// Setup
	router := chi.NewRouter()
	loggerService := services.NewMockLoggerService()
	uc := usecase.NewHelloUseCase(loggerService)
	ctrl := controller.NewHelloController(uc, loggerService)

	router.Get("/api/v1/hello/{name}", ctrl.HelloHandler)

	// Create request recorder pool
	recorderPool := sync.Pool{
		New: func() interface{} {
			return httptest.NewRecorder()
		},
	}

	// Create request once
	req := httptest.NewRequest(http.MethodGet, "/api/v1/hello/Joe", nil)
	req.Header.Set("Accept", "application/json")

	// Initial validation
	w := recorderPool.Get().(*httptest.ResponseRecorder)
	router.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		b.Fatalf("Initial request failed with status: %d", w.Code)
	}
	recorderPool.Put(w)

	// Benchmark
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			w := recorderPool.Get().(*httptest.ResponseRecorder)
			w.Body.Reset()
			router.ServeHTTP(w, req)

			if w.Code != http.StatusOK {
				b.Fatalf("Request failed with status: %d", w.Code)
			}
			recorderPool.Put(w)
		}
	})

	// Memory stats after
	runtime.ReadMemStats(&stats)
	heapAfter := stats.HeapAlloc

	b.ReportMetric(float64(heapAfter-heapBefore)/float64(b.N), "heap-delta-bytes/op")
	b.ReportMetric(float64(stats.NumGC), "GC-cycles")
}
