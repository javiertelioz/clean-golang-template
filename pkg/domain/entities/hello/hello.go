package hello

import (
	"fmt"
	"strings"

	"github.com/javiertelioz/clean_architecture/pkg/domain/entities/shared"
)

type Hello struct {
	name      string
	timestamp int64
}

func (h *Hello) GetName() string {
	return h.name
}

func (h *Hello) GetTimestamp() int64 {
	return h.timestamp
}

func (h *Hello) SayHello() string {
	return fmt.Sprintf("Hello, %s!", h.GetName())
}

func (h *Hello) Validate() *shared.ValidationErrors {
	errs := &shared.ValidationErrors{}

	if strings.TrimSpace(h.name) == "" {
		errs.Add(shared.NewDomainError("Name", "empty string", "Name cannot be empty", "Name is required", 3001))
	}

	return errs
}
