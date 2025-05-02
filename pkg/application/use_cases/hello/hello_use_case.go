package hello

import (
	dto "github.com/javiertelioz/clean_architecture/pkg/application/dto/hello"
	"github.com/javiertelioz/clean_architecture/pkg/application/events"
	"github.com/javiertelioz/clean_architecture/pkg/domain/contracts/services"
	"github.com/javiertelioz/clean_architecture/pkg/domain/entities/hello"
)

type HelloUseCase struct {
	publisherService services.EventPublisher
	loggerService    services.LoggerService
}

func NewHelloUseCase(
	publisherService services.EventPublisher,
	loggerService services.LoggerService,
) *HelloUseCase {
	return &HelloUseCase{
		publisherService: publisherService,
		loggerService:    loggerService,
	}
}

func (uc *HelloUseCase) Execute(input *dto.HelloInput) (*dto.HelloOutput, error) {
	h := hello.NewHello(input.ToDomainOptions()...)

	if errs := h.Validate(); errs.HasErrors() {
		// uc.loggerService.Error(fmt.Sprintf("Invalid input: %s", errs))
		return nil, errs
	}

	output := &dto.HelloOutput{
		Message:   h.SayHello(),
		Timestamp: h.GetTimestamp(),
	}

	event := events.NewHelloGreetedWasSuccessful(output.Message, output.Timestamp)
	_ = uc.publisherService.Publish(event)

	return output, nil
}
