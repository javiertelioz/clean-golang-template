package payment

import (
	"time"
	
	"github.com/javiertelioz/clean_architecture/pkg/domain/entities/payment/value_object"
)

type PaymentOption func(*Payment)

func NewPayment(opts ...PaymentOption) *Payment {
	p := &Payment{}
	for _, opt := range opts {
		opt(p)
	}
	return p
}

func WithID(id string) PaymentOption {
	return func(p *Payment) {
		p.id = id
	}
}

func WithAmount(amount value_object.Amount) PaymentOption {
	return func(p *Payment) {
		p.amount = amount
	}
}

func WithUserID(userID string) PaymentOption {
	return func(p *Payment) {
		p.userID = userID
	}
}

func WithMethod(method value_object.Method) PaymentOption {
	return func(p *Payment) {
		p.method = method
	}
}

func WithTraceability(trace value_object.Traceability) PaymentOption {
	return func(p *Payment) {
		p.traceability = trace
	}
}

func WithTraceabilityDate(date time.Time) PaymentOption {
	return func(p *Payment) {
		if p.traceability == (value_object.Traceability{}) {
			p.traceability = value_object.Traceability{}
		}
		p.traceability.Date = &date
	}
}

func WithTraceabilityBranchID(branchID string) PaymentOption {
	return func(p *Payment) {
		if p.traceability == (value_object.Traceability{}) {
			p.traceability = value_object.Traceability{}
		}
		p.traceability.BranchID = branchID
	}
}

func WithTraceabilityExternalID(externalID string) PaymentOption {
	return func(p *Payment) {
		if p.traceability == (value_object.Traceability{}) {
			p.traceability = value_object.Traceability{}
		}
		p.traceability.ExternalID = externalID
	}
}

func WithTraceabilityTerminalID(terminalID string) PaymentOption {
	return func(p *Payment) {
		if p.traceability == (value_object.Traceability{}) {
			p.traceability = value_object.Traceability{}
		}
		p.traceability.TerminalID = terminalID
	}
}
