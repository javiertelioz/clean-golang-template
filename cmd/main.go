package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/javiertelioz/clean_architecture/pkg/application/use_cases/hello"
	"github.com/javiertelioz/clean_architecture/pkg/application/use_cases/payment"
	"github.com/javiertelioz/clean_architecture/pkg/infrastructure/database"
	"github.com/javiertelioz/clean_architecture/pkg/infrastructure/logger"
	"github.com/javiertelioz/clean_architecture/pkg/interfaces/controllers"
	"github.com/javiertelioz/clean_architecture/pkg/interfaces/routes"
)

func main() {

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Services
	loggerService := logger.NewLogger()

	// Uses Cases
	helloUseCase := hello.NewHelloUseCase()
	helloController := controllers.NewHelloController(helloUseCase, loggerService)

	repo := database.NewInMemoryPaymentRepository()
	paymentUseCase := payment.NewCreatePaymentUseCase(repo)
	paymentsController := controllers.NewPaymentController(paymentUseCase)

	r.Route("/api", func(r chi.Router) {
		r.Mount("/v1/hello", routes.NewHelloRoutes(helloController).Mount())
		r.Mount("/v1/payments", routes.NewPaymentRoutes(paymentsController).Mount())
	})

	log.Printf("🚀 Starting application on: http://%s/\n", "localhost:8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatalf("Could not listen on %s: %v\n", "localhost:8080", err)
	}
}
