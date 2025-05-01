package hello

import domain "github.com/javiertelioz/clean_architecture/pkg/domain/entities/hello"

type HelloInput struct {
	Name string
}

func (dto *HelloInput) ToDomainOptions() []domain.HelloOption {
	return []domain.HelloOption{
		domain.WithName(dto.Name),
	}
}
