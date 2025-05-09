package payment

import (
	"time"

	domain "github.com/javiertelioz/clean_architecture/pkg/domain/entities/payment"
	"github.com/javiertelioz/clean_architecture/pkg/domain/entities/payment/value_object"
)

type CreateTransactionInput struct {
	UserID       string          `json:"user_id"`
	Amount       float64         `json:"amount"`
	Method       string          `json:"method"`
	ReferenceID  string          `json:"reference_id"`
	Traceability TraceabilityDto `json:"traceability"`
}

type TraceabilityDto struct {
	BranchId   string     `json:"branch_id"`
	Date       *time.Time `json:"date"`
	ExternalId string     `json:"external_id"`
	TerminalId string     `json:"terminal_id"`
}

func (input *CreateTransactionInput) ToDomainOptions() []domain.PaymentOption {
	opts := []domain.PaymentOption{
		domain.WithUserID(input.UserID),
		domain.WithAmount(value_object.Amount(input.Amount)),
		domain.WithMethod(value_object.Method(input.Method)),
		domain.WithTraceabilityBranchID(input.Traceability.BranchId),
		domain.WithTraceabilityExternalID(input.Traceability.ExternalId),
		domain.WithTraceabilityTerminalID(input.Traceability.TerminalId),
	}

	if input.Traceability.Date != nil {
		opts = append(opts, domain.WithTraceabilityDate(*input.Traceability.Date))
	}

	return opts
}
