package hello

import (
	dto "github.com/javiertelioz/clean_architecture/pkg/application/dto/hello"
	"github.com/javiertelioz/clean_architecture/pkg/domain/entities/hello"
)

type HelloUseCase struct{}

func NewHelloUseCase() *HelloUseCase {
	return &HelloUseCase{}
}

func (uc *HelloUseCase) Execute(input *dto.HelloInput) (*hello.Hello, error) {
	newHello := hello.NewHello(input.ToDomainOptions()...)

	if errs := newHello.Validate(); errs.HasErrors() {
		return nil, errs
	}

	return newHello, nil

	/*return presenters.HelloResponse{
		Message:   "Hello, " + name + "!",
		Code:      200,
		Timestamp: time.Now().UnixMilli(),
	}*/
}
