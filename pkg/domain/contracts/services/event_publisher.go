package services

type DomainEvent interface {
	EventName() string
}

type EventPublisher interface {
	Publish(event DomainEvent) error
}
