package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/javiertelioz/clean_architecture/pkg/interfaces/controllers"
)

type PaymentRoutes struct {
	router     chi.Router
	controller *controllers.PaymentController
}

func NewPaymentRoutes(controller *controllers.PaymentController) *PaymentRoutes {
	router := chi.NewRouter()
	return &PaymentRoutes{
		router:     router,
		controller: controller,
	}
}

func (pr *PaymentRoutes) Mount() chi.Router {
	pr.router.Post("/", pr.controller.CreatePaymentHandler)
	/*pr.router.Get("/", pr.controller.GetPaymentsHandler)
	pr.router.Get("/{id}", pr.controller.GetPaymentByIdHandler)
	pr.router.Put("/{id}", pr.controller.UpdatePaymentHandler)
	pr.router.Delete("/{id}", pr.controller.DeletePaymentHandler)*/

	return pr.router
}
