package repository

import (
	"errors"
	"sync"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"

	domain "github.com/javiertelioz/clean_architecture/pkg/domain/entities/payment"
)

type MockPaymentRepository struct {
	mock.Mock
	storage map[string]*domain.Payment
	mu      sync.RWMutex
}

func NewMockPaymentRepository() *MockPaymentRepository {
	return &MockPaymentRepository{
		storage: make(map[string]*domain.Payment),
	}
}

func (m *MockPaymentRepository) Save(p *domain.Payment) (*domain.Payment, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	args := m.Called(p)

	if p.GetID() == "" {
		newID := uuid.New().String()

		p = domain.NewPayment(
			domain.WithID(newID),
			domain.WithUserID(p.GetUserID()),
			domain.WithAmount(p.GetAmount()),
			domain.WithMethod(p.GetMethod()),
			domain.WithTraceabilityBranchID(p.GetTraceability().BranchID),
			domain.WithTraceabilityExternalID(p.GetTraceability().ExternalID),
			domain.WithTraceabilityTerminalID(p.GetTraceability().TerminalID),
		)

		if p.GetTraceability().Date != nil {
			p = domain.NewPayment(
				domain.WithID(newID),
				domain.WithUserID(p.GetUserID()),
				domain.WithAmount(p.GetAmount()),
				domain.WithMethod(p.GetMethod()),
				domain.WithTraceabilityBranchID(p.GetTraceability().BranchID),
				domain.WithTraceabilityExternalID(p.GetTraceability().ExternalID),
				domain.WithTraceabilityTerminalID(p.GetTraceability().TerminalID),
				domain.WithTraceabilityDate(*p.GetTraceability().Date),
			)
		}
	}

	m.storage[p.GetID()] = p
	return args.Get(0).(*domain.Payment), args.Error(1)
}

func (m *MockPaymentRepository) Approve(id string) (*domain.Payment, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if payment, ok := m.storage[id]; ok {
		// Aquí podrías actualizar un campo payment.status si existiera
		return payment, nil
	}

	return nil, errors.New("payment not found")
}

func (m *MockPaymentRepository) Update(payment *domain.Payment) (*domain.Payment, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, ok := m.storage[payment.GetID()]; ok {
		m.storage[payment.GetID()] = payment
		return payment, nil
	}

	return nil, errors.New("payment not found")
}
