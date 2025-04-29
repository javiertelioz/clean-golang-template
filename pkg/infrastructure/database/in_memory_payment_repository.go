package database

import (
	"errors"
	"sync"

	"github.com/google/uuid"
	"github.com/javiertelioz/clean_architecture/pkg/domain/contracts/repository"
	"github.com/javiertelioz/clean_architecture/pkg/domain/entities/payment"
)

type InMemoryPaymentRepository struct {
	data map[string]*payment.Payment
	mu   sync.RWMutex
}

func NewInMemoryPaymentRepository() repository.PaymentRepository {
	return &InMemoryPaymentRepository{
		data: make(map[string]*payment.Payment),
	}
}

func (r *InMemoryPaymentRepository) Save(p *payment.Payment) (*payment.Payment, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if p.GetID() == "" {
		id := uuid.New().String()
		p = payment.NewPayment(append([]payment.PaymentOption{
			payment.WithID(id),
			payment.WithUserID(p.GetUserID()),
			payment.WithAmount(p.GetAmount()),
			payment.WithMethod(p.GetMethod()),
			payment.WithTraceabilityBranchID(p.GetTraceability().BranchID),
			payment.WithTraceabilityExternalID(p.GetTraceability().ExternalID),
			payment.WithTraceabilityTerminalID(p.GetTraceability().TerminalID),
		}, conditionalDateOption(p)...)...)
	}

	r.data[p.GetID()] = p
	return p, nil
}

func (r *InMemoryPaymentRepository) Update(p *payment.Payment) (*payment.Payment, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.data[p.GetID()]; !exists {
		return nil, errors.New("payment not found")
	}

	r.data[p.GetID()] = p
	return p, nil
}

func (r *InMemoryPaymentRepository) Approve(id string) (*payment.Payment, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	p, exists := r.data[id]
	if !exists {
		return nil, errors.New("payment not found")
	}

	// Si agregas campo p.SetStatus("approved"), lo haces aquí
	// Por ahora solo lo devuelve como si fuera aprobado
	return p, nil
}

func conditionalDateOption(p *payment.Payment) []payment.PaymentOption {
	if p.GetTraceability().Date != nil {
		return []payment.PaymentOption{payment.WithTraceabilityDate(*p.GetTraceability().Date)}
	}
	return []payment.PaymentOption{}
}
