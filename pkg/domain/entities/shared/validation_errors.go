package shared

import (
	"strings"
)

type ValidationErrors struct {
	Errors []*DomainError
}

func (ve *ValidationErrors) Add(err *DomainError) {
	ve.Errors = append(ve.Errors, err)
}

func (ve *ValidationErrors) Error() string {
	var messages []string
	for _, err := range ve.Errors {
		messages = append(messages, err.Error())
	}
	return strings.Join(messages, ", ")
}

func (ve *ValidationErrors) HasErrors() bool {
	return len(ve.Errors) > 0
}

func (ve *ValidationErrors) IsEmpty() bool {
	return len(ve.Errors) == 0
}
