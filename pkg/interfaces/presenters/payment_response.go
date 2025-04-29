package presenters

import "time"

type PaymentResponse struct {
	ID           string               `json:"id"`
	UserID       string               `json:"user_id"`
	Amount       float64              `json:"amount"`
	Method       string               `json:"method"`
	Status       string               `json:"status"`
	Traceability TraceabilityResponse `json:"traceability"`
}

type TraceabilityResponse struct {
	BranchID   string     `json:"branch_id"`
	Date       *time.Time `json:"date"`
	ExternalID string     `json:"external_id"`
	TerminalID string     `json:"terminal_id"`
}
