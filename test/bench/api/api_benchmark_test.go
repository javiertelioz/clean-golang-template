package api_test

import (
	"net/http"
	"testing"
)

func BenchmarkHelloEndpoint(b *testing.B) {
	client := &http.Client{}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		resp, err := client.Get("http://localhost:8080/api/v1/hello/world")
		if err != nil {
			b.Fatal(err)
		}
		resp.Body.Close()
	}
}

func BenchmarkPaymentsEndpoint(b *testing.B) {
	client := &http.Client{}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		resp, err := client.Post("http://localhost:8080/api/v1/payments", "application/json", nil)
		if err != nil {
			b.Fatal(err)
		}
		resp.Body.Close()
	}
}
