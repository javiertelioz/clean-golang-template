package messaging

import (
	"log"

	contracts "github.com/javiertelioz/clean_architecture/pkg/domain/contracts/services"
)

type AsyncEventPublisher struct{}

func NewAsyncEventPublisher() *AsyncEventPublisher {
	return &AsyncEventPublisher{}
}

func (p *AsyncEventPublisher) Publish(event contracts.DomainEvent) error {
	go func() {
		log.Printf("[EVENT] %s: %+v\n", event.EventName(), event)
	}()
	return nil
}
