package payment

import (
	"github.com/javiertelioz/clean_architecture/pkg/domain/entities/payment/value_object"
	"github.com/javiertelioz/clean_architecture/pkg/domain/entities/shared"
)

type Payment struct {
	id           string
	amount       value_object.Amount
	userID       string
	method       value_object.Method
	traceability value_object.Traceability
}

func (p *Payment) GetID() string {
	return p.id
}

func (p *Payment) GetAmount() value_object.Amount {
	return p.amount
}

func (p *Payment) GetUserID() string {
	return p.userID
}

func (p *Payment) GetMethod() value_object.Method {
	return p.method
}

func (p *Payment) GetTraceability() value_object.Traceability {
	return p.traceability
}

func (p *Payment) Validate() *shared.ValidationErrors {
	errs := &shared.ValidationErrors{}

	if err := value_object.ValidateAmount(p.amount); err != nil {
		errs.Add(err)
	}
	if err := value_object.ValidateMethod(p.method); err != nil {
		errs.Add(err)
	}
	if err := value_object.ValidateTraceability(p.traceability); err != nil {
		errs.Add(err)
	}
	if p.userID == "" {
		errs.Add(shared.NewDomainError("UserID", "empty string", "UserID cannot be empty", "UserID is required", 3001))
	}

	return errs
}
