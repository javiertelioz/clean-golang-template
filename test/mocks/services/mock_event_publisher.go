package services

import contracts "github.com/javiertelioz/clean_architecture/pkg/domain/contracts/services"

type MockEventPublisher struct {
	Published []contracts.DomainEvent
}

func NewMockEventPublisher() *MockEventPublisher {
	return &MockEventPublisher{
		Published: []contracts.DomainEvent{},
	}
}
func (m *MockEventPublisher) Publish(event contracts.DomainEvent) error {
	m.Published = append(m.Published, event)
	return nil
}
