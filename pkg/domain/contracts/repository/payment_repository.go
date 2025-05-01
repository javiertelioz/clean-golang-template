package repository

import (
	"github.com/javiertelioz/clean_architecture/pkg/domain/entities/payment"
)

type PaymentRepository interface {
	Save(payment *payment.Payment) (*payment.Payment, error)
	Update(payment *payment.Payment) (*payment.Payment, error)
	Approve(id string) (*payment.Payment, error)
}
