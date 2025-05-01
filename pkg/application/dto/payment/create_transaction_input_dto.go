package payment

import (
	"time"

	domain "github.com/javiertelioz/clean_architecture/pkg/domain/entities/payment"
	"github.com/javiertelioz/clean_architecture/pkg/domain/entities/payment/value_object"
)

type CreateTransactionDto struct {
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

func (dto *CreateTransactionDto) ToDomainOptions() []domain.PaymentOption {
	opts := []domain.PaymentOption{
		domain.WithUserID(dto.UserID),
		domain.WithAmount(value_object.Amount(dto.Amount)),
		domain.WithMethod(value_object.Method(dto.Method)),
		domain.WithTraceabilityBranchID(dto.Traceability.BranchId),
		domain.WithTraceabilityExternalID(dto.Traceability.ExternalId),
		domain.WithTraceabilityTerminalID(dto.Traceability.TerminalId),
	}

	if dto.Traceability.Date != nil {
		opts = append(opts, domain.WithTraceabilityDate(*dto.Traceability.Date))
	}

	return opts
}
