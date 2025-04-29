package value_object

import (
	"strings"

	"github.com/javiertelioz/clean_architecture/pkg/domain/entities/shared"
)

type Method string

func ValidateMethod(method Method) *shared.DomainError {
	if strings.TrimSpace(string(method)) == "" {
		return shared.NewDomainError("Method", "empty", "Payment method is required", "Empty method value", 3003)
	}
	return nil
}
