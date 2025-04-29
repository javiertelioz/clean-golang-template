package value_object

import "github.com/javiertelioz/clean_architecture/pkg/domain/entities/shared"

type Amount float64

func ValidateAmount(amount Amount) *shared.DomainError {
	if amount <= 0 {
		return shared.NewDomainError("Amount", "non-positive", "Amount must be greater than 0", "Invalid amount", 3002)
	}
	return nil
}
