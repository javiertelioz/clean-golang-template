package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/javiertelioz/clean_architecture/pkg/application/dto/payment"
	usecase "github.com/javiertelioz/clean_architecture/pkg/application/use_cases/payment"
	"github.com/javiertelioz/clean_architecture/pkg/interfaces/presenters"
)

type PaymentController struct {
	createPayment *usecase.CreatePaymentUseCase
}

func NewPaymentController(createPayment *usecase.CreatePaymentUseCase) *PaymentController {
	return &PaymentController{
		createPayment: createPayment,
	}
}

func (pc *PaymentController) CreatePaymentHandler(w http.ResponseWriter, r *http.Request) {
	var input payment.CreateTransactionInput

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	p, err := pc.createPayment.Execute(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	response := presenters.PaymentResponse{
		ID:     p.GetID(),
		UserID: p.GetUserID(),
		Amount: float64(p.GetAmount()),
		Method: string(p.GetMethod()),
		Status: "created", // puedes mapearlo de la entidad si agregas estado
		Traceability: presenters.TraceabilityResponse{
			BranchID:   p.GetTraceability().BranchID,
			Date:       p.GetTraceability().Date,
			ExternalID: p.GetTraceability().ExternalID,
			TerminalID: p.GetTraceability().TerminalID,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
