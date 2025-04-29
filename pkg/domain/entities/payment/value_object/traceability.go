package value_object

import (
	"github.com/javiertelioz/clean_architecture/pkg/domain/entities/shared"
	"time"
)

type Traceability struct {
	BranchID   string
	Date       *time.Time
	ExternalID string
	TerminalID string
}

func ValidateTraceability(t Traceability) *shared.DomainError {
	if t.BranchID == "" || t.Date == nil || t.ExternalID == "" || t.TerminalID == "" {
		return shared.NewDomainError("Traceability", "missing fields", "Incomplete traceability info", "All traceability fields are required", 3004)
	}
	return nil
}
