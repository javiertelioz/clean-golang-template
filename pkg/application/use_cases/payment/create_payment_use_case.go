package payment

import (
	dto "github.com/javiertelioz/clean_architecture/pkg/application/dto/payment"
	"github.com/javiertelioz/clean_architecture/pkg/domain/contracts/repository"
	"github.com/javiertelioz/clean_architecture/pkg/domain/entities/payment"
)

type CreatePaymentUseCase struct {
	PaymentRepository repository.PaymentRepository
}

func NewCreatePaymentUseCase(paymentRepository repository.PaymentRepository) *CreatePaymentUseCase {
	return &CreatePaymentUseCase{
		PaymentRepository: paymentRepository,
	}
}

func (uc *CreatePaymentUseCase) Execute(input *dto.CreateTransactionDto) (*payment.Payment, error) {
	newPayment := payment.NewPayment(input.ToDomainOptions()...)

	errs := newPayment.Validate()
	if !errs.IsEmpty() {
		return nil, errs
	}

	createdPayment, err := uc.PaymentRepository.Save(newPayment)
	if err != nil {
		return nil, err
	}

	return createdPayment, nil
}
