package hello

import domain "github.com/javiertelioz/clean_architecture/pkg/domain/entities/hello"

type HelloInput struct {
	Name string
}

func (input *HelloInput) ToDomainOptions() []domain.HelloOption {
	return []domain.HelloOption{
		domain.WithName(input.Name),
	}
}
