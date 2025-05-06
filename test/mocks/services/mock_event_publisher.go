package services

import (
	"sync"

	contracts "github.com/javiertelioz/clean_architecture/pkg/domain/contracts/services"
)

type MockEventPublisher struct {
	mu        sync.Mutex
	Published []contracts.DomainEvent
}

func NewMockEventPublisher() *MockEventPublisher {
	return &MockEventPublisher{
		Published: make([]contracts.DomainEvent, 0),
	}
}

func (m *MockEventPublisher) Publish(event contracts.DomainEvent) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.Published = append(m.Published, event)
	return nil
}

func (m *MockEventPublisher) GetPublishedEvents() []contracts.DomainEvent {
	m.mu.Lock()
	defer m.mu.Unlock()
	return append([]contracts.DomainEvent(nil), m.Published...) // safe copy
}
