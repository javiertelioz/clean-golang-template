package hello

import (
	"strings"

	"github.com/javiertelioz/clean_architecture/pkg/domain/entities/shared"
)

type Hello struct {
	name string
}

func (h *Hello) GetName() string {
	return h.name
}

func (h *Hello) SayHello() string {
	return "Hello, " + h.name + "!"
}

func (h *Hello) Validate() *shared.ValidationErrors {
	errs := &shared.ValidationErrors{}

	if strings.TrimSpace(h.name) == "" {
		errs.Add(shared.NewDomainError("Name", "empty string", "Name cannot be empty", "Name is required", 3001))
	}

	return errs
}
